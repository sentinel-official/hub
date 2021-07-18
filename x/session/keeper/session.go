package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func (k *Keeper) SetCount(ctx sdk.Context, count uint64) {
	key := types.CountKey
	value := k.cdc.MustMarshalBinaryBare(&protobuf.UInt64Value{Value: count})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetCount(ctx sdk.Context) uint64 {
	store := k.Store(ctx)

	key := types.CountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	var count protobuf.UInt64Value
	k.cdc.MustUnmarshalBinaryBare(value, &count)

	return count.Value
}

func (k *Keeper) SetSession(ctx sdk.Context, session types.Session) {
	key := types.SessionKey(session.Id)
	value := k.cdc.MustMarshalBinaryBare(&session)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetSession(ctx sdk.Context, id uint64) (session types.Session, found bool) {
	store := k.Store(ctx)

	key := types.SessionKey(id)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &session)
	return session, true
}

func (k *Keeper) GetSessions(ctx sdk.Context, skip, limit int64) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Session
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

func (k *Keeper) IterateSessions(ctx sdk.Context, fn func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var session types.Session
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &session)

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSessionForSubscription(ctx sdk.Context, subscription, id uint64) {
	key := types.SessionForSubscriptionKey(subscription, id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetSessionsForSubscription(ctx sdk.Context, id uint64, skip, limit int64) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
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

func (k *Keeper) SetSessionForNode(ctx sdk.Context, address hubtypes.NodeAddress, id uint64) {
	key := types.SessionForNodeKey(address, id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetSessionsForNode(ctx sdk.Context, address hubtypes.NodeAddress, skip, limit int64) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
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

func (k *Keeper) SetInactiveSessionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.InactiveSessionForAddressKey(address, id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteInactiveSessionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.InactiveSessionForAddressKey(address, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetInactiveSessionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int64) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetInactiveSessionForAddressKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSession(ctx, types.IDFromStatusSessionForAddressKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k *Keeper) SetActiveSessionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.ActiveSessionForAddressKey(address, id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteActiveSessionForAddress(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	key := types.ActiveSessionForAddressKey(address, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetActiveSessionsForAddress(ctx sdk.Context, address sdk.AccAddress, skip, limit int64) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActiveSessionForAddressKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetSession(ctx, types.IDFromStatusSessionForAddressKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k *Keeper) SetInactiveSessionAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.InactiveSessionAtKey(at, id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteInactiveSessionAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.InactiveSessionAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateInactiveSessionsAt(ctx sdk.Context, end time.Time, fn func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.InactiveSessionAtKeyPrefix, sdk.PrefixEndBytes(types.GetInactiveSessionAtKeyPrefix(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var (
			key        = iter.Key()
			session, _ = k.GetSession(ctx, types.IDFromActiveSessionAtKey(key))
		)

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}
