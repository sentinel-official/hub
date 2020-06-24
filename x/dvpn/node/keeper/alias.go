package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func (k Keeper) GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider.Provider, bool) {
	return k.provider.GetProvider(ctx, address)
}

func (k Keeper) GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress) subscription.Plans {
	return k.subscription.GetPlansForProvider(ctx, address)
}

func (k Keeper) DeleteNodeAddressForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) {
	k.subscription.DeleteNodeAddressForPlan(ctx, id, address)
}
