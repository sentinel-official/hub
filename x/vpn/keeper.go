package vpn

import (
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
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

func (k Keeper) SetNodeDetails(ctx csdkTypes.Context, id string, details *sdkTypes.VPNNodeDetails) csdkTypes.Error {
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

func (k Keeper) GetNodeDetails(ctx csdkTypes.Context, id string) (*sdkTypes.VPNNodeDetails, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return nil, errorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var details sdkTypes.VPNNodeDetails
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &details); err != nil {
		return nil, errorUnmarshal()
	}

	return &details, nil
}

func (k Keeper) SetActiveNodeIDs(ctx csdkTypes.Context, ids []string) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveNodeIDs)
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
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveNodeIDs)
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
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.VPNNodesCountKey(owner))
	if err != nil {
		return errorMarshal()
	}

	value, err := k.cdc.MarshalBinaryLengthPrefixed(count)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.NodeStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetNodesCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress) (uint64, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.VPNNodesCountKey(owner))
	if err != nil {
		return 0, errorMarshal()
	}

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

func (k Keeper) SetSessionDetails(ctx csdkTypes.Context, id string, details *sdkTypes.VPNSessionDetails) csdkTypes.Error {
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

func (k Keeper) GetSessionDetails(ctx csdkTypes.Context, id string) (*sdkTypes.VPNSessionDetails, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return nil, errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var details sdkTypes.VPNSessionDetails
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &details); err != nil {
		return nil, errorUnmarshal()
	}

	return &details, nil
}

func (k Keeper) SetActiveSessionIDs(ctx csdkTypes.Context, ids []string) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveSessionIDs)
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
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveSessionIDs)
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
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.VPNSessionsCountKey(owner))
	if err != nil {
		return errorMarshal()
	}

	value, err := k.cdc.MarshalBinaryLengthPrefixed(count)
	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetSessionsCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress) (uint64, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.VPNSessionsCountKey(owner))
	if err != nil {
		return 0, errorMarshal()
	}

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

/* ____________________________________________________________________________________________________ */

func AddVPN(ctx csdkTypes.Context, nodeKeeper Keeper, bankKeeper bank.Keeper, details *sdkTypes.VPNNodeDetails) (csdkTypes.Tags, csdkTypes.Error) {
	allTags := csdkTypes.EmptyTags()
	count, err := nodeKeeper.GetNodesCount(ctx, details.Owner)
	if err != nil {
		return nil, err
	}

	id := sdkTypes.VPNNodeKey(details.Owner, count)
	if details, err := nodeKeeper.GetNodeDetails(ctx, id); true {
		if err != nil {
			return nil, err
		}
		if details != nil {
			return nil, nil
		}
	}

	_, tags, err := bankKeeper.SubtractCoins(ctx, details.Owner, details.LockedAmount)
	if err != nil {
		return nil, err
	}
	allTags.AppendTags(tags)

	details.Status = sdkTypes.StatusRegister
	if err := nodeKeeper.SetNodeDetails(ctx, id, details); err != nil {
		return nil, err
	}
	allTags.AppendTag("node_id", []byte(id))

	if err := nodeKeeper.SetNodesCount(ctx, details.Owner, count+1); err != nil {
		return nil, err
	}

	return allTags, nil
}
