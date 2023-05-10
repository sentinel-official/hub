// DO NOT COVER

package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func txShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "share [subscription-id] [account-addr] [bytes]",
		Short: "Add a quota for a subscription",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			bytes, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgShareRequest(
				ctx.FromAddress,
				id,
				addr,
				sdk.NewInt(bytes),
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txCancel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [subscription-id]",
		Short: "Cancel a subscription",
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

			msg := types.NewMsgCancelRequest(
				ctx.FromAddress,
				id,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
