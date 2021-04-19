package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

// SetProvider is for inserting a provider into the KVStore.
func (k *Keeper) SetProvider(ctx sdk.Context, provider types.Provider) {
	key := types.ProviderKey(provider.GetAddress())
	value := k.cdc.MustMarshalBinaryBare(&provider)

	store := k.Store(ctx)
	store.Set(key, value)
}

// HasProvider is for checking whether a provider with an address exists or not in the KVStore.
func (k *Keeper) HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool {
	store := k.Store(ctx)

	key := types.ProviderKey(address)
	return store.Has(key)
}

// GetProvider is for getting a provider with an address from the KVStore.
func (k *Keeper) GetProvider(ctx sdk.Context, address hubtypes.ProvAddress) (provider types.Provider, found bool) {
	store := k.Store(ctx)

	key := types.ProviderKey(address)
	value := store.Get(key)
	if value == nil {
		return provider, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &provider)
	return provider, true
}

// GetProviders is for getting the providers from the KVStore.
func (k *Keeper) GetProviders(ctx sdk.Context, skip, limit int) (items types.Providers) {
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
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &item)
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
		var provider types.Provider
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &provider)

		if stop := fn(i, provider); stop {
			break
		}
		i++
	}
}
