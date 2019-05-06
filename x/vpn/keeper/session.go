package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetSession(ctx csdkTypes.Context, session *types.Session) csdkTypes.Error {
	key := types.SessionKey(session.ID)
	value, err := k.cdc.MarshalBinaryLengthPrefixed(session)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetSession(ctx csdkTypes.Context, id sdkTypes.ID) (*types.Session, csdkTypes.Error) {
	key := types.SessionKey(id)

	store := ctx.KVStore(k.SessionStoreKey)
	value := store.Get(key)
	if value == nil {
		return nil, nil
	}

	var session types.Session
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &session); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return &session, nil
}

func (k Keeper) SetActiveSessionIDsAtHeight(ctx csdkTypes.Context, height int64, ids sdkTypes.IDs) csdkTypes.Error {
	key := types.ActiveSessionIDsAtHeightKey(height)

	ids = ids.Sort()
	value, err := k.cdc.MarshalBinaryLengthPrefixed(ids)
	if err != nil {
		return types.ErrorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetActiveSessionIDsAtHeight(ctx csdkTypes.Context, height int64) (sdkTypes.IDs, csdkTypes.Error) {
	key := types.ActiveSessionIDsAtHeightKey(height)

	var ids sdkTypes.IDs
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

func (k Keeper) AddSession(ctx csdkTypes.Context, session *types.Session) (csdkTypes.Tags, csdkTypes.Error) {
	tags := csdkTypes.EmptyTags()

	count, err := k.GetSessionsCount(ctx, session.Client)
	if err != nil {
		return nil, err
	}

	session.ID = sdkTypes.IDFromOwnerAndCount(session.Client, count)
	if err := k.SetSession(ctx, session); err != nil {
		return nil, err
	}
	tags = tags.AppendTag("session_id", session.ID.String())

	if err := k.SetSessionsCount(ctx, session.Client, count+1); err != nil {
		return nil, err
	}

	return tags, nil
}

func (k Keeper) AddActiveSessionIDsAtHeight(ctx csdkTypes.Context, height int64, id sdkTypes.ID) csdkTypes.Error {
	ids, err := k.GetActiveSessionIDsAtHeight(ctx, height)
	if err != nil {
		return err
	}

	if ids.Search(id) != ids.Len() {
		return nil
	}

	ids = ids.Append(id)
	return k.SetActiveSessionIDsAtHeight(ctx, height, ids)
}

func (k Keeper) RemoveActiveSessionIDsAtHeight(ctx csdkTypes.Context, height int64, id sdkTypes.ID) csdkTypes.Error {
	ids, err := k.GetActiveSessionIDsAtHeight(ctx, height)
	if err != nil {
		return err
	}

	index := ids.Search(id)
	if index == ids.Len() {
		return nil
	}

	ids = sdkTypes.NewIDs().Append(ids[:index]...).Append(ids[index+1:]...)
	return k.SetActiveSessionIDsAtHeight(ctx, height, ids)
}
