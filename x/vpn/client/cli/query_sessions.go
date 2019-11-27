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

func QuerySessionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Query session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			session, err := common.QuerySession(ctx, args[0])
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
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id := viper.GetString(flagSubscriptionID)

			var sessions []types.Session
			if id != "" {
				sessions, err = common.QuerySessionsOfSubscription(ctx, id)
			} else {
				sessions, err = common.QueryAllSessions(ctx)
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
