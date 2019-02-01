package cli

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
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

func QueryNodesCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Get details of nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			ownerAddress := viper.GetString(flagOwner)

			var paramBytes []byte
			var querier string
			var err error

			if ownerAddress == "" {
				querier = vpn.QueryNodes
			} else {
				querier = vpn.QueryNodesOfOwner
				accAddress, err := types.AccAddressFromBech32(ownerAddress)
				if err != nil {
					return err
				}
				params := vpn.NewQueryNodesOfOwnerParams(accAddress)
				paramBytes, err = cdc.MarshalJSON(params)
				if err != nil {
					return err
				}
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, querier), paramBytes)
			if err != nil {
				return err
			}

			var nodes []vpn.NodeDetails
			if err = cdc.UnmarshalJSON(res, &nodes); err != nil {
				return err
			}

			nodesData, err := json.MarshalIndent(nodes, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(nodesData))

			return nil
		},
	}

	cmd.Flags().String(flagOwner, "", "Owner address")

	return cmd
}
