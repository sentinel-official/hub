package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, plan := range state {
		k.SetPlan(ctx, plan)

		count := k.GetPlansCountForProvider(ctx, plan.Provider)
		k.SetPlansCountForProvider(ctx, plan.Provider, count+1)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return k.GetPlans(ctx)
}

func ValidateGenesis(state types.GenesisState) error {
	return nil
}
