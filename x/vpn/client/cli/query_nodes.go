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

func QueryNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Get node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			id := sdkTypes.NewIDFromString(args[0])

			res, err := common.QueryNode(cliCtx, cdc, id)
			if err != nil {
				return err
			}
			if res == nil {
				return fmt.Errorf("node not found")
			}

			var node vpn.Node
			if err := cdc.UnmarshalJSON(res, &node); err != nil {
				return err
			}

			fmt.Println(node)

			return nil
		},
	}

	return cmd
}

// nolint: gocyclo
func QueryNodesCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Get nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			address := viper.GetString(flagAddress)

			var res []byte
			var err error

			if len(address) != 0 {
				address, err := csdkTypes.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				res, err = common.QueryNodesOfAddress(cliCtx, cdc, address)
			} else {
				res, err = common.QueryAllNodes(cliCtx)
			}

			if err != nil {
				return err
			}
			if string(res) == "[]" || string(res) == "null" {
				return fmt.Errorf("no nodes found")
			}

			var nodes []vpn.Node
			if err := cdc.UnmarshalJSON(res, &nodes); err != nil {
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
