package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetNodesCount(ctx csdkTypes.Context, address csdkTypes.AccAddress, count uint64) {
	key := types.NodesCountKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetNodesCount(ctx csdkTypes.Context, address csdkTypes.AccAddress) (count uint64) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodesCountKey(address)
	value := store.Get(key)
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

func (k Keeper) GetNode(ctx csdkTypes.Context, id sdkTypes.ID) (node types.Node, found bool) {
	store := ctx.KVStore(k.nodeStoreKey)

	key := types.NodeKey(id)
	value := store.Get(key)
	if value == nil {
		return node, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &node)
	return node, true
}

func (k Keeper) SetActiveNodeIDs(ctx csdkTypes.Context, height int64, ids sdkTypes.IDs) {
	ids = ids.Sort()

	key := types.ActiveNodeIDsKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)

	store := ctx.KVStore(k.nodeStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveNodeIDs(ctx csdkTypes.Context, height int64) (ids sdkTypes.IDs) {
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
	count := k.GetNodesCount(ctx, address)

	nodes = make([]types.Node, 0, count)
	for index := uint64(0); index < count; index++ {
		id := types.NodeID(address, index)
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

func (k Keeper) AddNode(ctx csdkTypes.Context, node types.Node) (allTags csdkTypes.Tags, err csdkTypes.Error) {
	allTags = csdkTypes.EmptyTags()

	count := k.GetNodesCount(ctx, node.Owner)
	node.ID = types.NodeID(node.Owner, count)

	node.OwnerPubKey, err = k.accountKeeper.GetPubKey(ctx, node.Owner)
	if err != nil {
		return nil, err
	}

	if count >= k.FreeNodesCount(ctx) {
		node.Deposit = k.Deposit(ctx)

		tags, err := k.AddDeposit(ctx, node.Owner, node.Deposit)
		if err != nil {
			return nil, err
		}

		allTags = allTags.AppendTags(tags)
	}

	k.SetNode(ctx, node)
	allTags = allTags.AppendTag(types.TagNodeID, node.ID.String())

	k.SetNodesCount(ctx, node.Owner, count+1)
	return allTags, nil
}

func (k Keeper) AddActiveNodeID(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveNodeIDs(ctx, height)
	if ids.Search(id) != ids.Len() {
		return
	}

	ids = ids.Append(id)
	k.SetActiveNodeIDs(ctx, height, ids)
}

func (k Keeper) RemoveActiveNodeID(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveNodeIDs(ctx, height)

	index := ids.Search(id)
	if index == ids.Len() {
		return
	}

	ids = ids.Delete(index)
	k.SetActiveNodeIDs(ctx, height, ids)
}
