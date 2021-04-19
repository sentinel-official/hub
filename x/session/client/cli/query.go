package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func querySession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
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

			res, err := qc.QuerySession(context.Background(),
				types.NewQuerySessionRequest(id))
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

			subscription, err := cmd.Flags().GetUint64(flagSubscription)
			if err != nil {
				return err
			}

			bech32Node, err := cmd.Flags().GetString(flagNodeAddress)
			if err != nil {
				return err
			}

			bech32Address, err := cmd.Flags().GetString(flagAddress)
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

			if subscription > 0 {
				res, err := qc.QuerySessionsForSubscription(context.Background(),
					types.NewQuerySessionsForSubscriptionRequest(subscription, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else if len(bech32Node) > 0 {
				address, err := hubtypes.NodeAddressFromBech32(bech32Node)
				if err != nil {
					return err
				}

				res, err := qc.QuerySessionsForNode(context.Background(), types.NewQuerySessionsForNodeRequest(address, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else if len(bech32Address) > 0 {
				address, err := sdk.AccAddressFromBech32(bech32Address)
				if err != nil {
					return err
				}

				var (
					active bool
					status hubtypes.Status
				)

				active, err = cmd.Flags().GetBool(flagActive)
				if err != nil {
					return err
				}

				if active {
					status = hubtypes.StatusActive
				}

				res, err := qc.QuerySessionsForAddress(context.Background(),
					types.NewQuerySessionsForAddressRequest(address, status, pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			} else {
				res, err := qc.QuerySessions(context.Background(),
					types.NewQuerySessionsRequest(pagination))
				if err != nil {
					return err
				}

				return ctx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "sessions")
	cmd.Flags().String(flagAddress, "", "account address")
	cmd.Flags().Uint64(flagSubscription, 0, "subscription ID")
	cmd.Flags().String(flagNodeAddress, "", "node address")
	cmd.Flags().Bool(flagActive, false, "active sessions only")

	return cmd
}
