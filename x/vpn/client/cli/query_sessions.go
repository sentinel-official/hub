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

func QuerySessionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Query session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			session, err := common.QuerySession(cliCtx, cdc, args[0])
			if err != nil {
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
		Short: "Query sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			id := viper.GetString(flagSubscriptionID)

			var sessions []vpn.Session
			var err error

			if id != "" {
				sessions, err = common.QuerySessionsOfSubscription(cliCtx, cdc, id)
			} else {
				sessions, err = common.QueryAllSessions(cliCtx, cdc)
			}

			if err != nil {
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
