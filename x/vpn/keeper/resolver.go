package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
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

func (k Keeper) SetResolverOfNode(ctx sdk.Context, nodeID hub.NodeID, resolver sdk.AccAddress) {
	key := types.ResolverOfNodeKey(nodeID, resolver)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(resolver)

	store := ctx.KVStore(k.resolverKey)
	store.Set(key, value)
}

func (k Keeper) GetNodeOfResolver(ctx sdk.Context, resolver sdk.AccAddress, nodeID hub.NodeID) hub.NodeID {
	key := types.NodeOfResolverKey(resolver, nodeID)

	store := ctx.KVStore(k.resolverKey)
	value := store.Get(key)

	var _nodeID hub.NodeID
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &_nodeID)

	return _nodeID
}

func (k Keeper) GetNodesOfResolver(ctx sdk.Context, resolver sdk.AccAddress) (nodeIDs []hub.NodeID) {
	store := ctx.KVStore(k.resolverKey)
	key := types.NodeOfResolverKey(resolver, nil)

	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var _nodeID hub.NodeID
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &_nodeID)
		nodeIDs = append(nodeIDs, _nodeID)
	}

	return nodeIDs
}

func (k Keeper) SetNodeOfResolver(ctx sdk.Context, resolver sdk.AccAddress, nodeID hub.NodeID) {
	key := types.NodeOfResolverKey(resolver, nodeID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(nodeID)

	store := ctx.KVStore(k.resolverKey)
	store.Set(key, value)
}

func (k Keeper) GetResolverOfNode(ctx sdk.Context, nodeID hub.NodeID, resolver sdk.AccAddress) sdk.AccAddress {
	key := types.ResolverOfNodeKey(nodeID, resolver)

	store := ctx.KVStore(k.resolverKey)
	value := store.Get(key)

	var _resolver sdk.AccAddress
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &_resolver)

	return _resolver
}

func (k Keeper) GetResolversOfNode(ctx sdk.Context, nodeID hub.NodeID) (resolvers []sdk.AccAddress) {
	store := ctx.KVStore(k.resolverKey)
	key := types.ResolverOfNodeKey(nodeID, nil)

	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var resolver sdk.AccAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &resolver)
		resolvers = append(resolvers, resolver)
	}

	return resolvers
}

func (k Keeper) RemoveVPNNodeOnResolver(ctx sdk.Context, nodeID hub.NodeID, resolver sdk.AccAddress) {
	nodeKey := types.ResolverOfNodeKey(nodeID, resolver)
	resolverKey := types.NodeOfResolverKey(resolver, nodeID)

	store := ctx.KVStore(k.resolverKey)
	store.Delete(nodeKey)
	store.Delete(resolverKey)
}
