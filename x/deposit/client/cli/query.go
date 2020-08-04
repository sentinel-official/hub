package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/deposit/client/common"
)

func queryDeposit(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit",
		Short: "Query a deposit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			deposit, err := common.QueryDeposit(ctx, address)
			if err != nil {
				return err
			}

			fmt.Println(deposit)
			return nil
		},
	}
}

func queryDeposits(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits",
		Short: "Query deposits",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			page, err := cmd.Flags().GetInt(flagPage)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			deposits, err := common.QueryDeposits(ctx, page, limit)
			if err != nil {
				return err
			}

			for _, deposit := range deposits {
				fmt.Println(deposit)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagPage, 1, "page")
	cmd.Flags().Int(flagLimit, 0, "limit")

	return cmd
}
