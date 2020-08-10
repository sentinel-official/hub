package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k Keeper) SetSubscriptionsCount(ctx sdk.Context, count uint64) {
	key := types.CountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsCount(ctx sdk.Context) (count uint64) {
	store := k.Store(ctx)

	key := types.CountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	key := types.SubscriptionKey(subscription.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(subscription)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription types.Subscription, found bool) {
	store := k.Store(ctx)

	key := types.SubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &subscription)
	return subscription, true
}

func (k Keeper) GetSubscriptions(ctx sdk.Context) (items types.Subscriptions) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Subscription
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.SubscriptionForAddressKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) HasSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) bool {
	key := types.SubscriptionForAddressKey(address, id)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k Keeper) DeleteSubscriptionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.SubscriptionForAddressKey(address, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) GetSubscriptionsForAddress(ctx sdk.Context, address sdk.AccAddress) (items types.Subscriptions) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetSubscriptionForAddressKeyPrefix(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &id)

		item, _ := k.GetSubscription(ctx, id)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetSubscriptionForPlan(ctx sdk.Context, plan, id uint64) {
	key := types.SubscriptionForPlanKey(plan, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) HasSubscriptionForPlan(ctx sdk.Context, plan, id uint64) bool {
	key := types.SubscriptionForPlanKey(plan, id)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k Keeper) GetSubscriptionsForPlan(ctx sdk.Context, plan uint64) (items types.Subscriptions) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, sdk.Uint64ToBigEndian(plan))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &id)

		item, _ := k.GetSubscription(ctx, id)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) {
	key := types.SubscriptionForNodeKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) HasSubscriptionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) bool {
	key := types.SubscriptionForNodeKey(address, id)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k Keeper) GetSubscriptionsForNode(ctx sdk.Context, address hub.NodeAddress) (items types.Subscriptions) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetSubscriptionForNodeKeyPrefix(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &id)

		item, _ := k.GetSubscription(ctx, id)
		items = append(items, item)
	}

	return items
}
