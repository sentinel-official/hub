package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetSubscriptionsCount(ctx csdkTypes.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(types.SubscriptionsCountKey, value)
}

func (k Keeper) GetSubscriptionsCount(ctx csdkTypes.Context) (count uint64) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	value := store.Get(types.SubscriptionsCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscription(ctx csdkTypes.Context, subscription types.Subscription) {
	key := types.SubscriptionKey(subscription.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(subscription)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscription(ctx csdkTypes.Context, id sdkTypes.ID) (subscription types.Subscription, found bool) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &subscription)
	return subscription, true
}

func (k Keeper) SetSubscriptionIDByNodeID(ctx csdkTypes.Context, i sdkTypes.ID, j uint64, id sdkTypes.ID) {
	key := types.SubscriptionIDByNodeIDKey(i, j)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionIDByNodeID(ctx csdkTypes.Context, i sdkTypes.ID, j uint64) (id sdkTypes.ID, found bool) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionIDByNodeIDKey(i, j)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetSubscriptionsCountOfAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress, count uint64) {
	key := types.SubscriptionsCountOfAddressKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsCountOfAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress) (count uint64) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionsCountOfAddressKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscriptionIDByAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress, i uint64, id sdkTypes.ID) {
	key := types.SubscriptionIDByAddressKey(address, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.subscriptionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionIDByAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress, i uint64) (id sdkTypes.ID, found bool) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	key := types.SubscriptionIDByAddressKey(address, i)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) GetAllSubscriptions(ctx csdkTypes.Context) (subscriptions []types.Subscription) {
	store := ctx.KVStore(k.subscriptionStoreKey)

	iter := csdkTypes.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var subscription types.Subscription
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &subscription)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

// nolint
func (k Keeper) IterateSubscriptions(ctx csdkTypes.Context,
	fn func(index int64, subscription types.Subscription) (stop bool)) {

	store := ctx.KVStore(k.subscriptionStoreKey)

	iterator := csdkTypes.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
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

func (k Keeper) AddSubscription(ctx csdkTypes.Context, node types.Node,
	subscription types.Subscription) (allTags csdkTypes.Tags, err csdkTypes.Error) {

	allTags = csdkTypes.EmptyTags()

	subscription.ClientPubKey, err = k.accountKeeper.GetPubKey(ctx, subscription.Client)
	if err != nil {
		return nil, err
	}

	tags, err := k.AddDeposit(ctx, subscription.Client, subscription.TotalDeposit)
	if err != nil {
		return nil, err
	}

	allTags = allTags.AppendTags(tags)

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionIDByNodeID(ctx, node.ID, node.SubscriptionsCount, subscription.ID)

	count := k.GetSubscriptionsCountOfAddress(ctx, subscription.Client)
	k.SetSubscriptionIDByAddress(ctx, subscription.Client, count, subscription.ID)

	node.SubscriptionsCount++
	k.SetNode(ctx, node)
	k.SetSubscriptionsCountOfAddress(ctx, subscription.Client, count+1)
	k.SetSubscriptionsCount(ctx, subscription.ID.UInt64()+1)

	return allTags, nil
}
