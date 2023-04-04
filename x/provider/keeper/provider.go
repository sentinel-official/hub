package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

func (k *Keeper) setActiveProvider(ctx sdk.Context, v types.Provider) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveProviderKey(v.GetAddress())
		value = k.cdc.MustMarshal(&v)
	)

	store.Set(key, value)
}

func (k *Keeper) hasActiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActiveProviderKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) getActiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) (v types.Provider, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveProviderKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) deleteActiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveProviderKey(addr)
	)

	store.Delete(key)
}

func (k *Keeper) setInactiveProvider(ctx sdk.Context, v types.Provider) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveProviderKey(v.GetAddress())
		value = k.cdc.MustMarshal(&v)
	)

	store.Set(key, value)
}

func (k *Keeper) hasInactiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactiveProviderKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) getInactiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) (v types.Provider, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveProviderKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) deleteInactiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveProviderKey(addr)
	)

	store.Delete(key)
}

// SetProvider is for inserting a provider into the KVStore.
func (k *Keeper) SetProvider(ctx sdk.Context, provider types.Provider) {
	switch provider.Status {
	case hubtypes.StatusActive:
		k.setActiveProvider(ctx, provider)
	case hubtypes.StatusInactive:
		k.setInactiveProvider(ctx, provider)
	default:
		panic(fmt.Errorf("invalid status for the provider %v", provider))
	}
}

// HasProvider is for checking whether a provider with an address exists or not in the KVStore.
func (k *Keeper) HasProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	return k.hasActiveProvider(ctx, addr) ||
		k.hasInactiveProvider(ctx, addr)
}

// GetProvider is for getting a provider with an address from the KVStore.
func (k *Keeper) GetProvider(ctx sdk.Context, addr hubtypes.ProvAddress) (provider types.Provider, found bool) {
	provider, found = k.getActiveProvider(ctx, addr)
	if found {
		return
	}

	provider, found = k.getInactiveProvider(ctx, addr)
	if found {
		return
	}

	return provider, false
}

// GetProviders is for getting the providers from the KVStore.
func (k *Keeper) GetProviders(ctx sdk.Context, skip, limit int64) (items types.Providers) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Provider
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

// IterateProviders is for iterating over the providers to perform an action.
func (k *Keeper) IterateProviders(ctx sdk.Context, fn func(index int, item types.Provider) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item types.Provider
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}
