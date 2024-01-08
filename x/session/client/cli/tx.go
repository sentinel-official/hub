// DO NOT COVER

package cli

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
)

func txStart() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [node-addr]",
		Short: "Start a session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := cmd.Flags().GetUint64(flagSubscriptionID)
			if err != nil {
				return err
			}

			addr, err := hubtypes.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgStartRequest(
				ctx.FromAddress,
				id,
				addr,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Uint64(flagSubscriptionID, 0, "")

	return cmd
}

func txUpdateDetails() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [session-id] [upload] [download] [duration]",
		Short: "Update the session details",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			upload, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			download, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			duration, err := time.ParseDuration(args[3])
			if err != nil {
				return err
			}

			signature, err := GetSignature(cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDetailsRequest(
				ctx.FromAddress.Bytes(),
				types.Proof{
					ID:        id,
					Duration:  duration,
					Bandwidth: hubtypes.NewBandwidthFromInt64(upload, download),
				},
				signature,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagSignature, "", "client signature of the bandwidth info")

	return cmd
}

func txEnd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end [session-id]",
		Short: "End a session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			rating, err := cmd.Flags().GetUint64(flagRating)
			if err != nil {
				return err
			}

			msg := types.NewMsgEndRequest(
				ctx.FromAddress,
				id,
				rating,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Uint64(flagRating, 0, "rate the session quality [0, 10]")

	return cmd
}
