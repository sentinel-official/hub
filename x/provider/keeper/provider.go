package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/provider/types"
)

func (k *Keeper) SetActiveProvider(ctx sdk.Context, v types.Provider) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveProviderKey(v.GetAddress())
		value = k.cdc.MustMarshal(&v)
	)

	store.Set(key, value)
}

func (k *Keeper) HasActiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActiveProviderKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) GetActiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) (v types.Provider, found bool) {
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

func (k *Keeper) DeleteActiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveProviderKey(addr)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactiveProvider(ctx sdk.Context, v types.Provider) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveProviderKey(v.GetAddress())
		value = k.cdc.MustMarshal(&v)
	)

	store.Set(key, value)
}

func (k *Keeper) HasInactiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactiveProviderKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) GetInactiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) (v types.Provider, found bool) {
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

func (k *Keeper) DeleteInactiveProvider(ctx sdk.Context, addr hubtypes.ProvAddress) {
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
		k.SetActiveProvider(ctx, provider)
	case hubtypes.StatusInactive:
		k.SetInactiveProvider(ctx, provider)
	default:
		panic(fmt.Errorf("failed to set the provider %v", provider))
	}
}

// HasProvider is for checking whether a provider with an address exists or not in the KVStore.
func (k *Keeper) HasProvider(ctx sdk.Context, addr hubtypes.ProvAddress) bool {
	return k.HasActiveProvider(ctx, addr) || k.HasInactiveProvider(ctx, addr)
}

// GetProvider is for getting a provider with an address from the KVStore.
func (k *Keeper) GetProvider(ctx sdk.Context, addr hubtypes.ProvAddress) (provider types.Provider, found bool) {
	provider, found = k.GetActiveProvider(ctx, addr)
	if found {
		return
	}

	provider, found = k.GetInactiveProvider(ctx, addr)
	if found {
		return
	}

	return provider, false
}

// GetProviders is for getting the providers from the KVStore.
func (k *Keeper) GetProviders(ctx sdk.Context) (items types.Providers) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Provider
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

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
