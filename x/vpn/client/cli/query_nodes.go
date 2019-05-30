// nolint:dupl
package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func QueryNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Query node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			node, err := common.QueryNode(cliCtx, cdc, args[0])
			if err != nil {
				return err
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
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			address := viper.GetString(flagAddress)

			var nodes []vpn.Node
			var err error

			if address != "" {
				nodes, err = common.QueryNodesOfAddress(cliCtx, cdc, address)
			} else {
				nodes, err = common.QueryAllNodes(cliCtx, cdc)
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
