package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/subscription/client/common"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func queryPlanCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "plan",
		Short: "Query plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			plan, err := common.QueryPlan(ctx, id)
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

			provider, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			var plans types.Plans

			if len(provider) > 0 {
				address, err := hub.ProvAddressFromBech32(provider)
				if err != nil {
					return err
				}

				plans, err = common.QueryPlansForProvider(ctx, address)
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

func querySubscriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Query subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			membersOnly, err := cmd.Flags().GetBool(flagMembersOnly)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			if membersOnly {
				members, err := common.QueryMembersForSubscription(ctx, id)
				if err != nil {
					return err
				}

				for _, member := range members {
					fmt.Println(member.String())
				}

				return nil
			}

			subscription, err := common.QuerySubscription(ctx, id)
			if err != nil {
				return err
			}

			fmt.Println(subscription)
			return nil
		},
	}

	cmd.Flags().Bool(flagMembersOnly, false, "Show members only")
	return cmd
}

func querySubscriptionsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := cmd.Flags().GetString(flagAddress)
			if err != nil {
				return err
			}

			plan, err := cmd.Flags().GetUint64(flagPlan)
			if err != nil {
				return err
			}

			nodeAddress, err := cmd.Flags().GetString(flagNodeAddress)
			if err != nil {
				return err
			}

			var subscriptions types.Subscriptions

			if len(address) > 0 {
				address, err := sdk.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				subscriptions, err = common.QuerySubscriptionsForAddress(ctx, address)
				if err != nil {
					return err
				}
			} else if plan > 0 {
				subscriptions, err = common.QuerySubscriptionsForPlan(ctx, plan)
				if err != nil {
					return err
				}
			} else if len(nodeAddress) > 0 {
				address, err := hub.NodeAddressFromBech32(nodeAddress)
				if err != nil {
					return err
				}

				subscriptions, err = common.QuerySubscriptionsForNode(ctx, address)
				if err != nil {
					return err
				}
			} else {
				subscriptions, err = common.QuerySubscriptions(ctx)
				if err != nil {
					return err
				}
			}

			for _, subscription := range subscriptions {
				fmt.Println(subscription)
			}

			return nil
		},
	}

	cmd.Flags().String(flagAddress, "", "Account address")
	cmd.Flags().Uint64(flagPlan, 0, "Plan ID")
	cmd.Flags().String(flagNodeAddress, "", "Node address")

	return cmd
}
