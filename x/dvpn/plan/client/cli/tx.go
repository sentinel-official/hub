package cli

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/plan/types"
)

func txAddPlanCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a plan",
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

			s, err = cmd.Flags().GetString(flagValidity)
			if err != nil {
				return err
			}

			validity, err := time.ParseDuration(s)
			if err != nil {
				return err
			}

			upload, err := cmd.Flags().GetInt64(flagUpload)
			if err != nil {
				return err
			}

			download, err := cmd.Flags().GetInt64(flagDownload)
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

			msg := types.NewMsgAddPlan(ctx.FromAddress.Bytes(), price, validity,
				hub.NewBandwidthFromInt64(upload, download), duration)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagPrice, "", "Plan price")
	cmd.Flags().String(flagValidity, "", "Plan validity")
	cmd.Flags().Int64(flagUpload, 0, "Plan upload bandwidth")
	cmd.Flags().Int64(flagDownload, 0, "Plan download bandwidth")
	cmd.Flags().String(flagDuration, "", "Plan duration")

	_ = cmd.MarkFlagRequired(flagPrice)
	_ = cmd.MarkFlagRequired(flagValidity)
	_ = cmd.MarkFlagRequired(flagUpload)
	_ = cmd.MarkFlagRequired(flagDownload)
	_ = cmd.MarkFlagRequired(flagDuration)

	return cmd
}

func txSetPlanStatusCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status-set",
		Short: "Set plan status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPlanStatus(ctx.FromAddress.Bytes(), id, hub.StatusFromString(args[1]))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txAddNodeForPlanCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node-add",
		Short: "Add a node for a plan",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			node, err := hub.NodeAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddNodeForPlan(ctx.FromAddress.Bytes(), id, node)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txRemoveNodeForPlanCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node-remove",
		Short: "Remove a node for a plan",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			node, err := hub.NodeAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveNodeForPlan(ctx.FromAddress.Bytes(), id, node)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}
