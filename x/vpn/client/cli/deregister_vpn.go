package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	ckeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

const (
	flagLockerID = "locker-id"
)

func DeregisterVPNCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deregister-vpn",
		Short: "deregister for sentinel vpn service",
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(authCli.GetAccountDecoder(cdc))

			vpnID := viper.GetString(flagNodeID)
			lockerID := viper.GetString(flagLockerID)

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

			pubKey := account.GetPubKey()

			msgReleaseCoins := hub.MsgReleaseCoins{
				LockerID: lockerID,
				PubKey:   pubKey,
			}
			kb, err := ckeys.GetKeyBase()

			if err != nil {
				return err
			}

			name, err := cliCtx.GetFromName()

			if err != nil {
				return err
			}

			passPhrase, err := ckeys.GetPassphrase(name)

			if err != nil {
				return err
			}

			signature, _, err := kb.Sign(name, passPhrase, msgReleaseCoins.GetUnSignBytes())

			if err != nil {
				return err
			}

			msg := vpn.NewMsgDeregisterVPN(from, vpnID, lockerID, pubKey, signature)

			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "VPN node ID")
	cmd.Flags().String(flagLockerID, "", "VPN locker ID")

	return cmd
}
