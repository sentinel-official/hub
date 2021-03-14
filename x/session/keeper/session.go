package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func (k Keeper) SetCount(ctx sdk.Context, count uint64) {
	key := types.CountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetCount(ctx sdk.Context) (count uint64) {
	store := k.Store(ctx)

	key := types.CountKey
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

func (k Keeper) GetSessions(ctx sdk.Context, skip, limit int) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Session
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

func (k Keeper) IterateSessions(ctx sdk.Context, fn func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var session types.Session
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &session)

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}

func (k Keeper) SetSessionForSubscription(ctx sdk.Context, subscription, id uint64) {
	key := types.SessionForSubscriptionKey(subscription, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSessionsForSubscription(ctx sdk.Context, id uint64, skip, limit int) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetSessionForSubscriptionKeyPrefix(id)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSession(ctx, types.IDFromSessionForSubscriptionKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetSessionForNode(ctx sdk.Context, address hub.NodeAddress, id uint64) {
	key := types.SessionForNodeKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSessionsForNode(ctx sdk.Context, address hub.NodeAddress, skip, limit int) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetSessionForNodeKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSession(ctx, types.IDFromSessionForNodeKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetSessionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.SessionForAddressKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSessionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetSessionForAddressKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSession(ctx, types.IDFromSessionForAddressKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetActiveSession(ctx sdk.Context, address sdk.AccAddress, subscription uint64, node hub.NodeAddress, id uint64) {
	key := types.ActiveSessionForAddressKey(address, subscription, node)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetActiveSession(ctx sdk.Context, address sdk.AccAddress, id uint64, node hub.NodeAddress) (session types.Session, found bool) {
	store := k.Store(ctx)

	key := types.ActiveSessionForAddressKey(address, id, node)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &id)
	return k.GetSession(ctx, id)
}

func (k Keeper) DeleteActiveSession(ctx sdk.Context, address sdk.AccAddress, id uint64, node hub.NodeAddress) {
	key := types.ActiveSessionForAddressKey(address, id, node)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) GetActiveSessionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActiveSessionForAddressKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var id uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &id)

		item, _ := k.GetSession(ctx, id)
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetActiveSessionAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.ActiveSessionAtKey(at, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteActiveSessionAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.ActiveSessionAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) IterateActiveSessionsAt(ctx sdk.Context, end time.Time, fn func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.ActiveSessionAtKeyPrefix, sdk.PrefixEndBytes(types.GetActiveSessionAtKeyPrefix(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		session, _ := k.GetSession(ctx, types.IDFromActiveSessionAtKey(iter.Key()))
		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}
