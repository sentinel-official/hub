package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	var (
		inactiveDuration = k.InactiveDuration(ctx)
	)

	for _, node := range state.Nodes {
		var (
			address  = node.GetAddress()
			provider = node.GetProvider()
		)

		k.SetNode(ctx, node)

		if node.Status.Equal(hubtypes.StatusActive) {
			k.SetActiveNode(ctx, address)
			if node.Provider != "" {
				k.SetActiveNodeForProvider(ctx, provider, address)
			}

			k.SetInactiveNodeAt(ctx, node.StatusAt.Add(inactiveDuration), address)
		} else {
			k.SetInactiveNode(ctx, address)
			if node.Provider != "" {
				k.SetInactiveNodeForProvider(ctx, provider, address)
			}
		}
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetNodes(ctx, 0, 0),
		k.GetParams(ctx),
	)
}
