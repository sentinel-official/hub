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

func QueryResolversClientsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolvers-node",
		Short: "Query resolvers of node",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id := viper.GetString(flagNodeID)

			resolvers, err := common.QueryResolversOfNode(ctx, id)
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

	cmd.Flags().String(flagNodeID, "", "Node ID")
	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

func QueryNodesOfResolverCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes-resolver",
		Short: "Query nodes of resolver",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			address := viper.GetString(flagAddress)

			nodes, err := common.QueryNodesOfResolver(ctx, address)
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
	_ = cmd.MarkFlagRequired(flagAddress)

	return cmd
}
