// DO NOT COVER

package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node [node-addr]",
		Short: "Query a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := hubtypes.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryNode(
				context.Background(),
				types.NewQueryNodeRequest(
					addr,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryNodes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Query nodes",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := cmd.Flags().GetUint64(flagPlan)
			if err != nil {
				return err
			}

			status, err := hubtypes.StatusFromFlags(cmd.Flags())
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

			if id > 0 {
				res, err := qc.QueryNodesForPlan(
					context.Background(),
					types.NewQueryNodesForPlanRequest(
						id,
						status,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QueryNodes(
				context.Background(),
				types.NewQueryNodesRequest(
					status,
					pagination,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nodes")
	cmd.Flags().String(hubtypes.FlagStatus, "", "filter the nodes by status (active|inactive)")
	cmd.Flags().Uint64(flagPlan, 0, "query the nodes of a subscription plan")

	return cmd
}

func queryLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease [id]",
		Short: "Query a lease",
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

			res, err := qc.QueryLease(
				context.Background(),
				types.NewQueryLeaseRequest(
					id,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryLeases() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leases",
		Short: "Query leases",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
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

			res, err := qc.QueryLeases(
				context.Background(),
				types.NewQueryLeasesRequest(
					pagination,
				),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "leases")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node-params",
		Short: "Query node module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryParams(
				context.Background(),
				types.NewQueryParamsRequest(),
			)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
