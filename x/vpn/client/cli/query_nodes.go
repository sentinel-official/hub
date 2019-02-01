package cli

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func QueryNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Args:  cobra.ExactArgs(1),
		Short: "Get details of a node",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			nodeID := args[0]

			params := vpn.NewQueryNodeParams(nodeID)
			paramBytes, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNode), paramBytes)
			if err != nil {
				return err
			}

			var node vpn.NodeDetails
			if err = cdc.UnmarshalJSON(res, &node); err != nil {
				return err
			}

			nodeData, err := json.MarshalIndent(node, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(nodeData))

			return nil
		},
	}

	return cmd
}

func queryNodes(cliCtx context.CLIContext, cdc *codec.Codec) ([]vpn.NodeDetails, error) {
	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodes), nil)
	if err != nil {
		return nil, err
	}

	var nodes []vpn.NodeDetails
	if err = cdc.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

func queryNodesOfOwner(cliCtx context.CLIContext, cdc *codec.Codec, owner csdkTypes.AccAddress) ([]vpn.NodeDetails, error) {
	params := vpn.NewQueryNodesOfOwnerParams(owner)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodes), paramBytes)
	if err != nil {
		return nil, err
	}

	var nodes []vpn.NodeDetails
	if err = cdc.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

func QueryNodesCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Get details of nodes",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			owner := viper.GetString(flagOwnerAddress)

			var nodes []vpn.NodeDetails

			if owner == "" {
				nodes, err = queryNodes(cliCtx, cdc)
				if err != nil {
					return err
				}
			} else {
				owner, err := csdkTypes.AccAddressFromBech32(owner)
				if err != nil {
					return err
				}

				nodes, err = queryNodesOfOwner(cliCtx, cdc, owner)
				if err != nil {
					return err
				}
			}

			if nodes == nil || len(nodes) == 0 {
				return fmt.Errorf("no nodes found")
			}

			nodesData, err := json.MarshalIndent(nodes, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(nodesData))

			return nil
		},
	}

	cmd.Flags().String(flagOwnerAddress, "", "Owner address")

	return cmd
}
