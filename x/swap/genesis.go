package swap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/swap/keeper"
	"github.com/sentinel-official/hub/v12/x/swap/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)
	for _, item := range state.Swaps {
		k.SetSwap(ctx, item)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetSwaps(ctx),
		k.GetParams(ctx),
	)
}
