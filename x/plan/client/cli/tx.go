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
	"github.com/sentinel-official/hub/x/plan/types"
)

func txAdd(cdc *codec.Codec) *cobra.Command {
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

			bytes, err := cmd.Flags().GetInt64(flagBytes)
			if err != nil {
				return err
			}

			msg := types.NewMsgAdd(ctx.FromAddress.Bytes(), price, validity, sdk.NewInt(bytes))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagPrice, "", "plan price")
	cmd.Flags().String(flagValidity, "", "plan validity")
	cmd.Flags().Int64(flagBytes, 0, "plan bytes (upload + download)")

	_ = cmd.MarkFlagRequired(flagPrice)
	_ = cmd.MarkFlagRequired(flagValidity)
	_ = cmd.MarkFlagRequired(flagBytes)

	return cmd
}

func txSetStatus(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status-set",
		Short: "Set a plan status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetStatus(ctx.FromAddress.Bytes(), id, hub.StatusFromString(args[1]))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txAddNode(cdc *codec.Codec) *cobra.Command {
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

			msg := types.NewMsgAddNode(ctx.FromAddress.Bytes(), id, node)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txRemoveNode(cdc *codec.Codec) *cobra.Command {
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

			msg := types.NewMsgRemoveNode(ctx.FromAddress.Bytes(), id, node)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}
