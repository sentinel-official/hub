// nolint: dupl
package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func QuerySubscriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Query subscription",
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
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			id := viper.GetString(flagNodeID)
			address := viper.GetString(flagAddress)

			var subscriptions []vpn.Subscription
			var err error

			if id != "" {
				subscriptions, err = common.QuerySubscriptionsOfNode(cliCtx, cdc, id)
			} else if address != "" {
				subscriptions, err = common.QuerySubscriptionsOfAddress(cliCtx, cdc, address)
			} else {
				subscriptions, err = common.QueryAllSubscriptions(cliCtx, cdc)
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
