package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func txAddPlanCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan-add",
		Short: "Add plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			s, err := cmd.Flags().GetString(flagPrice)
			if err != nil {
				return err
			}

			price, err := sdk.ParseCoins(s)
			if err != nil {
				return err
			}

			s, err = cmd.Flags().GetString(flagDuration)
			if err != nil {
				return err
			}

			duration, err := time.ParseDuration(s)
			if err != nil {
				return err
			}

			upload, err := cmd.Flags().GetUint64(flagMaxUpload)
			if err != nil {
				return err
			}

			download, err := cmd.Flags().GetUint64(flagMaxDownload)
			if err != nil {
				return err
			}

			s, err = cmd.Flags().GetString(flagMaxDuration)
			if err != nil {
				return err
			}

			maxDuration, err := time.ParseDuration(s)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddPlan(ctx.FromAddress.Bytes(), price, duration,
				hub.NewBandwidth(upload, download), maxDuration)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagPrice, "", "Plan price")
	cmd.Flags().String(flagDuration, "", "Plan duration")
	cmd.Flags().Uint64(flagMaxUpload, 0, "Plan max upload bandwidth")
	cmd.Flags().Uint64(flagMaxDownload, 0, "Plan max download bandwidth")
	cmd.Flags().String(flagMaxDuration, "", "Plan max duration")

	_ = cmd.MarkFlagRequired(flagPrice)
	_ = cmd.MarkFlagRequired(flagDuration)
	_ = cmd.MarkFlagRequired(flagMaxUpload)
	_ = cmd.MarkFlagRequired(flagMaxDownload)
	_ = cmd.MarkFlagRequired(flagMaxDuration)

	return cmd
}

func txSetPlanStatusCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan-set-status",
		Short: "Set plan status",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := hub.ProvAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			if !address.Equals(ctx.FromAddress) {
				return fmt.Errorf("provider address is not equal to from address")
			}

			id, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPlanStatus(ctx.FromAddress.Bytes(), id, hub.StatusFromString(args[2]))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}
