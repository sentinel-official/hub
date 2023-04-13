package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	inactiveDuration := k.InactiveDuration(ctx)
	for _, item := range state.Nodes {
		k.SetNode(ctx, item)
		if item.Status.Equal(hubtypes.StatusActive) {
			var (
				addr     = item.GetAddress()
				statusAt = item.StatusAt.Add(inactiveDuration)
			)

			k.SetInactiveNodeAt(ctx, statusAt, addr)
		}
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetNodes(ctx),
		k.GetParams(ctx),
	)
}
