package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetSessionsCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.sessionKey)
	store.Set(types.SessionsCountKey, value)
}

func (k Keeper) GetSessionsCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.sessionKey)

	value := store.Get(types.SessionsCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSession(ctx sdk.Context, session types.Session) {
	key := types.SessionKey(session.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(session)

	store := ctx.KVStore(k.sessionKey)
	store.Set(key, value)
}

func (k Keeper) GetSession(ctx sdk.Context, id hub.SessionID) (session types.Session, found bool) {
	store := ctx.KVStore(k.sessionKey)

	key := types.SessionKey(id)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &session)
	return session, true
}

func (k Keeper) SetSessionsCountOfSubscription(ctx sdk.Context, id hub.SubscriptionID, count uint64) {
	key := types.SessionsCountOfSubscriptionKey(id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.sessionKey)
	store.Set(key, value)
}

func (k Keeper) GetSessionsCountOfSubscription(ctx sdk.Context, id hub.SubscriptionID) (count uint64) {
	store := ctx.KVStore(k.sessionKey)

	key := types.SessionsCountOfSubscriptionKey(id)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)

	return count
}

func (k Keeper) SetSessionIDBySubscriptionID(ctx sdk.Context, i hub.SubscriptionID, j uint64, id hub.SessionID) {
	key := types.SessionIDBySubscriptionIDKey(i, j)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.sessionKey)
	store.Set(key, value)
}

func (k Keeper) GetSessionIDBySubscriptionID(ctx sdk.Context,
	i hub.SubscriptionID, j uint64) (id hub.SessionID, found bool) {
	store := ctx.KVStore(k.sessionKey)

	key := types.SessionIDBySubscriptionIDKey(i, j)
	value := store.Get(key)
	if value == nil {
		return hub.NewSessionID(0), false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetActiveSessionIDs(ctx sdk.Context, height int64, ids hub.IDs) {
	ids.Sort()

	key := types.ActiveSessionIDsKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)

	store := ctx.KVStore(k.sessionKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveSessionIDs(ctx sdk.Context, height int64) (ids hub.IDs) {
	store := ctx.KVStore(k.sessionKey)

	key := types.ActiveSessionIDsKey(height)
	value := store.Get(key)
	if value == nil {
		return ids
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &ids)
	return ids
}

func (k Keeper) DeleteActiveSessionIDs(ctx sdk.Context, height int64) {
	store := ctx.KVStore(k.sessionKey)

	key := types.ActiveSessionIDsKey(height)
	store.Delete(key)
}

func (k Keeper) GetSessionsOfSubscription(ctx sdk.Context, id hub.SubscriptionID) (sessions []types.Session) {
	count := k.GetSessionsCountOfSubscription(ctx, id)

	sessions = make([]types.Session, 0, count)
	for i := uint64(0); i < count; i++ {
		_id, _ := k.GetSessionIDBySubscriptionID(ctx, id, i)

		session, _ := k.GetSession(ctx, _id)
		sessions = append(sessions, session)
	}

	return sessions
}

func (k Keeper) GetAllSessions(ctx sdk.Context) (sessions []types.Session) {
	store := ctx.KVStore(k.sessionKey)

	iter := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var session types.Session
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &session)
		sessions = append(sessions, session)
	}

	return sessions
}

func (k Keeper) AddSessionIDToActiveList(ctx sdk.Context, height int64, id hub.SessionID) {
	ids := k.GetActiveSessionIDs(ctx, height)

	index := ids.Search(id)
	if index != len(ids) {
		return
	}

	ids = ids.Append(id)
	k.SetActiveSessionIDs(ctx, height, ids)
}

func (k Keeper) RemoveSessionIDFromActiveList(ctx sdk.Context, height int64, id hub.SessionID) {
	ids := k.GetActiveSessionIDs(ctx, height)

	index := ids.Search(id)
	if index == len(ids) {
		return
	}

	ids = ids.Delete(index)
	k.SetActiveSessionIDs(ctx, height, ids)
}
