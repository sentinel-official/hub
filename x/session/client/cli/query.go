// DO NOT COVER

package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/v1/x/session/types"
)

func querySession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session [session-id]",
		Short: "Query a session",
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

			res, err := qc.QuerySession(
				context.Background(),
				types.NewQuerySessionRequest(
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

func querySessions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sessions",
		Short: "Query sessions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			accAddr, err := GetAddress(cmd.Flags())
			if err != nil {
				return err
			}

			nodeAddr, err := GetNodeAddress(cmd.Flags())
			if err != nil {
				return err
			}

			subscriptionID, err := cmd.Flags().GetUint64(flagSubscriptionID)
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

			if accAddr != nil {
				res, err := qc.QuerySessionsForAccount(
					context.Background(),
					types.NewQuerySessionsForAccountRequest(
						accAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			if nodeAddr != nil {
				res, err := qc.QuerySessionsForNode(
					context.Background(),
					types.NewQuerySessionsForNodeRequest(
						nodeAddr,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			if subscriptionID != 0 {
				res, err := qc.QuerySessionsForSubscription(
					context.Background(),
					types.NewQuerySessionsForSubscriptionRequest(
						subscriptionID,
						pagination,
					),
				)
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}

			res, err := qc.QuerySessions(
				context.Background(),
				types.NewQuerySessionsRequest(
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
	flags.AddPaginationFlagsToCmd(cmd, "sessions")
	cmd.Flags().String(flagAddress, "", "filter the sessions by an account address")
	cmd.Flags().String(flagNodeAddress, "", "filter the sessions by a node address")
	cmd.Flags().Uint64(flagSubscriptionID, 0, "filter the sessions by a subscription id")

	return cmd
}

func queryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session-params",
		Short: "Query session module parameters",
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
