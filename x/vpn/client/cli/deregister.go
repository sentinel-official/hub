package cli

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	ckeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func DeregisterCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deregister",
		Short: "Deregister Sentinel VPN service node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			nodeID := viper.GetString(flagNodeID)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			from, err := cliCtx.GetFromAddress()

			if err != nil {
				return err
			}

			kb, err := ckeys.GetKeyBase()

			if err != nil {
				return err
			}

			name, err := cliCtx.GetFromName()

			if err != nil {
				return err
			}

			keyInfo, err := kb.Get(name)

			if err != nil {
				return err
			}

			pubKey := keyInfo.GetPubKey()

			var vpnDetails sdkTypes.VPNDetails
			vpnDetailsBytes, err := cliCtx.QueryStore(cdc.MustMarshalBinaryLengthPrefixed(nodeID), "vpn")

			if err := cdc.UnmarshalBinaryLengthPrefixed(vpnDetailsBytes, &vpnDetails); err != nil {
				return err
			}

			lockerID := vpnDetails.LockerID
			msgReleaseCoins := hub.MsgReleaseCoins{
				LockerID: lockerID,
				PubKey:   pubKey,
			}

			passPhrase, err := ckeys.GetPassphrase(name)

			if err != nil {
				return err
			}

			signature, _, err := kb.Sign(name, passPhrase, msgReleaseCoins.GetUnSignBytes())

			if err != nil {
				return err
			}

			msg := vpn.NewMsgDeregisterNode(from, nodeID, lockerID, pubKey, signature)

			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "VPN node ID")

	return cmd
}
