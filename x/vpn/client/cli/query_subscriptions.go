// nolint: dupl
package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
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

			id := sdkTypes.NewIDFromString(args[0])

			subscription, err := common.QuerySubscription(cliCtx, cdc, id)
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

			nodeID := viper.GetString(flagNodeID)
			address := viper.GetString(flagAddress)

			var subscriptions []vpn.Subscription
			var err error

			if nodeID != "" {
				id := sdkTypes.NewIDFromString(nodeID)
				subscriptions, err = common.QuerySubscriptionsOfNode(cliCtx, cdc, id)
			} else if address != "" {
				var _address csdkTypes.AccAddress

				_address, err = csdkTypes.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				subscriptions, err = common.QuerySubscriptionsOfAddress(cliCtx, cdc, _address)
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
