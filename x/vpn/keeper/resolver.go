package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetResolverCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)
	
	store := ctx.KVStore(k.resolverKey)
	store.Set(types.ResolverCountKey, value)
}

func (k Keeper) GetResolverCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.resolverKey)
	
	value := store.Get(types.ResolverCountKey)
	if value == nil {
		return 0
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetResolver(ctx sdk.Context, resolver types.Resolver) {
	store := ctx.KVStore(k.resolverKey)
	
	key := types.ResolverKey(resolver.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(resolver)
	
	store.Set(key, value)
}

func (k Keeper) GetResolver(ctx sdk.Context, id hub.ResolverID) (resolver types.Resolver, found bool) {
	store := ctx.KVStore(k.resolverKey)
	
	key := types.ResolverKey(id)
	
	value := store.Get(key)
	if value == nil {
		return resolver, false
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &resolver)
	return resolver, true
}

func (k Keeper) SetResolverCountOfAddress(ctx sdk.Context, address sdk.AccAddress, count uint64) {
	key := types.ResolversCountOfAddressKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)
	
	store := ctx.KVStore(k.resolverKey)
	store.Set(key, value)
}

func (k Keeper) GetResolversCountOfAddress(ctx sdk.Context, address sdk.AccAddress) (count uint64) {
	store := ctx.KVStore(k.resolverKey)
	
	key := types.ResolversCountOfAddressKey(address)
	
	value := store.Get(key)
	if value == nil {
		return 0
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetResolverIDByAddress(ctx sdk.Context, address sdk.AccAddress, i uint64, id hub.ResolverID) {
	key := types.ResolverIDByAddressKey(address, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)
	
	store := ctx.KVStore(k.resolverKey)
	store.Set(key, value)
}

func (k Keeper) GetResolverIDByAddress(ctx sdk.Context, address sdk.AccAddress,
	i uint64) (id hub.ResolverID, found bool) {
	key := types.ResolverIDByAddressKey(address, i)
	
	store := ctx.KVStore(k.resolverKey)
	value := store.Get(key)
	if value == nil {
		return hub.NewResolverID(0), false
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) GetResolversOfAddress(ctx sdk.Context, address sdk.AccAddress) (resolvers types.Resolvers) {
	count := k.GetResolversCountOfAddress(ctx, address)
	
	for i := uint64(0); i < count; i++ {
		id, _ := k.GetResolverIDByAddress(ctx, address, i)
		
		resolver, _ := k.GetResolver(ctx, id)
		resolvers = append(resolvers, resolver)
	}
	return resolvers
}

func (k Keeper) GetAllResolvers(ctx sdk.Context) (resolvers types.Resolvers) {
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

func (k Keeper) SetResolverOfNode(ctx sdk.Context, nodeID hub.NodeID, resolver hub.ResolverID) {
	key := types.ResolverOfNodeKey(nodeID, resolver)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(resolver)
	
	store := ctx.KVStore(k.resolverKey)
	store.Set(key, value)
}

func (k Keeper) GetNodeOfResolver(ctx sdk.Context, resolverID hub.ResolverID,
	nodeID hub.NodeID) (id hub.NodeID, found bool) {
	key := types.NodeOfResolverKey(resolverID, nodeID)
	
	store := ctx.KVStore(k.resolverKey)
	value := store.Get(key)
	
	if value == nil {
		return nil, false
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	
	return id, true
}

func (k Keeper) GetNodesOfResolver(ctx sdk.Context, resolverID hub.ResolverID) (nodeIDs []hub.NodeID) {
	store := ctx.KVStore(k.resolverKey)
	key := types.NodeOfResolverKey(resolverID, nil)
	
	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()
	
	for ; iter.Valid(); iter.Next() {
		var _nodeID hub.NodeID
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &_nodeID)
		nodeIDs = append(nodeIDs, _nodeID)
	}
	
	return nodeIDs
}

func (k Keeper) SetNodeOfResolver(ctx sdk.Context, resolverID hub.ResolverID, nodeID hub.NodeID) {
	key := types.NodeOfResolverKey(resolverID, nodeID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(nodeID)
	
	store := ctx.KVStore(k.resolverKey)
	store.Set(key, value)
}

func (k Keeper) GetResolverOfNode(ctx sdk.Context, nodeID hub.NodeID,
	resolverID hub.ResolverID) (id hub.ResolverID, found bool) {
	key := types.ResolverOfNodeKey(nodeID, resolverID)
	
	store := ctx.KVStore(k.resolverKey)
	value := store.Get(key)
	
	if value == nil {
		return nil, false
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) GetResolversOfNode(ctx sdk.Context, nodeID hub.NodeID) (resolvers []hub.ResolverID) {
	store := ctx.KVStore(k.resolverKey)
	key := types.ResolverOfNodeKey(nodeID, nil)
	
	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()
	
	for ; iter.Valid(); iter.Next() {
		var resolver hub.ResolverID
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &resolver)
		resolvers = append(resolvers, resolver)
	}
	
	return resolvers
}

func (k Keeper) RemoveVPNNodeOnResolver(ctx sdk.Context, nodeID hub.NodeID, resolverID hub.ResolverID) {
	nodeKey := types.ResolverOfNodeKey(nodeID, resolverID)
	resolverKey := types.NodeOfResolverKey(resolverID, nodeID)
	
	store := ctx.KVStore(k.resolverKey)
	store.Delete(nodeKey)
	store.Delete(resolverKey)
}
