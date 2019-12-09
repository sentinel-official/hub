package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/sentinel-official/hub/x/vpn/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func QueryResolversCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers ",
		Short: "Query the Resolvers, to filter use --address to get the particular resolvers",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			var res []byte
			var err error

			address := viper.GetString(flagAddress)
			if address != "" {
				bytes, err := ctx.Codec.MarshalJSON(address)
				if err != nil {
					return err
				}
				res, _, err = ctx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolvers), bytes)
				if err != nil {
					return err
				}
			} else {
				res, _, err = ctx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolvers), nil)
				if err != nil {
					return err
				}
			}

			var resolvers types.Resolvers
			if res == nil {
				return nil
			}
			if err := ctx.Codec.UnmarshalJSON(res, &resolvers); err != nil {
				return err
			}

			ctx.PrintOutput(resolvers)
			return nil
		},
	}
	cmd.Flags().String(flagAddress, "", "Resolver address")
	return cmd
}
