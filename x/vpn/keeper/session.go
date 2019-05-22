package keeper

import (
	"sort"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

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

func (k Keeper) GetSession(ctx csdkTypes.Context, id uint64) (session types.Session, found bool) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.SessionKey(id)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &session)
	return session, true
}

func (k Keeper) SetSessionIDBySubscriptionID(ctx csdkTypes.Context, i, j, id uint64) {
	key := types.SessionIDBySubscriptionIDKey(i, j)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSessionIDBySubscriptionID(ctx csdkTypes.Context, i, j uint64) (id uint64, found bool) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.SessionIDBySubscriptionIDKey(i, j)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) SetActiveSessionIDs(ctx csdkTypes.Context, height int64, ids []uint64) {
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	key := types.ActiveSessionIDsKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveSessionIDs(ctx csdkTypes.Context, height int64) (ids []uint64) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.ActiveSessionIDsKey(height)
	value := store.Get(key)
	if value == nil {
		return ids
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &ids)
	return ids
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

func (k Keeper) AddActiveSessionID(ctx csdkTypes.Context, height int64, id uint64) {
	ids := k.GetActiveSessionIDs(ctx, height)

	index := sort.Search(len(ids), func(i int) bool {
		return ids[i] >= id
	})

	if (index == len(ids)) ||
		(index < len(ids) && ids[index] != id) {

		index = len(ids)
	}

	if index != len(ids) {
		return
	}

	ids = append(ids, id)
	k.SetActiveSessionIDs(ctx, height, ids)
}

func (k Keeper) RemoveActiveSessionID(ctx csdkTypes.Context, height int64, id uint64) {
	ids := k.GetActiveSessionIDs(ctx, height)

	index := sort.Search(len(ids), func(i int) bool {
		return ids[i] >= id
	})

	if (index == len(ids)) ||
		(index < len(ids) && ids[index] != id) {

		index = len(ids)
	}

	if index == len(ids) {
		return
	}

	ids[index] = ids[len(ids)-1]
	ids = ids[:len(ids)-1]

	k.SetActiveSessionIDs(ctx, height, ids)
}
