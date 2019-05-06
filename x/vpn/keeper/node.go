package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetNode(ctx csdkTypes.Context, node *types.Node) csdkTypes.Error {
	key := types.NodeKey(node.ID)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(node)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetNode(ctx csdkTypes.Context, id sdkTypes.ID) (*types.Node, csdkTypes.Error) {
	key := types.NodeKey(id)

	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var node types.Node
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &node); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return &node, nil
}

func (k Keeper) SetActiveNodeIDsAtHeight(ctx csdkTypes.Context, height int64, ids sdkTypes.IDs) csdkTypes.Error {
	key := types.ActiveNodeIDsAtHeightKey(height)

	ids = ids.Sort()
	value, err := k.cdc.MarshalBinaryLengthPrefixed(ids)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetActiveNodeIDsAtHeight(ctx csdkTypes.Context, height int64) (sdkTypes.IDs, csdkTypes.Error) {
	key := types.ActiveNodeIDsAtHeightKey(height)

	var ids sdkTypes.IDs
	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return ids, nil
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &ids); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return ids, nil
}

func (k Keeper) SetNodesCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress, count uint64) csdkTypes.Error {
	key := types.NodesCountKey(owner)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(count)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetNodesCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress) (uint64, csdkTypes.Error) {
	key := types.NodesCountKey(owner)

	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return 0, nil
	}

	var count uint64
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &count); err != nil {
		return 0, types.ErrorUnmarshal()
	}

	return count, nil
}

func (k Keeper) GetNodesOfOwner(ctx csdkTypes.Context, owner csdkTypes.AccAddress) ([]*types.Node, csdkTypes.Error) {
	count, err := k.GetNodesCount(ctx, owner)
	if err != nil {
		return nil, err
	}

	var nodes []*types.Node
	for index := uint64(0); index < count; index++ {
		id := sdkTypes.IDFromOwnerAndCount(owner, index)
		node, err := k.GetNode(ctx, id)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (k Keeper) GetNodes(ctx csdkTypes.Context) ([]*types.Node, csdkTypes.Error) {
	var nodes []*types.Node
	store := ctx.KVStore(k.NodeStoreKey)

	iter := csdkTypes.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var node types.Node
		if err := k.cdc.UnmarshalBinaryLengthPrefixed(iter.Value(), &node); err != nil {
			return nil, types.ErrorUnmarshal()
		}

		nodes = append(nodes, &node)
	}

	return nodes, nil
}

func (k Keeper) AddNode(ctx csdkTypes.Context, node *types.Node) (csdkTypes.Tags, csdkTypes.Error) {
	tags := csdkTypes.EmptyTags()

	count, err := k.GetNodesCount(ctx, node.Owner)
	if err != nil {
		return nil, err
	}

	node.ID = sdkTypes.IDFromOwnerAndCount(node.Owner, count)
	if err := k.SetNode(ctx, node); err != nil {
		return nil, err
	}
	tags = tags.AppendTag("node_id", node.ID.String())

	if err := k.SetNodesCount(ctx, node.Owner, count+1); err != nil {
		return nil, err
	}

	return tags, nil
}

func (k Keeper) AddActiveNodeIDAtHeight(ctx csdkTypes.Context, height int64, id sdkTypes.ID) csdkTypes.Error {
	ids, err := k.GetActiveNodeIDsAtHeight(ctx, height)
	if err != nil {
		return err
	}

	if ids.Search(id) != ids.Len() {
		return nil
	}

	ids = ids.Append(id)
	return k.SetActiveNodeIDsAtHeight(ctx, height, ids)
}

func (k Keeper) RemoveActiveNodeIDAtHeight(ctx csdkTypes.Context, height int64, id sdkTypes.ID) csdkTypes.Error {
	ids, err := k.GetActiveNodeIDsAtHeight(ctx, height)
	if err != nil {
		return err
	}

	index := ids.Search(id)
	if index == ids.Len() {
		return nil
	}

	ids = sdkTypes.NewIDs().Append(ids[:index]...).Append(ids[index+1:]...)
	return k.SetActiveNodeIDsAtHeight(ctx, height, ids)
}
