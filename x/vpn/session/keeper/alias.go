package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	subscription "github.com/sentinel-official/hub/x/vpn/subscription/types"
)

func (k Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool {
	return k.plan.HasNodeForPlan(ctx, id, address)
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription subscription.Subscription) {
	k.subscription.SetSubscription(ctx, subscription)
}

func (k Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription.Subscription, bool) {
	return k.subscription.GetSubscription(ctx, id)
}

func (k Keeper) HasSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool {
	return k.subscription.HasSubscriptionForNode(ctx, address, id)
}

func (k Keeper) HasSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) bool {
	return k.subscription.HasSubscriptionForAddress(ctx, address, id)
}
