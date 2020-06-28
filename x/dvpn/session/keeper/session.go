package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/session/types"
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

func (k Keeper) SetActiveSessionID(ctx sdk.Context, subscription uint64, node hub.NodeAddress, address sdk.AccAddress, id uint64) {
	key := types.ActiveSessionIDKey(subscription, node, address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetActiveSessionID(ctx sdk.Context, subscription uint64, node hub.NodeAddress, address sdk.AccAddress) (id uint64, found bool) {
	store := k.Store(ctx)

	key := types.ActiveSessionIDKey(subscription, node, address)
	value := store.Get(key)
	if value == nil {
		return 0, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return id, true
}

func (k Keeper) GetActiveSession(ctx sdk.Context, subscription uint64, node hub.NodeAddress, address sdk.AccAddress) (session types.Session, found bool) {
	id, found := k.GetActiveSessionID(ctx, subscription, node, address)
	if !found {
		return session, false
	}

	return k.GetSession(ctx, id)
}
