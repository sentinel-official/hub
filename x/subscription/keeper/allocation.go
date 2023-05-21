package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) SetAllocation(ctx sdk.Context, id uint64, allocation types.Allocation) {
	// TODO: add field ID for allocation?

	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, allocation.GetAddress())
		value = k.cdc.MustMarshal(&allocation)
	)

	store.Set(key, value)
}

func (k *Keeper) GetAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) (allocation types.Allocation, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
		value = store.Get(key)
	)

	if value == nil {
		return allocation, false
	}

	k.cdc.MustUnmarshal(value, &allocation)
	return allocation, true
}

func (k *Keeper) HasAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
	)

	store.Delete(key)
}

func (k *Keeper) GetAllocations(ctx sdk.Context, id uint64) (items types.Allocations) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetAllocationKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Allocation
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateAllocations(ctx sdk.Context, id uint64, fn func(index int, item types.Allocation) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetAllocationKeyPrefix(id))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var allocation types.Allocation
		k.cdc.MustUnmarshal(iter.Value(), &allocation)

		if stop := fn(i, allocation); stop {
			break
		}
		i++
	}
}
