package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetResolver(ctx sdk.Context, resolver types.Resolver) {
	store := ctx.KVStore(k.resolverKey)

	key := types.ResolverKey(resolver.Owner)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(resolver)

	store.Set(key, value)
}

func (k Keeper) GetResolver(ctx sdk.Context, id sdk.AccAddress) (resolver types.Resolver, found bool) {
	store := ctx.KVStore(k.resolverKey)

	key := types.ResolverKey(id)

	value := store.Get(key)
	if value == nil {
		return resolver, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &resolver)
	return resolver, true
}

func (k Keeper) GetAllResolvers(ctx sdk.Context) (resolvers []types.Resolver) {
	store := ctx.KVStore(k.resolverKey)

	iter := sdk.KVStorePrefixIterator(store, types.ResolverKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var resolver types.Resolver
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &resolver)
		resolvers = append(resolvers, resolver)
	}

	return resolvers
}
