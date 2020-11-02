package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/client/common"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Query a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			subscription, err := common.QuerySubscription(ctx, id)
			if err != nil {
				return err
			}

			fmt.Println(subscription)
			return nil
		},
	}

	return cmd
}

func querySubscriptions(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

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

			skip, err := cmd.Flags().GetInt(flagSkip)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			var (
				address       sdk.AccAddress
				node          hub.NodeAddress
				subscriptions types.Subscriptions
			)

			if len(bech32Address) > 0 {
				address, err = sdk.AccAddressFromBech32(bech32Address)
				if err != nil {
					return err
				}

				subscriptions, err = common.QuerySubscriptionsForAddress(ctx, address, skip, limit)
			} else if plan > 0 {
				subscriptions, err = common.QuerySubscriptionsForPlan(ctx, plan, skip, limit)
			} else if len(bech32Node) > 0 {
				node, err = hub.NodeAddressFromBech32(bech32Node)
				if err != nil {
					return err
				}

				subscriptions, err = common.QuerySubscriptionsForNode(ctx, node, skip, limit)
			} else {
				subscriptions, err = common.QuerySubscriptions(ctx, skip, limit)
			}

			if err != nil {
				return err
			}

			for _, subscription := range subscriptions {
				fmt.Printf("%s\n\n", subscription)
			}

			return nil
		},
	}

	cmd.Flags().String(flagAddress, "", "account address")
	cmd.Flags().Uint64(flagPlan, 0, "plan ID")
	cmd.Flags().String(flagNodeAddress, "", "node address")
	cmd.Flags().Int(flagSkip, 0, "skip")
	cmd.Flags().Int(flagLimit, 25, "limit")

	return cmd
}

func queryQuota(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quota",
		Short: "Query a quota of a subscription",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			address, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			quota, err := common.QueryQuota(ctx, id, address)
			if err != nil {
				return err
			}

			fmt.Println(quota)
			return nil
		},
	}

	return cmd
}

func queryQuotas(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quotas",
		Short: "Query quotas of a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := strconv.ParseUint(args[0], 10, 64)
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

			quotas, err := common.QueryQuotas(ctx, id, skip, limit)
			if err != nil {
				return err
			}

			for _, quota := range quotas {
				fmt.Printf("%s\n\n", quota)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagSkip, 0, "skip")
	cmd.Flags().Int(flagLimit, 25, "limit")

	return cmd
}
