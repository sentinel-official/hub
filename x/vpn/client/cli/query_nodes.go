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

func QueryNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Query node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			node, err := common.QueryNode(ctx, args[0])
			if err != nil {
				return nil
			}
			
			fmt.Println(node)
			return nil
		},
	}
	
	return cmd
}

func QueryNodesCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Query nodes",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			address := viper.GetString(flagAddress)
			
			var nodes []types.Node
			if address != "" {
				nodes, err = common.QueryNodesOfAddress(ctx, address)
			} else {
				nodes, err = common.QueryAllNodes(ctx)
			}
			
			if err != nil {
				return err
			}
			
			for _, node := range nodes {
				fmt.Println(node)
			}
			
			return nil
		},
	}
	
	cmd.Flags().String(flagAddress, "", "Account address")
	
	return cmd
}
