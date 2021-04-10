package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode() *cobra.Command {
	return &cobra.Command{
		Use:   "node",
		Short: "Query a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			address, err := hub.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryNode(context.Background(),
				types.NewQueryNodeRequest(address))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
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

			provider, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			plan, err := cmd.Flags().GetUint64(flagPlan)
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

				res, err := qc.QueryNodesForProvider(context.Background(),
					types.NewQueryNodesForProviderRequest(address, status, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else if plan > 0 {
				return nil
			} else {
				res, err := qc.QueryNodes(context.Background(),
					types.NewQueryNodesRequest(status, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}
		},
	}

	cmd.Flags().String(flagProvider, "", "provider address")
	cmd.Flags().Uint64(flagPlan, 0, "subscription plan ID")
	cmd.Flags().String(flagStatus, "", "status")

	return cmd
}
