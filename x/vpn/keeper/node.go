package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetNodesCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(types.NodesCountKey, value)
}

func (k Keeper) GetNodesCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.nodeKey)
	
	value := store.Get(types.NodesCountKey)
	if value == nil {
		return 0
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetNode(ctx sdk.Context, node types.Node) {
	key := types.NodeKey(node.ID)
	
	value := k.cdc.MustMarshalBinaryLengthPrefixed(node)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) GetNode(ctx sdk.Context, id hub.NodeID) (node types.Node, found bool) {
	store := ctx.KVStore(k.nodeKey)
	
	key := types.NodeKey(id)
	value := store.Get(key)
	if value == nil {
		return node, false
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &node)
	return node, true
}

func (k Keeper) SetNodesCountOfAddress(ctx sdk.Context, address sdk.AccAddress, count uint64) {
	key := types.NodesCountOfAddressKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) GetNodesCountOfAddress(ctx sdk.Context, address sdk.AccAddress) (count uint64) {
	store := ctx.KVStore(k.nodeKey)
	
	key := types.NodesCountOfAddressKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetNodeIDByAddress(ctx sdk.Context, address sdk.AccAddress, i uint64, id hub.NodeID) {
	key := types.NodeIDByAddressKey(address, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) GetNodeIDByAddress(ctx sdk.Context, address sdk.AccAddress, i uint64) (id hub.NodeID, found bool) {
	store := ctx.KVStore(k.nodeKey)
	
	key := types.NodeIDByAddressKey(address, i)
	value := store.Get(key)
	if value == nil {
		return hub.NewNodeID(0), false
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetActiveNodeIDs(ctx sdk.Context, height int64, ids hub.IDs) {
	ids = ids.Sort()
	
	key := types.ActiveNodeIDsKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveNodeIDs(ctx sdk.Context, height int64) (ids hub.IDs) {
	store := ctx.KVStore(k.nodeKey)
	
	key := types.ActiveNodeIDsKey(height)
	value := store.Get(key)
	if value == nil {
		return ids
	}
	
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &ids)
	return ids
}

func (k Keeper) DeleteActiveNodeIDs(ctx sdk.Context, height int64) {
	store := ctx.KVStore(k.nodeKey)
	
	key := types.ActiveNodeIDsKey(height)
	store.Delete(key)
}

func (k Keeper) GetNodesOfAddress(ctx sdk.Context, address sdk.AccAddress) (nodes []types.Node) {
	count := k.GetNodesCountOfAddress(ctx, address)
	
	nodes = make([]types.Node, 0, count)
	for i := uint64(0); i < count; i++ {
		id, _ := k.GetNodeIDByAddress(ctx, address, i)
		
		node, _ := k.GetNode(ctx, id)
		nodes = append(nodes, node)
	}
	
	return nodes
}

func (k Keeper) GetAllNodes(ctx sdk.Context) (nodes []types.Node) {
	store := ctx.KVStore(k.nodeKey)
	
	iter := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()
	
	for ; iter.Valid(); iter.Next() {
		var node types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &node)
		nodes = append(nodes, node)
	}
	
	return nodes
}

func (k Keeper) AddNodeIDToActiveList(ctx sdk.Context, height int64, id hub.NodeID) {
	ids := k.GetActiveNodeIDs(ctx, height)
	
	index := ids.Search(id)
	if index != len(ids) {
		return
	}
	
	ids = ids.Append(id)
	k.SetActiveNodeIDs(ctx, height, ids)
}

func (k Keeper) RemoveNodeIDFromActiveList(ctx sdk.Context, height int64, id hub.NodeID) {
	ids := k.GetActiveNodeIDs(ctx, height)
	
	index := ids.Search(id)
	if index == len(ids) {
		return
	}
	
	ids = ids.Delete(index)
	k.SetActiveNodeIDs(ctx, height, ids)
}
