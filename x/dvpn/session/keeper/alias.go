package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func (k Keeper) SetSubscription(ctx sdk.Context, subscription subscription.Subscription) {
	k.subscription.SetSubscription(ctx, subscription)
}

func (k Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription.Subscription, bool) {
	return k.subscription.GetSubscription(ctx, id)
}

func (k Keeper) HasNodeAddressForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool {
	return k.subscription.HasNodeAddressForPlan(ctx, id, address)
}

func (k Keeper) HasSubscriptionIDForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool {
	return k.subscription.HasSubscriptionIDForNode(ctx, address, id)
}

func (k Keeper) HasAddressForSubscriptionID(ctx sdk.Context, id uint64, address sdk.AccAddress) bool {
	return k.subscription.HasAddressForSubscriptionID(ctx, id, address)
}
