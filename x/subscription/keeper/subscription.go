package keeper

import (
	"fmt"
	"time"

	hubtypes "github.com/sentinel-official/hub/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(subscription.GetID())
	)

	value, err := k.cdc.MarshalInterface(subscription)
	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}

func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription types.Subscription, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return subscription, false
	}
	if err := k.cdc.UnmarshalInterface(value, &subscription); err != nil {
		panic(err)
	}

	return subscription, true
}

func (k *Keeper) DeleteSubscription(ctx sdk.Context, id uint64) {
	key := types.SubscriptionKey(id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetSubscriptions(ctx sdk.Context) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Subscription
		if err := k.cdc.UnmarshalInterface(iter.Value(), &item); err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForAccountKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForAccountKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for account key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForNode(ctx sdk.Context, addr hubtypes.NodeAddress) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForNodeKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForNodeKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for node key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForPlan(ctx sdk.Context, id uint64) (items types.Subscriptions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForPlanKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForPlanKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("subscription for plan key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionExpiryAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SubscriptionExpiryAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteSubscriptionExpiryAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SubscriptionExpiryAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateSubscriptionExpirys(ctx sdk.Context, end time.Time, fn func(index int, item types.Subscription) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.SubscriptionExpiryAtKeyPrefix, sdk.PrefixEndBytes(types.GetSubscriptionExpiryAtKeyPrefix(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var (
			key             = iter.Key()
			subscription, _ = k.GetSubscription(ctx, types.IDFromSubscriptionExpiryAtKey(key))
		)

		if stop := fn(i, subscription); stop {
			break
		}
		i++
	}
}

func (k *Keeper) CreateNodeSubscription(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, hours int64, price sdk.Coin) (uint64, error) {
	return 0, nil
}

func (k *Keeper) CreatePlanSubscription(ctx sdk.Context, accAddr sdk.AccAddress, id uint64, price sdk.Coin) (uint64, error) {
	return 0, nil
}
