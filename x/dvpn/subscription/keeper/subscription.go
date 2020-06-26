package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func (k Keeper) SetSubscriptionsCount(ctx sdk.Context, count uint64) {
	key := types.SubscriptionsCountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.SubscriptionStore(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsCount(ctx sdk.Context) (count uint64) {
	store := k.SubscriptionStore(ctx)

	key := types.PlansCountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	key := types.PlanKey(subscription.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(subscription)

	store := k.SubscriptionStore(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription types.Subscription, found bool) {
	store := k.SubscriptionStore(ctx)

	key := types.SubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &subscription)
	return subscription, true
}

func (k Keeper) GetSubscriptions(ctx sdk.Context) (subscriptions types.Subscriptions) {
	store := k.SubscriptionStore(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var subscription types.Subscription
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &subscription)
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

func (k Keeper) SetSubscriptionIDForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.SubscriptionIDForAddressKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.SubscriptionStore(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSubscriptionsForAddress(ctx sdk.Context, address sdk.AccAddress) (subscriptions types.Subscriptions) {
	store := k.SubscriptionStore(ctx)

	iter := sdk.KVStorePrefixIterator(store, address.Bytes())
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &id)

		s, _ := k.GetSubscription(ctx, id)
		subscriptions = append(subscriptions, s)
	}

	return subscriptions
}
