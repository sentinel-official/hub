package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v1/x/mint/keeper"
	"github.com/sentinel-official/hub/v1/x/mint/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	for _, item := range state.Inflations {
		k.SetInflation(ctx, item)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetInflations(ctx),
	)
}
