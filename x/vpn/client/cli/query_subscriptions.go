// nolint: dupl
package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func QuerySubscriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Get subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			subscription, err := common.QuerySubscription(cliCtx, cdc, args[0])
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
		Short: "Get subscriptions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			var subscriptions []vpn.Subscription
			res, err := cliCtx.QuerySubspace(vpn.SubscriptionKeyPrefix, vpn.StoreKeySubscription)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("no subscriptions found")
			}

			for _, kv := range res {
				var subscription vpn.Subscription
				if err := cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &subscription); err != nil {
					return err
				}

				subscriptions = append(subscriptions, subscription)
			}

			for _, subscription := range subscriptions {
				fmt.Println(subscription)
			}

			return nil
		},
	}

	return cmd
}
