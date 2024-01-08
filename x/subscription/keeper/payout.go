package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
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

func (k *Keeper) SetPayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountByNodeKey(accAddr, nodeAddr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashPayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountByNodeKey(accAddr, nodeAddr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeletePayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountByNodeKey(accAddr, nodeAddr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetLatestPayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr hubtypes.NodeAddress) (payout types.Payout, found bool) {
	store := k.Store(ctx)

	iter := sdk.KVStoreReversePrefixIterator(store, types.GetPayoutForAccountByNodeKeyPrefix(accAddr, nodeAddr))
	defer iter.Close()

	if iter.Valid() {
		payout, found = k.GetPayout(ctx, types.IDFromPayoutForAccountByNodeKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("payout for account by node key %X does not exist", iter.Key()))
		}
	}

	return payout, found
}

func (k *Keeper) SetPayoutForNextAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNextAtKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeletePayoutForNextAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNextAtKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IteratePayoutsForNextAt(ctx sdk.Context, at time.Time, fn func(index int, item types.Payout) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.PayoutForNextAtKeyPrefix, sdk.PrefixEndBytes(types.GetPayoutForNextAtKeyPrefix(at)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		payout, found := k.GetPayout(ctx, types.IDFromPayoutForNextAtKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("payout for next_at key %X does not exist", iter.Key()))
		}

		if stop := fn(i, payout); stop {
			break
		}
		i++
	}
}
