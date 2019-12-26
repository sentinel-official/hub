package cli

import (
	"fmt"
	
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
			
			resolverID := viper.GetString(flagResolverID)
			resolvers, err := common.QueryResolvers(ctx, resolverID)
			if err != nil {
				return err
			}
			
			if len(resolvers) == 0 {
				return nil
			}
			
			for _, resolver := range resolvers {
				fmt.Println(resolver)
			}
			return nil
		},
	}
	cmd.Flags().String(flagResolverID, "", "Resolver address")
	
	return cmd
}
