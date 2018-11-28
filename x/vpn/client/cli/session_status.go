package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func UpdateSessionStatusCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-session-status",
		Short: "Update VPN session status",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(authCli.GetAccountDecoder(cdc))

			sessionID := viper.GetString(flagSessionID)
			status := viper.GetString(flagStatus)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			from, err := cliCtx.GetFromAddress()

			if err != nil {
				return err
			}

			msg := vpn.NewMsgUpdateSessionStatus(from, sessionID, status)

			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagSessionID, "", "Session ID")
	cmd.Flags().Bool(flagStatus, false, "Session status")

	return cmd
}
