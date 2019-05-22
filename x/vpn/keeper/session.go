package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func (k Keeper) SetSessionsCount(ctx csdkTypes.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(types.SessionsCountKey, value)
}

func (k Keeper) GetSessionsCount(ctx csdkTypes.Context) (count uint64) {
	store := ctx.KVStore(k.sessionStoreKey)

	value := store.Get(types.SessionsCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSession(ctx csdkTypes.Context, session types.Session) {
	key := types.SessionKey(session.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(session)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSession(ctx csdkTypes.Context, id sdkTypes.ID) (session types.Session, found bool) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.SessionKey(id)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &session)
	return session, true
}

func (k Keeper) SetSessionsCountOfSubscription(ctx csdkTypes.Context, id sdkTypes.ID, count uint64) {
	key := types.SessionsCountOfSubscriptionKey(id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSessionsCountOfSubscription(ctx csdkTypes.Context, id sdkTypes.ID) (count uint64) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.SessionsCountOfSubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSessionIDBySubscriptionID(ctx csdkTypes.Context, i sdkTypes.ID, j uint64, id sdkTypes.ID) {
	key := types.SessionIDBySubscriptionIDKey(i, j)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSessionIDBySubscriptionID(ctx csdkTypes.Context, i sdkTypes.ID, j uint64) (id sdkTypes.ID, found bool) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.SessionIDBySubscriptionIDKey(i, j)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetActiveSessionIDs(ctx csdkTypes.Context, height int64, ids sdkTypes.IDs) {
	ids = ids.Sort()

	key := types.ActiveSessionIDsKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveSessionIDs(ctx csdkTypes.Context, height int64) (ids sdkTypes.IDs) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.ActiveSessionIDsKey(height)
	value := store.Get(key)
	if value == nil {
		return ids
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &ids)
	return ids
}

func (k Keeper) DeleteActiveSessionIDs(ctx csdkTypes.Context, height int64) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.ActiveSessionIDsKey(height)
	store.Delete(key)
}

func (k Keeper) GetSessionsOfSubscription(ctx csdkTypes.Context, id sdkTypes.ID) (sessions []types.Session) {
	count := k.GetSessionsCountOfSubscription(ctx, id)

	sessions = make([]types.Session, 0, count)
	for i := uint64(0); i < count; i++ {
		_id, _ := k.GetSessionIDBySubscriptionID(ctx, id, i)

		session, _ := k.GetSession(ctx, _id)
		sessions = append(sessions, session)
	}

	return sessions
}

func (k Keeper) GetAllSessions(ctx csdkTypes.Context) (sessions []types.Session) {
	store := ctx.KVStore(k.sessionStoreKey)

	iter := csdkTypes.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var session types.Session
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &session)
		sessions = append(sessions, session)
	}

	return sessions
}

func (k Keeper) AddSessionIDToActiveList(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveSessionIDs(ctx, height)

	index := ids.Search(id)
	if index != len(ids) {
		return
	}

	ids = ids.Append(id)
	k.SetActiveSessionIDs(ctx, height, ids)
}

func (k Keeper) RemoveSessionIDFromActiveList(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveSessionIDs(ctx, height)

	index := ids.Search(id)
	if index == len(ids) {
		return
	}

	ids = ids.Delete(index)
	k.SetActiveSessionIDs(ctx, height, ids)
}
