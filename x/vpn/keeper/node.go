package keeper

import (
	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetNodesCount(ctx csdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(types.NodesCountKey, value)
}

func (k Keeper) GetNodesCount(ctx csdk.Context) (count uint64) {
	store := ctx.KVStore(k.nodeStoreKey)

	value := store.Get(types.NodesCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetNode(ctx csdk.Context, node types.Node) {
	key := types.NodeKey(node.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(node)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetNode(ctx csdk.Context, id sdk.ID) (node types.Node, found bool) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodeKey(id)
	value := store.Get(key)
	if value == nil {
		return node, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &node)
	return node, true
}

func (k Keeper) SetNodesCountOfAddress(ctx csdk.Context, address csdk.AccAddress, count uint64) {
	key := types.NodesCountOfAddressKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetNodesCountOfAddress(ctx csdk.Context, address csdk.AccAddress) (count uint64) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodesCountOfAddressKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetNodeIDByAddress(ctx csdk.Context, address csdk.AccAddress, i uint64, id sdk.ID) {
	key := types.NodeIDByAddressKey(address, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetNodeIDByAddress(ctx csdk.Context, address csdk.AccAddress, i uint64) (id sdk.ID, found bool) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodeIDByAddressKey(address, i)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetActiveNodeIDs(ctx csdk.Context, height int64, ids sdk.IDs) {
	ids = ids.Sort()

	key := types.ActiveNodeIDsKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveNodeIDs(ctx csdk.Context, height int64) (ids sdk.IDs) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.ActiveNodeIDsKey(height)
	value := store.Get(key)
	if value == nil {
		return ids
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &ids)
	return ids
}

func (k Keeper) DeleteActiveNodeIDs(ctx csdk.Context, height int64) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.ActiveNodeIDsKey(height)
	store.Delete(key)
}

func (k Keeper) GetNodesOfAddress(ctx csdk.Context, address csdk.AccAddress) (nodes []types.Node) {
	count := k.GetNodesCountOfAddress(ctx, address)

	nodes = make([]types.Node, 0, count)
	for i := uint64(0); i < count; i++ {
		id, _ := k.GetNodeIDByAddress(ctx, address, i)

		node, _ := k.GetNode(ctx, id)
		nodes = append(nodes, node)
	}

	return nodes
}

func (k Keeper) GetAllNodes(ctx csdk.Context) (nodes []types.Node) {
	store := ctx.KVStore(k.nodeStoreKey)

	iter := csdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var node types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &node)
		nodes = append(nodes, node)
	}

	return nodes
}

// nolint: dupl
func (k Keeper) IterateNodes(ctx csdk.Context, fn func(index int64, node types.Node) (stop bool)) {
	store := ctx.KVStore(k.nodeStoreKey)

	iterator := csdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iterator.Close()

	for i := int64(0); iterator.Valid(); iterator.Next() {
		var node types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &node)

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}

func (k Keeper) AddNodeIDToActiveList(ctx csdk.Context, height int64, id sdk.ID) {
	ids := k.GetActiveNodeIDs(ctx, height)

	index := ids.Search(id)
	if index != len(ids) {
		return
	}

	ids = ids.Append(id)
	k.SetActiveNodeIDs(ctx, height, ids)
}

func (k Keeper) RemoveNodeIDFromActiveList(ctx csdk.Context, height int64, id sdk.ID) {
	ids := k.GetActiveNodeIDs(ctx, height)

	index := ids.Search(id)
	if index == len(ids) {
		return
	}

	ids = ids.Delete(index)
	k.SetActiveNodeIDs(ctx, height, ids)
}
