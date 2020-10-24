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
			k.SetActiveNode(ctx, node.Address)
			if node.Provider != nil {
				k.SetActiveNodeForProvider(ctx, node.Provider, node.Address)
			}

			k.SetInActiveNodeAt(ctx, node.StatusAt, node.Address)
		} else {
			k.SetInActiveNode(ctx, node.Address)
			if node.Provider != nil {
				k.SetInActiveNodeForProvider(ctx, node.Provider, node.Address)
			}
		}
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.NewGenesisState(k.GetNodes(ctx), k.GetParams(ctx))
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
		address := node.Address.String()
		if nodes[address] {
			return fmt.Errorf("found duplicate node address %s", address)
		}

		nodes[address] = true
	}

	return nil
}
