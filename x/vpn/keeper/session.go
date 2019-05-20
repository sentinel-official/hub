package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

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

func (k Keeper) AddActiveSessionID(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveSessionIDs(ctx, height)
	if ids.Search(id) != ids.Len() {
		return
	}

	ids = ids.Append(id)
	k.SetActiveSessionIDs(ctx, height, ids)
}

func (k Keeper) RemoveActiveSessionID(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveSessionIDs(ctx, height)

	index := ids.Search(id)
	if index == ids.Len() {
		return
	}

	ids = ids.Delete(index)
	k.SetActiveSessionIDs(ctx, height, ids)
}
