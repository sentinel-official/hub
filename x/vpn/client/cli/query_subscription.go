package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func QuerySubscriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Get details of a session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			res, err := common.QuerySubscription(cliCtx, cdc, args[0])
			if err != nil {
				return err
			}

			sessionData, err := cdc.MarshalJSONIndent(res, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(sessionData))

			return nil
		},
	}

	return cmd
}
