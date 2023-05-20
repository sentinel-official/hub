package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) SetPayout(ctx sdk.Context, payout types.Payout) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(payout.ID)
		value = k.cdc.MustMarshal(&payout)
	)

	store.Set(key, value)
}

func (k *Keeper) HasPayout(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetPayout(ctx sdk.Context, id uint64) (payout types.Payout, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return payout, false
	}

	k.cdc.MustUnmarshal(value, &payout)
	return payout, true
}

func (k *Keeper) DeletePayout(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetPayouts(ctx sdk.Context) (items types.Payouts) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.PayoutKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Payout
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IteratePayouts(ctx sdk.Context, fn func(index int, item types.Payout) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.PayoutKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item types.Payout
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetPayoutForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashPayoutForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeletePayoutForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetPayoutForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashPayoutForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNodeKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeletePayoutForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetPayoutForTimestamp(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForTimestampKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeletePayoutForTimestamp(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForTimestampKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IteratePayoutsForTimestamp(ctx sdk.Context, at time.Time, fn func(index int, item types.Payout) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.PayoutForTimestampKeyPrefix, sdk.PrefixEndBytes(types.GetPayoutForTimestampKeyPrefix(at)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		payout, found := k.GetPayout(ctx, types.IDFromPayoutForTimestampKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("payout for timestamp key %X does not exist", iter.Key()))
		}

		if stop := fn(i, payout); stop {
			break
		}
		i++
	}
}
