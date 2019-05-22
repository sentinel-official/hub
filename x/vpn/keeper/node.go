package keeper

import (
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetNodesCount(ctx csdkTypes.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(types.NodesCountKey, value)
}

func (k Keeper) GetNodesCount(ctx csdkTypes.Context) (count uint64) {
	store := ctx.KVStore(k.nodeStoreKey)

	value := store.Get(types.NodesCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetNode(ctx csdkTypes.Context, node types.Node) {
	key := types.NodeKey(node.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(node)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetNode(ctx csdkTypes.Context, id uint64) (node types.Node, found bool) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodeKey(id)
	value := store.Get(key)
	if value == nil {
		return node, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &node)
	return node, true
}

func (k Keeper) SetNodesCountOfAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress, count uint64) {
	key := types.NodesCountOfAddressKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetNodesCountOfAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress) (count uint64) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodesCountOfAddressKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetNodeIDByAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress, i, id uint64) {
	key := types.NodeIDByAddressKey(address, i)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetNodeIDByAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress, i uint64) (id uint64, found bool) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodeIDByAddressKey(address, i)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetActiveNodeIDs(ctx csdkTypes.Context, height int64, ids []uint64) {
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	key := types.ActiveNodeIDsKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveNodeIDs(ctx csdkTypes.Context, height int64) (ids []uint64) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.ActiveNodeIDsKey(height)
	value := store.Get(key)
	if value == nil {
		return ids
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &ids)
	return ids
}

func (k Keeper) GetNodesOfAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress) (nodes []types.Node) {
	count := k.GetNodesCountOfAddress(ctx, address)

	nodes = make([]types.Node, 0, count)
	for i := uint64(0); i < count; i++ {
		id, _ := k.GetNodeIDByAddress(ctx, address, i)

		node, _ := k.GetNode(ctx, id)
		nodes = append(nodes, node)
	}

	return nodes
}

func (k Keeper) GetAllNodes(ctx csdkTypes.Context) (nodes []types.Node) {
	store := ctx.KVStore(k.nodeStoreKey)

	iter := csdkTypes.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var node types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &node)
		nodes = append(nodes, node)
	}

	return nodes
}

// nolint: dupl
func (k Keeper) IterateNodes(ctx csdkTypes.Context, fn func(index int64, node types.Node) (stop bool)) {
	store := ctx.KVStore(k.nodeStoreKey)

	iterator := csdkTypes.KVStorePrefixIterator(store, types.NodeKeyPrefix)
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

func (k Keeper) AddNode(ctx csdkTypes.Context, node types.Node) (allTags csdkTypes.Tags, err csdkTypes.Error) {
	allTags = csdkTypes.EmptyTags()

	node.OwnerPubKey, err = k.accountKeeper.GetPubKey(ctx, node.Owner)
	if err != nil {
		return nil, err
	}

	count := k.GetNodesCountOfAddress(ctx, node.Owner)
	if count >= k.FreeNodesCount(ctx) {
		node.Deposit = k.Deposit(ctx)

		tags, err := k.AddDeposit(ctx, node.Owner, node.Deposit)
		if err != nil {
			return nil, err
		}

		allTags = allTags.AppendTags(tags)
	}

	k.SetNode(ctx, node)
	k.SetNodeIDByAddress(ctx, node.Owner, count, node.ID)

	k.SetNodesCount(ctx, node.ID+1)
	k.SetNodesCountOfAddress(ctx, node.Owner, count+1)

	return allTags, nil
}

func (k Keeper) AddActiveNodeID(ctx csdkTypes.Context, height int64, id uint64) {
	ids := k.GetActiveNodeIDs(ctx, height)

	index := sort.Search(len(ids), func(i int) bool {
		return ids[i] >= id
	})

	if (index == len(ids)) ||
		(index < len(ids) && ids[index] != id) {

		index = len(ids)
	}

	if index != len(ids) {
		return
	}

	ids = append(ids, id)
	k.SetActiveNodeIDs(ctx, height, ids)
}

func (k Keeper) RemoveActiveNodeID(ctx csdkTypes.Context, height int64, id uint64) {
	ids := k.GetActiveNodeIDs(ctx, height)

	index := sort.Search(len(ids), func(i int) bool {
		return ids[i] >= id
	})

	if (index == len(ids)) ||
		(index < len(ids) && ids[index] != id) {

		index = len(ids)
	}

	if index == len(ids) {
		return
	}

	ids[index] = ids[len(ids)-1]
	ids = ids[:len(ids)-1]

	k.SetActiveNodeIDs(ctx, height, ids)
}
