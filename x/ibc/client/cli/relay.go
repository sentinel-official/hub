package cli

import (
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	flagFromChainID      = "from-chain-id"
	flagFromChainNodeURI = "from-chain-node-uri"
	flagToChainID        = "to-chain-id"
	flagToChainNodeURI   = "to-chain-node-uri"
)

type relayCommander struct {
	cdc        *codec.Codec
	accDecoder auth.AccountDecoder

	address     csdkTypes.AccAddress
	ibcStoreKey string
	accStoreKey string

	logger log.Logger
}

func IBCRelayCmd(cdc *codec.Codec) *cobra.Command {
	cmdr := relayCommander{
		cdc:         cdc,
		accDecoder:  authCli.GetAccountDecoder(cdc),
		ibcStoreKey: "ibc",
		accStoreKey: "acc",
		logger:      log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
	}

	cmd := &cobra.Command{
		Use: "relay",
		Run: cmdr.runIBCRelay,
	}

	cmd.Flags().String(flagFromChainID, "", "Chain ID for ibc node to check outgoing packets")
	cmd.Flags().String(flagFromChainNodeURI, "tcp://localhost:26657", "<host>:<port> to tendermint rpc interface for this chain")
	cmd.Flags().String(flagToChainID, "", "Chain ID for ibc node to broadcast incoming packets")
	cmd.Flags().String(flagToChainNodeURI, "tcp://localhost:36657", "<host>:<port> to tendermint rpc interface for this chain")

	cmd.MarkFlagRequired(flagFromChainID)
	cmd.MarkFlagRequired(flagFromChainNodeURI)
	cmd.MarkFlagRequired(flagToChainID)
	cmd.MarkFlagRequired(flagToChainNodeURI)

	viper.BindPFlag(flagFromChainID, cmd.Flags().Lookup(flagFromChainID))
	viper.BindPFlag(flagFromChainNodeURI, cmd.Flags().Lookup(flagFromChainNodeURI))
	viper.BindPFlag(flagToChainID, cmd.Flags().Lookup(flagToChainID))
	viper.BindPFlag(flagToChainNodeURI, cmd.Flags().Lookup(flagToChainNodeURI))

	return cmd
}

func (c relayCommander) runIBCRelay(cmd *cobra.Command, args []string) {
	fromChainID := viper.GetString(flagFromChainID)
	fromChainNodeURI := viper.GetString(flagFromChainNodeURI)
	toChainID := viper.GetString(flagToChainID)
	toChainNodeURI := viper.GetString(flagToChainNodeURI)
	address, err := context.NewCLIContext().GetFromAddress()

	if err != nil {
		panic(err)
	}

	c.address = address
	c.loop(fromChainID, fromChainNodeURI, toChainID, toChainNodeURI)
}

func (c relayCommander) loop(fromChainID, fromChainNodeURI, toChainID, toChainNodeURI string) {
	cliCtx := context.NewCLIContext()
	name, err := cliCtx.GetFromName()

	if err != nil {
		panic(err)
	}

	passphrase, err := keys.ReadPassphraseFromStdin(name)

	if err != nil {
		panic(err)
	}

	ingressKey := ibc.IngressSequenceKey(fromChainID)
	eggressLengthKey := ibc.EgressLengthKey(toChainID)

	for {
		var processed, egressLength int64
		processedBytes, err := query(toChainNodeURI, ingressKey, c.ibcStoreKey)

		if err != nil {
			panic(err)
		}

		if processedBytes == nil {
			processed = 0
		} else if err = c.cdc.UnmarshalBinary(processedBytes, &processed); err != nil {
			panic(err)
		}

		egressLengthBytes, err := query(fromChainNodeURI, eggressLengthKey, c.ibcStoreKey)

		if err != nil {
			c.logger.Error("error querying outgoing packet list length", "err", err)
			break
		}

		if egressLengthBytes == nil {
			egressLength = 0
		} else if err = c.cdc.UnmarshalBinary(egressLengthBytes, &egressLength); err != nil {
			panic(err)
		}

		if egressLength > processed {
			c.logger.Info("Detected IBC packet", "number", egressLength-1)
		}

		seq := c.getSequence(toChainNodeURI)

		for i := processed; i < egressLength; i++ {
			egressbz, err := query(fromChainNodeURI, ibc.EgressKey(toChainID, i), c.ibcStoreKey)

			if err != nil {
				c.logger.Error("error querying egress packet", "err", err)
				break
			}

			err = c.broadcastTx(seq+i-processed, toChainNodeURI, c.refine(egressbz, i, passphrase))

			if err != nil {
				c.logger.Error("error broadcasting ingress packet", "err", err)
				break
			}

			c.logger.Info("Relayed IBC packet", "number", i)
		}

		time.Sleep(5 * time.Second)
	}
}

func query(nodeURI string, key []byte, storeName string) (res []byte, err error) {
	return context.NewCLIContext().WithNodeURI(nodeURI).QueryStore(key, storeName)
}

func (c relayCommander) broadcastTx(seq int64, nodeURI string, tx []byte) error {
	_, err := context.NewCLIContext().WithNodeURI(nodeURI).BroadcastTx(tx)

	return err
}

func (c relayCommander) getSequence(nodeURI string) int64 {
	res, err := query(nodeURI, c.address, c.accStoreKey)

	if err != nil {
		panic(err)
	}

	if nil != res {
		account, err := c.accDecoder(res)

		if err != nil {
			panic(err)
		}

		return account.GetSequence()
	}

	return 0
}

func (c relayCommander) refine(bz []byte, sequence int64, passphrase string) []byte {
	var packet types.IBCPacket

	if err := c.cdc.UnmarshalBinary(bz, &packet); err != nil {
		panic(err)
	}

	msg := ibc.MsgIBCTransaction{
		Relayer:   c.address,
		Sequence:  sequence,
		IBCPacket: packet,
	}
	txBuilder := authTxBuilder.NewTxBuilderFromCLI().WithSequence(sequence).WithCodec(c.cdc)
	cliCtx := context.NewCLIContext()
	name, err := cliCtx.GetFromName()

	if err != nil {
		panic(err)
	}

	res, err := txBuilder.BuildAndSign(name, passphrase, []csdkTypes.Msg{msg})

	if err != nil {
		panic(err)
	}

	return res
}
