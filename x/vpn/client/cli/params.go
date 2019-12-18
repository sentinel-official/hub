package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/vpn/client/common"
)

func QueryParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current vpn parameters information",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			res, err := common.QueryParams(ctx)
			if err != nil {
				return err
			}

			return ctx.PrintOutput(res)
		},
	}
}
