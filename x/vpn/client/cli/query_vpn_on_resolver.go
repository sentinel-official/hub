package cli

import (
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/bech32"
	
	"github.com/sentinel-official/hub/x/vpn/client/common"
)

func QueryResolversOfNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-of-node [node-id]",
		Short: "Query resolvers of node ",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			resolvers, err := common.QueryResolversOfNode(ctx, args[0])
			if err != nil {
				return err
			}
			
			bech32AccAddPrefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
			for _, resolver := range resolvers {
				_resolver, err := bech32.ConvertAndEncode(bech32AccAddPrefix, resolver)
				if err != nil {
					return err
				}
				fmt.Println(_resolver)
			}
			
			return nil
		},
	}
	
	return cmd
}

func QueryNodesOfResolverCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes-of-resolver [address]",
		Short: "Query nodes of resolver",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			nodes, err := common.QueryNodesOfResolver(ctx, args[0])
			if err != nil {
				return err
			}
			
			for _, node := range nodes {
				fmt.Println(node)
			}
			return nil
		},
	}
	
	return cmd
}
