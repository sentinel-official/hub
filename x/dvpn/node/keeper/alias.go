package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	plan "github.com/sentinel-official/hub/x/dvpn/plan/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func (k Keeper) GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider.Provider, bool) {
	return k.provider.GetProvider(ctx, address)
}

func (k Keeper) GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress) plan.Plans {
	return k.plan.GetPlansForProvider(ctx, address)
}

func (k Keeper) DeleteNodeAddressForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) {
	k.plan.DeleteNodeAddressForPlan(ctx, id, address)
}
