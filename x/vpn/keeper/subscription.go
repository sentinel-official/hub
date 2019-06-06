package keeper

import (
	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetSubscriptionsCount(ctx csdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(types.SubscriptionsCountKey, value)
}

func (k Keeper) GetSubscriptionsCount(ctx csdk.Context) (count uint64) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	value := store.Get(types.SubscriptionsCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscription(ctx csdk.Context, subscription types.Subscription) {
	key := types.SubscriptionKey(subscription.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(subscription)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscription(ctx csdk.Context, id sdk.ID) (subscription types.Subscription, found bool) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &subscription)
	return subscription, true
}

func (k Keeper) SetSubscriptionsCountOfNode(ctx csdk.Context, id sdk.ID, count uint64) {
	key := types.SubscriptionsCountOfNodeKey(id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsCountOfNode(ctx csdk.Context, id sdk.ID) (count uint64) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionsCountOfNodeKey(id)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscriptionIDByNodeID(ctx csdk.Context, i sdk.ID, j uint64, id sdk.ID) {
	key := types.SubscriptionIDByNodeIDKey(i, j)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionIDByNodeID(ctx csdk.Context, i sdk.ID, j uint64) (id sdk.ID, found bool) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionIDByNodeIDKey(i, j)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetSubscriptionsCountOfAddress(ctx csdk.Context, address csdk.AccAddress, count uint64) {
	key := types.SubscriptionsCountOfAddressKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsCountOfAddress(ctx csdk.Context, address csdk.AccAddress) (count uint64) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionsCountOfAddressKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscriptionIDByAddress(ctx csdk.Context, address csdk.AccAddress, i uint64, id sdk.ID) {
	key := types.SubscriptionIDByAddressKey(address, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionIDByAddress(ctx csdk.Context, address csdk.AccAddress, i uint64) (id sdk.ID, found bool) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionIDByAddressKey(address, i)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) GetSubscriptionsOfNode(ctx csdk.Context, id sdk.ID) (subscriptions []types.Subscription) {
	count := k.GetSubscriptionsCountOfNode(ctx, id)

	subscriptions = make([]types.Subscription, 0, count)
	for i := uint64(0); i < count; i++ {
		_id, _ := k.GetSubscriptionIDByNodeID(ctx, id, i)

		subscription, _ := k.GetSubscription(ctx, _id)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

func (k Keeper) GetSubscriptionsOfAddress(ctx csdk.Context, address csdk.AccAddress) (subscriptions []types.Subscription) {
	count := k.GetSubscriptionsCountOfAddress(ctx, address)

	subscriptions = make([]types.Subscription, 0, count)
	for i := uint64(0); i < count; i++ {
		id, _ := k.GetSubscriptionIDByAddress(ctx, address, i)

		subscription, _ := k.GetSubscription(ctx, id)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

func (k Keeper) GetAllSubscriptions(ctx csdk.Context) (subscriptions []types.Subscription) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	iter := csdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var subscription types.Subscription
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &subscription)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

// nolint
func (k Keeper) IterateSubscriptions(ctx csdk.Context,
	fn func(index int64, subscription types.Subscription) (stop bool)) {

	store := ctx.KVStore(k.subscriptionStoreKey)

	iterator := csdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	defer iterator.Close()

	for i := int64(0); iterator.Valid(); iterator.Next() {
		var subscription types.Subscription
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &subscription)

		if stop := fn(i, subscription); stop {
			break
		}
		i++
	}
}
