package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/client/common"
	"github.com/sentinel-official/hub/x/session/types"
)

func querySession(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Query a session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			session, err := common.QuerySession(ctx, id)
			if err != nil {
				return err
			}

			fmt.Println(session)
			return nil
		},
	}

	return cmd
}

func querySessions(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sessions",
		Short: "Query sessions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

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

			skip, err := cmd.Flags().GetInt(flagSkip)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			var (
				address  sdk.AccAddress
				node     hub.NodeAddress
				sessions types.Sessions
			)

			if subscription > 0 {
				sessions, err = common.QuerySessionsForSubscription(ctx, subscription, skip, limit)
			} else if len(bech32Node) > 0 {
				node, err = hub.NodeAddressFromBech32(bech32Node)
				if err != nil {
					return err
				}

				sessions, err = common.QuerySessionsForNode(ctx, node, skip, limit)
			} else if len(bech32Address) > 0 {
				address, err = sdk.AccAddressFromBech32(bech32Address)
				if err != nil {
					return err
				}

				var (
					active = false
					status hub.Status
				)

				active, err = cmd.Flags().GetBool(flagActive)
				if err != nil {
					return err
				}

				if active {
					status = hub.StatusActive
				}

				sessions, err = common.QuerySessionsForAddress(ctx, address, status, skip, limit)
			} else {
				sessions, err = common.QuerySessions(ctx, skip, limit)
			}

			if err != nil {
				return err
			}

			for _, session := range sessions {
				fmt.Printf("%s\n\n", session)
			}

			return nil
		},
	}

	cmd.Flags().String(flagAddress, "", "account address")
	cmd.Flags().Uint64(flagSubscription, 0, "subscription ID")
	cmd.Flags().String(flagNodeAddress, "", "node address")
	cmd.Flags().Bool(flagActive, false, "active sessions only")
	cmd.Flags().Int(flagSkip, 0, "skip")
	cmd.Flags().Int(flagLimit, 25, "limit")

	return cmd
}
