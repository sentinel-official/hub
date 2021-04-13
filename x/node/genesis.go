package node

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, node := range state.Nodes {
		k.SetNode(ctx, node)

		if node.Status.Equal(hub.StatusActive) {
			k.SetActiveNode(ctx, node.GetAddress())
			if node.Provider != "" {
				k.SetActiveNodeForProvider(ctx, node.GetProvider(), node.GetAddress())
			}

			k.SetInactiveNodeAt(ctx, node.StatusAt, node.GetAddress())
		} else {
			k.SetInactiveNode(ctx, node.GetAddress())
			if node.Provider != "" {
				k.SetInactiveNodeForProvider(ctx, node.GetProvider(), node.GetAddress())
			}
		}
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(k.GetNodes(ctx, 0, 0), k.GetParams(ctx))
}

func ValidateGenesis(state types.GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	for _, node := range state.Nodes {
		if err := node.Validate(); err != nil {
			return err
		}
	}

	nodes := make(map[string]bool)
	for _, node := range state.Nodes {
		if nodes[node.Address] {
			return fmt.Errorf("found duplicate node address %s", node.Address)
		}

		nodes[node.Address] = true
	}

	return nil
}
