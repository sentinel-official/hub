package keeper

import (
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetSessionDetails(ctx csdkTypes.Context, id string, details *types.SessionDetails) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return types.ErrorMarshal()
	}

	value, err := k.cdc.MarshalBinaryLengthPrefixed(details)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetSessionDetails(ctx csdkTypes.Context, id string) (*types.SessionDetails, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(id)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var details types.SessionDetails
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &details); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return &details, nil
}

func (k Keeper) SetActiveSessionIDs(ctx csdkTypes.Context, ids []string) csdkTypes.Error {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(types.KeyActiveSessionIDs)
	if err != nil {
		return types.ErrorMarshal()
	}

	sort.Strings(ids)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(ids)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetActiveSessionIDs(ctx csdkTypes.Context) ([]string, csdkTypes.Error) {
	key, err := k.cdc.MarshalBinaryLengthPrefixed(types.KeyActiveSessionIDs)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	var ids []string

	store := ctx.KVStore(k.SessionStoreKey)
	value := store.Get(key)
	if value == nil {
		return ids, nil
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &ids); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return ids, nil
}

func (k Keeper) SetSessionsCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress, count uint64) csdkTypes.Error {
	key := types.SessionsCountKey(owner)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(count)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetSessionsCount(ctx csdkTypes.Context, owner csdkTypes.AccAddress) (uint64, csdkTypes.Error) {
	key := types.SessionsCountKey(owner)
	store := ctx.KVStore(k.SessionStoreKey)
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
