package cli

import (
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	ckeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func PaymentCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payment",
		Short: "Pay amount for using the VPN service",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			nodeID := viper.GetString(flagNodeID)
			amount := viper.GetString(flagAmount)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			from, err := cliCtx.GetFromAddress()

			if err != nil {
				return err
			}

			account, err := cliCtx.GetAccount(from)

			if err != nil {
				return err
			}

			coins, err := csdkTypes.ParseCoins(amount)

			if err != nil {
				return err
			}

			if !account.GetCoins().IsAllGTE(coins) {
				return errors.Errorf("Address %s doesn't have enough coins to pay for this transaction.", from)
			}

			sessionsCountBytes, err := cliCtx.QueryStore(cdc.MustMarshalBinaryLengthPrefixed(sdkTypes.KeySessionsCount), sdkTypes.KeySession)

			if err != nil {
				return err
			}

			var sessionsCount uint64

			if sessionsCountBytes != nil {
				if err := cdc.UnmarshalBinaryLengthPrefixed(sessionsCountBytes, &sessionsCount); err != nil {
					return err
				}
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
			lockerID := sdkTypes.KeySession + "/" + from.String() + "/" + strconv.Itoa(int(sessionsCount))
			msgLockerCoins := hub.MsgLockCoins{
				LockerID: lockerID,
				Coins:    coins,
				PubKey:   pubKey,
			}

			passPhrase, err := ckeys.GetPassphrase(name)

			if err != nil {
				return err
			}

			signature, _, err := kb.Sign(name, passPhrase, msgLockerCoins.GetUnSignBytes())

			if err != nil {
				return err
			}

			msg := vpn.NewMsgPayVPNService(from, nodeID, lockerID, coins, pubKey, signature)

			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "VPN Node ID")
	cmd.Flags().String(flagAmount, "100sent", "Amount of coins that you want to lock")

	return cmd
}
