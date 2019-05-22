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
		Short: "Get subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			id := sdkTypes.NewIDFromString(args[0])

			res, err := common.QuerySubscription(cliCtx, cdc, id)
			if err != nil {
				return err
			}
			if res == nil {
				return fmt.Errorf("subscription not found")
			}

			var subscription vpn.Subscription
			if err := cdc.UnmarshalJSON(res, &subscription); err != nil {
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

			nodeID := viper.GetString(flagNodeID)
			address := viper.GetString(flagAddress)

			var res []byte
			var err error

			if len(nodeID) != 0 {
				id := sdkTypes.NewIDFromString(nodeID)
				res, err = common.QuerySubscriptionsOfNode(cliCtx, cdc, id)
			} else if len(address) != 0 {
				var _address csdkTypes.AccAddress

				_address, err = csdkTypes.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				res, err = common.QuerySubscriptionsOAddress(cliCtx, cdc, _address)
			} else {
				res, err = common.QueryAllSubscriptions(cliCtx)
			}

			if err != nil {
				return err
			}
			if string(res) == "[]" || string(res) == "null" {
				return fmt.Errorf("no subscriptions found")
			}

			var subscriptions []vpn.Subscription
			if err := cdc.UnmarshalJSON(res, &subscriptions); err != nil {
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
