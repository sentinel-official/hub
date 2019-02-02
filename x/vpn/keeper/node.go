package keeper

import (
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetNodeDetails(ctx csdkTypes.Context, id string, details *types.NodeDetails) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return types.ErrorMarshal()
	}

	value, err := k.cdc.MarshalBinaryLengthPrefixed(details)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetNodeDetails(ctx csdkTypes.Context, id string) (*types.NodeDetails, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var details types.NodeDetails
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &details); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return &details, nil
}

func (k Keeper) SetActiveNodeIDs(ctx csdkTypes.Context, ids []string) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(types.KeyActiveNodeIDs)
	if err != nil {
		return types.ErrorMarshal()
	}

	sort.Strings(ids)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(ids)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetActiveNodeIDs(ctx csdkTypes.Context) ([]string, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(types.KeyActiveNodeIDs)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	var ids []string

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

func (k Keeper) GetNodesOfOwner(ctx csdkTypes.Context, owner csdkTypes.AccAddress) ([]types.NodeDetails, csdkTypes.Error) {
	count, err := k.GetNodesCount(ctx, owner)
	if err != nil {
		return nil, err
	}

	var nodes []types.NodeDetails
	for index := uint64(0); index < count; index++ {
		id := types.NodeKey(owner, index)
		details, err := k.GetNodeDetails(ctx, id)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, *details)
	}

	return nodes, nil
}

func (k Keeper) GetNodes(ctx csdkTypes.Context) ([]types.NodeDetails, csdkTypes.Error) {
	var nodes []types.NodeDetails
	store := ctx.KVStore(k.NodeStoreKey)

	iter := csdkTypes.KVStorePrefixIterator(store, types.NodesCountKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		owner := csdkTypes.AccAddress(iter.Key()[len(types.NodesCountKeyPrefix):])
		_nodes, err := k.GetNodesOfOwner(ctx, owner)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, _nodes...)
	}

	return nodes, nil
}
