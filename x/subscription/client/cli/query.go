package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription [subscription-id]",
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

			res, err := qc.QuerySubscription(
				context.Background(),
				types.NewQuerySubscriptionRequest(
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

func querySubscriptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			accAddr, err := GetAccountAddress(cmd.Flags())
			if err != nil {
				return err
			}

			nodeAddr, err := GetNodeAddress(cmd.Flags())
			if err != nil {
				return err
			}

			planID, err := cmd.Flags().GetUint64(flagPlanID)
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

			if !accAddr.Empty() {
				res, err := qc.QuerySubscriptionsForAccount(
					context.Background(),
					types.NewQuerySubscriptionsForAccountRequest(
						accAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else if !nodeAddr.Empty() {
				res, err := qc.QuerySubscriptionsForNode(
					context.Background(),
					types.NewQuerySubscriptionsForNodeRequest(
						nodeAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else if planID > 0 {
				res, err := qc.QuerySubscriptionsForPlan(
					context.Background(),
					types.NewQuerySubscriptionsForPlanRequest(
						planID,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QuerySubscriptions(
				context.Background(),
				types.NewQuerySubscriptionsRequest(
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
	flags.AddPaginationFlagsToCmd(cmd, "subscriptions")
	cmd.Flags().String(flagAccountAddress, "", "query the subscriptions of an account address")
	cmd.Flags().String(flagNodeAddress, "", "query the subscriptions of a node address")
	cmd.Flags().Uint64(flagPlanID, 0, "query the subscriptions of a subscription plan")

	return cmd
}

func queryQuota() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quota [subscription-id] [account-addr]",
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

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			var (
				qc = types.NewQueryServiceClient(ctx)
			)

			res, err := qc.QueryQuota(
				context.Background(),
				types.NewQueryQuotaRequest(
					id,
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

func queryQuotas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quotas [subscription-id]",
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

			res, err := qc.QueryQuotas(
				context.Background(),
				types.NewQueryQuotasRequest(
					id,
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
