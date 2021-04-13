package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func queryPlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Query a plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryPlan(context.Background(),
				types.NewQueryPlanRequest(id))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryPlans() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plans",
		Short: "Query plans",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			provider, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			s, err := cmd.Flags().GetString(flagStatus)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var (
				status = hub.StatusFromString(s)
				qc     = types.NewQueryServiceClient(ctx)
			)

			if len(provider) > 0 {
				address, err := hub.ProvAddressFromBech32(provider)
				if err != nil {
					return err
				}

				res, err := qc.QueryPlansForProvider(context.Background(),
					types.NewQueryPlansForProviderRequest(address, status, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QueryPlans(context.Background(), types.NewQueryPlansRequest(status, pagination))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "providers")
	cmd.Flags().String(flagProvider, "", "provider address")
	cmd.Flags().String(flagStatus, "", "status")

	return cmd
}
