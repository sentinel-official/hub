package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Query a subscription",
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

			res, err := qc.QuerySubscription(context.Background(),
				types.NewQuerySubscriptionRequest(id))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func querySubscriptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			bech32Address, err := cmd.Flags().GetString(flagAddress)
			if err != nil {
				return err
			}

			plan, err := cmd.Flags().GetUint64(flagPlan)
			if err != nil {
				return err
			}

			bech32Node, err := cmd.Flags().GetString(flagNodeAddress)
			if err != nil {
				return err
			}

			status, err := cmd.Flags().GetString(flagStatus)
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

			if len(bech32Address) > 0 {
				address, err := sdk.AccAddressFromBech32(bech32Address)
				if err != nil {
					return err
				}

				res, err := qc.QuerySubscriptionsForAddress(context.Background(),
					types.NewQuerySubscriptionsForAddressRequest(address, hubtypes.StatusFromString(status), pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else if plan > 0 {
				res, err := qc.QuerySubscriptionsForPlan(context.Background(),
					types.NewQuerySubscriptionsForPlanRequest(plan, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else if len(bech32Node) > 0 {
				address, err := hubtypes.NodeAddressFromBech32(bech32Node)
				if err != nil {
					return err
				}

				res, err := qc.QuerySubscriptionsForNode(context.Background(),
					types.NewQuerySubscriptionsForNodeRequest(address, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else {
				res, err := qc.QuerySubscriptions(context.Background(),
					types.NewQuerySubscriptionsRequest(pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "subscriptions")
	cmd.Flags().String(flagAddress, "", "account address")
	cmd.Flags().Uint64(flagPlan, 0, "plan ID")
	cmd.Flags().String(flagNodeAddress, "", "node address")
	cmd.Flags().String(flagStatus, "", "status")

	return cmd
}

func queryQuota() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quota [id] [address]",
		Short: "Query a quota of an address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			address, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryQuota(context.Background(),
				types.NewQueryQuotaRequest(id, address))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryQuotas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quotas",
		Short: "Query quotas of a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
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

			res, err := qc.QueryQuotas(context.Background(), types.NewQueryQuotasRequest(id, pagination))
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "quotas")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription-params",
		Short: "Query subscription module parameters",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryParams(context.Background(),
				types.NewQueryParamsRequest())
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
