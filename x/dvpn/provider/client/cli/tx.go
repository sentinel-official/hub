package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
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

func getTxUpdateProviderCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update provider",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := hub.ProvAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			if !address.Equals(ctx.FromAddress) {
				return fmt.Errorf("provider address does not belong from address")
			}

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

			msg := types.NewMsgUpdateProvider(address.Bytes(), name, website, description)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagName, "", "Provider name")
	cmd.Flags().String(flagWebsite, "", "Provider website")
	cmd.Flags().String(flagDescription, "", "Provider description")

	return cmd
}
