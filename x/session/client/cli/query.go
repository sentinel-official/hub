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

			address, err := cmd.Flags().GetString(flagAddress)
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

			if len(address) > 0 {
				address, err := sdk.AccAddressFromBech32(address)
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
	cmd.Flags().Bool(flagActive, false, "active sessions only")

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
