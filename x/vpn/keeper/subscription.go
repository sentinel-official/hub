package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetSubscriptionsCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionKey)
	store.Set(types.SubscriptionsCountKey, value)
}

func (k Keeper) GetSubscriptionsCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.subscriptionKey)

	value := store.Get(types.SubscriptionsCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	key := types.SubscriptionKey(subscription.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(subscription)

	store := ctx.KVStore(k.subscriptionKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscription(ctx sdk.Context, id hub.SubscriptionID) (subscription types.Subscription, found bool) {
	store := ctx.KVStore(k.subscriptionKey)

	key := types.SubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &subscription)
	return subscription, true
}

func (k Keeper) SetSubscriptionsCountOfNode(ctx sdk.Context, id hub.NodeID, count uint64) {
	key := types.SubscriptionsCountOfNodeKey(id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsCountOfNode(ctx sdk.Context, id hub.NodeID) (count uint64) {
	store := ctx.KVStore(k.subscriptionKey)

	key := types.SubscriptionsCountOfNodeKey(id)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscriptionIDByNodeID(ctx sdk.Context, i hub.NodeID, j uint64, id hub.SubscriptionID) {
	key := types.SubscriptionIDByNodeIDKey(i, j)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.subscriptionKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionIDByNodeID(ctx sdk.Context, i hub.NodeID, j uint64) (id hub.SubscriptionID, found bool) {
	store := ctx.KVStore(k.subscriptionKey)

	key := types.SubscriptionIDByNodeIDKey(i, j)
	value := store.Get(key)
	if value == nil {
		return hub.NewSubscriptionID(0), false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetSubscriptionsCountOfAddress(ctx sdk.Context, address sdk.AccAddress, count uint64) {
	key := types.SubscriptionsCountOfAddressKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsCountOfAddress(ctx sdk.Context, address sdk.AccAddress) (count uint64) {
	store := ctx.KVStore(k.subscriptionKey)

	key := types.SubscriptionsCountOfAddressKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscriptionIDByAddress(ctx sdk.Context, address sdk.AccAddress, i uint64, id hub.SubscriptionID) {
	key := types.SubscriptionIDByAddressKey(address, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.subscriptionKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionIDByAddress(ctx sdk.Context,
	address sdk.AccAddress, i uint64) (id hub.SubscriptionID, found bool) {
	store := ctx.KVStore(k.subscriptionKey)

	key := types.SubscriptionIDByAddressKey(address, i)
	value := store.Get(key)
	if value == nil {
		return hub.NewSubscriptionID(0), false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) GetSubscriptionsOfNode(ctx sdk.Context, id hub.NodeID) (subscriptions []types.Subscription) {
	count := k.GetSubscriptionsCountOfNode(ctx, id)

	subscriptions = make([]types.Subscription, 0, count)
	for i := uint64(0); i < count; i++ {
		_id, _ := k.GetSubscriptionIDByNodeID(ctx, id, i)

		subscription, _ := k.GetSubscription(ctx, _id)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

func (k Keeper) GetSubscriptionsOfAddress(ctx sdk.Context,
	address sdk.AccAddress) (subscriptions []types.Subscription) {
	count := k.GetSubscriptionsCountOfAddress(ctx, address)

	subscriptions = make([]types.Subscription, 0, count)
	for i := uint64(0); i < count; i++ {
		id, _ := k.GetSubscriptionIDByAddress(ctx, address, i)

		subscription, _ := k.GetSubscription(ctx, id)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

func (k Keeper) GetAllSubscriptions(ctx sdk.Context) (subscriptions []types.Subscription) {
	store := ctx.KVStore(k.subscriptionKey)

	iter := sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var subscription types.Subscription
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &subscription)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

func (k Keeper) IterateSubscriptions(ctx sdk.Context,
	fn func(index int64, subscription types.Subscription) (stop bool)) {
	store := ctx.KVStore(k.subscriptionKey)

	iterator := sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
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
