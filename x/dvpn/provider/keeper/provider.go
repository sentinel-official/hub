package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func (k Keeper) SetProvidersCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.Store(ctx)
	store.Set(types.ProvidersCountKey, value)
}

func (k Keeper) GetProvidersCount(ctx sdk.Context) (count uint64) {
	store := k.Store(ctx)

	value := store.Get(types.ProvidersCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetProvider(ctx sdk.Context, provider types.Provider) {
	key := types.ProviderKey(provider.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(provider)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetProvider(ctx sdk.Context, id hub.ProviderID) (provider types.Provider, found bool) {
	store := k.Store(ctx)

	key := types.ProviderKey(id)
	value := store.Get(key)
	if value == nil {
		return provider, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &provider)
	return provider, true
}

func (k Keeper) GetProviders(ctx sdk.Context) (providers types.Providers) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var provider types.Provider
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &provider)
		providers = append(providers, provider)
	}

	return providers
}

func (k Keeper) IterateProviders(ctx sdk.Context, f func(i int, provider types.Provider) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.ProviderKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var provider types.Provider
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &provider)

		if stop := f(i, provider); stop {
			break
		}
		i++
	}
}

func (k Keeper) SetProviderIDForAddress(ctx sdk.Context, address sdk.AccAddress, id hub.ProviderID) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(types.ProviderIDForAddressKey(address), value)
}

func (k Keeper) GetProviderIDForAddress(ctx sdk.Context, address sdk.AccAddress) (id hub.ProviderID, found bool) {
	store := k.Store(ctx)

	value := store.Get(types.ProviderIDForAddressKey(address))
	if value == nil {
		return nil, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}
