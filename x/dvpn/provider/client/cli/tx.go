package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func getTxRegisterProviderCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			name, err := cmd.Flags().GetString(flagName)
			if err != nil {
				return err
			}

			website, err := cmd.Flags().GetString(flagWebsite)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(flagDescription)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterProvider(ctx.FromAddress, name, website, description)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagName, "", "Provider name")
	cmd.Flags().String(flagWebsite, "", "Provider website")
	cmd.Flags().String(flagDescription, "", "Provider description")

	_ = cmd.MarkFlagRequired(flagName)
	_ = cmd.MarkFlagRequired(flagWebsite)
	_ = cmd.MarkFlagRequired(flagDescription)

	return cmd
}
