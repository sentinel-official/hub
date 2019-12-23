package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"github.com/sentinel-official/hub/x/vpn/client/common"
)

func QueryResolversCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers ",
		Short: "Query the Resolvers, to filter use --address to get the particular resolvers",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			address := viper.GetString(flagAddress)
			
			resolvers, err := common.QueryResolvers(ctx, address)
			if err != nil {
				return err
			}
			
			ctx.PrintOutput(resolvers)
			return nil
		},
	}
	cmd.Flags().String(flagAddress, "", "Resolver address")
	
	return cmd
}
