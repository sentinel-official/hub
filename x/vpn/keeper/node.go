package keeper

import (
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetNodeDetails(ctx csdkTypes.Context, details *types.NodeDetails) csdkTypes.Error {
	key := types.NodeKey(details.ID)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(details)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetNodeDetails(ctx csdkTypes.Context, id types.NodeID) (*types.NodeDetails, csdkTypes.Error) {
	key := types.NodeKey(id)
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
	key := types.KeyActiveNodeIDs

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
	key := types.KeyActiveNodeIDs

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

func (k Keeper) GetNodesOfOwner(ctx csdkTypes.Context, owner csdkTypes.AccAddress) ([]*types.NodeDetails, csdkTypes.Error) {
	count, err := k.GetNodesCount(ctx, owner)
	if err != nil {
		return nil, err
	}

	var nodes []*types.NodeDetails
	for index := uint64(0); index < count; index++ {
		id := types.NodeIDFromOwnerCount(owner, index)
		details, err := k.GetNodeDetails(ctx, id)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, details)
	}

	return nodes, nil
}

func (k Keeper) GetNodes(ctx csdkTypes.Context) ([]*types.NodeDetails, csdkTypes.Error) {
	var nodes []*types.NodeDetails
	store := ctx.KVStore(k.NodeStoreKey)

	iter := csdkTypes.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var details types.NodeDetails
		if err := k.cdc.UnmarshalBinaryLengthPrefixed(iter.Value(), &details); err != nil {
			return nil, types.ErrorUnmarshal()
		}

		nodes = append(nodes, &details)
	}

	return nodes, nil
}

func (k Keeper) AddNode(ctx csdkTypes.Context, details *types.NodeDetails) (csdkTypes.Tags, csdkTypes.Error) {
	tags := csdkTypes.EmptyTags()

	count, err := k.GetNodesCount(ctx, details.Owner)
	if err != nil {
		return nil, err
	}

	details.ID = types.NodeIDFromOwnerCount(details.Owner, count)
	if err := k.SetNodeDetails(ctx, details); err != nil {
		return nil, err
	}
	tags = tags.AppendTag("node_id", details.ID.Bytes())

	if err := k.SetNodesCount(ctx, details.Owner, count+1); err != nil {
		return nil, err
	}

	return tags, nil
}
