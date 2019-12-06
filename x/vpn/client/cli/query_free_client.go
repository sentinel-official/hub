package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/bech32"

	"github.com/sentinel-official/hub/x/vpn/client/common"
)

func QueryFreeClientsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "free-clients",
		Short: "Query free clients of node",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id := viper.GetString(flagNodeID)

			freeClients, err := common.QueryQueryFreeClientsOfNode(ctx, id)
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

	cmd.Flags().String(flagNodeID, "", "Node ID")
	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

func QueryFreeNodesCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "free-nodes",
		Short: "Query free nodes of client",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			address := viper.GetString(flagAddress)

			freeNodes, err := common.QueryFreeNodesOfClient(ctx, address)
			if err != nil {
				return err
			}

			for _, freeNode := range freeNodes {
				fmt.Println(freeNode)
			}
			return nil
		},
	}

	cmd.Flags().String(flagAddress, "", "Account address")
	_ = cmd.MarkFlagRequired(flagAddress)
	return cmd
}
