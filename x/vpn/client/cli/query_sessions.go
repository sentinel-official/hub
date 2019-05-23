// nolint: dupl
package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func QuerySessionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Get session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			id := sdkTypes.NewIDFromString(args[0])

			res, err := common.QuerySession(cliCtx, cdc, id)
			if err != nil {
				return err
			}
			if res == nil {
				return fmt.Errorf("session not found")
			}

			var session vpn.Session
			if err := cdc.UnmarshalJSON(res, &session); err != nil {
				return err
			}

			fmt.Println(session)

			return nil
		},
	}

	return cmd
}

func QuerySessionsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sessions",
		Short: "Get sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			subscriptionID := viper.GetString(flagSubscriptionID)

			var res []byte
			var err error

			if subscriptionID != "" {
				id := sdkTypes.NewIDFromString(subscriptionID)
				res, err = common.QuerySessionsOfSubscription(cliCtx, cdc, id)
			} else {
				res, err = common.QueryAllSessions(cliCtx)
			}

			if err != nil {
				return err
			}
			if string(res) == "[]" || string(res) == "null" {
				return fmt.Errorf("no sessions found")
			}

			var sessions []vpn.Session
			if err := cdc.UnmarshalJSON(res, &sessions); err != nil {
				return err
			}

			for _, session := range sessions {
				fmt.Println(session)
			}

			return nil
		},
	}

	cmd.Flags().String(flagSubscriptionID, "", "Subscription ID")

	return cmd
}
