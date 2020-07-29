package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func (k Keeper) SetSessionsCount(ctx sdk.Context, count uint64) {
	key := types.SessionsCountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSessionsCount(ctx sdk.Context) (count uint64) {
	store := k.Store(ctx)

	key := types.SessionsCountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetSession(ctx sdk.Context, session types.Session) {
	key := types.SessionKey(session.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(session)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSession(ctx sdk.Context, id uint64) (session types.Session, found bool) {
	store := k.Store(ctx)

	key := types.SessionKey(id)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &session)
	return session, true
}

func (k Keeper) GetSessions(ctx sdk.Context) (items types.Sessions) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Session
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) IterateSessions(ctx sdk.Context, f func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var session types.Session
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &session)

		if stop := f(i, session); stop {
			break
		}
		i++
	}
}

func (k Keeper) SetActiveSession(ctx sdk.Context, s uint64, n hub.NodeAddress, a sdk.AccAddress, id uint64) {
	key := types.ActiveSessionKey(s, n, a)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetActiveSession(ctx sdk.Context, s uint64, n hub.NodeAddress, a sdk.AccAddress) (session types.Session, found bool) {
	store := k.Store(ctx)

	key := types.ActiveSessionKey(s, n, a)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	var id uint64
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)

	return k.GetSession(ctx, id)
}
