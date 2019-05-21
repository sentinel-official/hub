package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

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

func (k Keeper) AddSubscription(ctx csdkTypes.Context, subscription types.Subscription) (allTags csdkTypes.Tags, err csdkTypes.Error) {
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
	allTags = allTags.AppendTag(types.TagSubscriptionID, subscription.ID.String())

	return allTags, nil
}
