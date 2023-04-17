package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) SetLease(ctx sdk.Context, lease types.Lease) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(lease.ID)
		value = k.cdc.MustMarshal(&lease)
	)

	store.Set(key, value)
}

func (k *Keeper) HasLease(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetLease(ctx sdk.Context, id uint64) (lease types.Lease, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return lease, false
	}

	k.cdc.MustUnmarshal(value, &lease)
	return lease, true
}

func (k *Keeper) DeleteLease(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetLeases(ctx sdk.Context) (items types.Leases) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Lease
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateLeases(ctx sdk.Context, fn func(index int, item types.Lease) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item types.Lease
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetLeaseForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashLeaseForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForAccountKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteLeaseForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetLeaseForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashLeaseForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForNodeKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteLeaseForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForNodeKey(addr, id)
	)

	store.Delete(key)
}
