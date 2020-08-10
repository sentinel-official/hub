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

			page, err := cmd.Flags().GetInt(flagPage)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			var subscriptions types.Subscriptions
			if len(address) > 0 {
				address, err := sdk.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				subscriptions, err = common.QuerySubscriptionsForAddress(ctx, address, page, limit)
				if err != nil {
					return err
				}
			} else if plan > 0 {
				subscriptions, err = common.QuerySubscriptionsForPlan(ctx, plan, page, limit)
				if err != nil {
					return err
				}
			} else if len(nodeAddress) > 0 {
				address, err := hub.NodeAddressFromBech32(nodeAddress)
				if err != nil {
					return err
				}

				subscriptions, err = common.QuerySubscriptionsForNode(ctx, address, page, limit)
				if err != nil {
					return err
				}
			} else {
				subscriptions, err = common.QuerySubscriptions(ctx, page, limit)
				if err != nil {
					return err
				}
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
	cmd.Flags().Int(flagPage, 1, "page")
	cmd.Flags().Int(flagLimit, 0, "limit")

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

			page, err := cmd.Flags().GetInt(flagPlan)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagPlan)
			if err != nil {
				return err
			}

			quotas, err := common.QueryQuotas(ctx, id, page, limit)
			if err != nil {
				return err
			}

			for _, quota := range quotas {
				fmt.Printf("%s\n\n", quota)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagPage, 1, "page")
	cmd.Flags().Int(flagLimit, 0, "limit")

	return cmd
}
