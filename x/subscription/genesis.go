package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var (
		subscriptions = k.GetSubscriptions(ctx)
		items         = make(types.GenesisSubscriptions, 0, len(subscriptions))
	)

	return types.NewGenesisState(items, k.GetParams(ctx))
}
