package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

// nolint:dupl
func QueryNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Query node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			id := sdkTypes.NewIDFromString(args[0])

			node, err := common.QueryNode(cliCtx, cdc, id)
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
				var _address csdkTypes.AccAddress

				_address, err = csdkTypes.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				nodes, err = common.QueryNodesOfAddress(cliCtx, cdc, _address)
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
