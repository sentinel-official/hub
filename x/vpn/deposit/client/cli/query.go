package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/vpn/deposit/client/common"
)

func queryDepositCmd(cdc *codec.Codec) *cobra.Command {
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

func queryDepositsCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposits",
		Short: "Query deposits",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			deposits, err := common.QueryDeposits(ctx)
			if err != nil {
				return err
			}

			for _, deposit := range deposits {
				fmt.Println(deposit)
			}

			return nil
		},
	}
}
