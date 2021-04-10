package cli

import (
	"context"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/swap/types"
)

func querySwap() *cobra.Command {
	return &cobra.Command{
		Use:   "swap",
		Short: "Query a swap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			txHash, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QuerySwap(context.Background(),
				types.NewQuerySwapRequest(types.BytesToHash(txHash)))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
}

func querySwaps() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaps",
		Short: "Query swaps",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QuerySwaps(context.Background(),
				types.NewQuerySwapsRequest(pagination))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	return cmd
}
