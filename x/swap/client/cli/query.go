package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/swap/client/common"
	"github.com/sentinel-official/hub/x/swap/types"
)

func querySwap(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "swap",
		Short: "Query a swap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			txHash, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			fmt.Println(txHash, len(txHash))

			swap, err := common.QuerySwap(ctx, types.BytesToHash(txHash))
			if err != nil {
				return err
			}

			fmt.Println(swap)
			return nil
		},
	}
}

func querySwaps(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaps",
		Short: "Query swaps",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			skip, err := cmd.Flags().GetInt(flagSkip)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			swaps, err := common.QuerySwaps(ctx, skip, limit)
			if err != nil {
				return err
			}

			for _, swap := range swaps {
				fmt.Printf("%s\n\n", swap)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagSkip, 0, "skip")
	cmd.Flags().Int(flagLimit, 25, "limit")

	return cmd
}
