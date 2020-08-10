package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/client/common"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "node",
		Short: "Query a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := hub.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			node, err := common.QueryNode(ctx, address)
			if err != nil {
				return err
			}

			fmt.Println(node)
			return nil
		},
	}
}

func queryNodes(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "Query nodes",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			provider, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			plan, err := cmd.Flags().GetUint64(flagPlan)
			if err != nil {
				return err
			}

			page, err := cmd.Flags().GetInt(flagPage)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			var nodes types.Nodes

			if len(provider) > 0 {
				address, err := hub.ProvAddressFromBech32(provider)
				if err != nil {
					return err
				}

				nodes, err = common.QueryNodesForProvider(ctx, address, page, limit)
				if err != nil {
					return err
				}
			} else if plan > 0 {
				nodes, err = common.QueryNodesForPlan(ctx, plan, page, limit)
				if err != nil {
					return err
				}
			} else {
				nodes, err = common.QueryNodes(ctx, page, limit)
				if err != nil {
					return err
				}
			}

			for _, node := range nodes {
				fmt.Printf("%s\n\n", node)
			}

			return nil
		},
	}

	cmd.Flags().String(flagProvider, "", "provider address")
	cmd.Flags().Uint64(flagPlan, 0, "subscription plan ID")
	cmd.Flags().Int(flagPage, 1, "page")
	cmd.Flags().Int(flagLimit, 0, "limit")

	return cmd
}
