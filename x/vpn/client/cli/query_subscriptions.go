package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/hub/x/vpn/client/common"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func QuerySubscriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Query subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			subscription, err := common.QuerySubscription(ctx, args[0])
			if err != nil {
				return err
			}

			fmt.Println(subscription)
			return nil
		},
	}

	return cmd
}

func QuerySubscriptionsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id := viper.GetString(flagNodeID)
			address := viper.GetString(flagAddress)

			var subscriptions []types.Subscription
			if id != "" {
				subscriptions, err = common.QuerySubscriptionsOfNode(ctx, id)
			} else if address != "" {
				subscriptions, err = common.QuerySubscriptionsOfAddress(ctx, address)
			} else {
				subscriptions, err = common.QueryAllSubscriptions(ctx)
			}

			if err != nil {
				return err
			}

			for _, subscription := range subscriptions {
				fmt.Println(subscription)
			}

			return nil
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")
	cmd.Flags().String(flagAddress, "", "Account address")

	return cmd
}
