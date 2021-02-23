package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	plan "github.com/sentinel-official/hub/x/plan/types"
)

func (k Keeper) FundCommunityPool(ctx sdk.Context, from sdk.AccAddress, coin sdk.Coin) error {
	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), from)
}

func (k Keeper) HasProvider(ctx sdk.Context, address hub.ProvAddress) bool {
	if address == nil {
		return true
	}

	return k.provider.HasProvider(ctx, address)
}

func (k Keeper) GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress) plan.Plans {
	return k.plan.GetPlansForProvider(ctx, address, 0, 0)
}

func (k Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) {
	k.plan.DeleteNodeForPlan(ctx, id, address)
}
