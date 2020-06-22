package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/subscription/client/common"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func queryPlanCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "plan",
		Short: "Query plan",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := hub.ProvAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			plan, err := common.QueryPlan(ctx, address, id)
			if err != nil {
				return err
			}

			fmt.Println(plan)
			return nil
		},
	}
}

func queryPlansCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plans",
		Short: "Query plans",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			var plans types.Plans

			if len(s) > 0 {
				address, err := hub.ProvAddressFromBech32(s)
				if err != nil {
					return err
				}

				plans, err = common.QueryPlansOfProvider(ctx, address)
				if err != nil {
					return err
				}
			} else {
				plans, err = common.QueryPlans(ctx)
				if err != nil {
					return err
				}
			}

			for _, plan := range plans {
				fmt.Println(plan)
			}

			return nil
		},
	}

	cmd.Flags().String(flagProvider, "", "Provider address")
	return cmd
}
