package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, node := range state.Nodes {
		k.SetNode(ctx, node)

		if node.Status.Equal(hubtypes.StatusActive) {
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

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetNodes(ctx, 0, 0), k.GetParams(ctx))
}
