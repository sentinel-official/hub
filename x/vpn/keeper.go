package vpn

import (
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	NodeStoreKey    csdkTypes.StoreKey
	SessionStoreKey csdkTypes.StoreKey
	cdc             *codec.Codec
}

func NewKeeper(cdc *codec.Codec, nodeKey, sessionKey csdkTypes.StoreKey) Keeper {
	return Keeper{
		NodeStoreKey:    nodeKey,
		SessionStoreKey: sessionKey,
		cdc:             cdc,
	}
}

func (k Keeper) SetNodeDetails(ctx csdkTypes.Context, id string, details *NodeDetails) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return errorMarshal()
	}

	value, err := k.cdc.MarshalBinaryLengthPrefixed(details)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetNodeDetails(ctx csdkTypes.Context, id string) (*NodeDetails, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return nil, errorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var details NodeDetails
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &details); err != nil {
		return nil, errorUnmarshal()
	}

	return &details, nil
}

func (k Keeper) SetActiveNodeIDs(ctx csdkTypes.Context, ids []string) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(KeyActiveNodeIDs)
	if err != nil {
		return errorMarshal()
	}

	sort.Strings(ids)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(ids)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetActiveNodeIDs(ctx csdkTypes.Context) ([]string, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(KeyActiveNodeIDs)
	if err != nil {
		return nil, errorMarshal()
	}

	var ids []string

	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return ids, nil
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &ids); err != nil {
		return nil, errorUnmarshal()
	}

	return ids, nil
}

func (k Keeper) SetNodesCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress, count uint64) csdkTypes.Error {
	key := NodesCountKey(owner)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(count)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetNodesCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress) (uint64, csdkTypes.Error) {
	key := NodesCountKey(owner)
	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return 0, nil
	}

	var count uint64
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &count); err != nil {
		return 0, errorUnmarshal()
	}

	return count, nil
}

func (k Keeper) SetSessionDetails(ctx csdkTypes.Context, id string, details *SessionDetails) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return errorMarshal()
	}

	value, err := k.cdc.MarshalBinaryLengthPrefixed(details)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetSessionDetails(ctx csdkTypes.Context, id string) (*SessionDetails, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return nil, errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var details SessionDetails
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &details); err != nil {
		return nil, errorUnmarshal()
	}

	return &details, nil
}

func (k Keeper) SetActiveSessionIDs(ctx csdkTypes.Context, ids []string) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(KeyActiveSessionIDs)
	if err != nil {
		return errorMarshal()
	}

	sort.Strings(ids)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(ids)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetActiveSessionIDs(ctx csdkTypes.Context) ([]string, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(KeyActiveSessionIDs)
	if err != nil {
		return nil, errorMarshal()
	}

	var ids []string

	store := ctx.KVStore(k.SessionStoreKey)
	value := store.Get(key)
	if value == nil {
		return ids, nil
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &ids); err != nil {
		return nil, errorUnmarshal()
	}

	return ids, nil
}

func (k Keeper) SetSessionsCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress, count uint64) csdkTypes.Error {
	key := SessionsCountKey(owner)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(count)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetSessionsCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress) (uint64, csdkTypes.Error) {
	key := SessionsCountKey(owner)
	store := ctx.KVStore(k.SessionStoreKey)
	value := store.Get(key)
	if value == nil {
		return 0, nil
	}

	var count uint64
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &count); err != nil {
		return 0, errorUnmarshal()
	}

	return count, nil
}

func (k Keeper) GetNodesOfOwner(ctx csdkTypes.Context, owner csdkTypes.AccAddress) ([]NodeDetails, csdkTypes.Error) {
	count, err := k.GetNodesCount(ctx, owner)
	if err != nil {
		return nil, err
	}

	var nodes []NodeDetails
	for index := uint64(0); index < count; index++ {
		id := NodeKey(owner, index)
		details, err := k.GetNodeDetails(ctx, id)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, *details)
	}

	return nodes, nil
}

func (k Keeper) GetNodes(ctx csdkTypes.Context) ([]NodeDetails, csdkTypes.Error) {
	var nodes []NodeDetails
	store := ctx.KVStore(k.NodeStoreKey)

	iter := csdkTypes.KVStorePrefixIterator(store, NodesCountKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		owner := csdkTypes.AccAddress(iter.Key()[len(NodesCountKeyPrefix):])
		_nodes, err := k.GetNodesOfOwner(ctx, owner)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, _nodes...)
	}

	return nodes, nil
}
