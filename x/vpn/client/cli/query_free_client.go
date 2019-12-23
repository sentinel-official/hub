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

func QueryFreeClientsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "free-clients [nodeID]",
		Short: "Query free clients of node",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			freeClients, err := common.QueryFreeClientsOfNode(ctx, args[0])
			if err != nil {
				return err
			}
			
			bech32AccAddPrefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
			for _, freeClient := range freeClients {
				_freeClient, err := bech32.ConvertAndEncode(bech32AccAddPrefix, freeClient)
				if err != nil {
					return err
				}
				fmt.Println(_freeClient)
			}
			
			return nil
		},
	}
	
	return cmd
}

func QueryFreeNodesCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "free-nodes [address]",
		Short: "Query free nodes of client",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			freeNodes, err := common.QueryFreeNodesOfClient(ctx, args[0])
			if err != nil {
				return err
			}
			
			for _, freeNode := range freeNodes {
				fmt.Println(freeNode)
			}
			return nil
		},
	}
	
	return cmd
}
