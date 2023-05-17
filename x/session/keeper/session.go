package keeper

import (
	"fmt"
	"time"

	hubtypes "github.com/sentinel-official/hub/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func (k *Keeper) SetSession(ctx sdk.Context, session types.Session) {
	var (
		store = k.Store(ctx)
		key   = types.SessionKey(session.ID)
		value = k.cdc.MustMarshal(&session)
	)

	store.Set(key, value)
}

func (k *Keeper) GetSession(ctx sdk.Context, id uint64) (session types.Session, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.SessionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshal(value, &session)
	return session, true
}

func (k *Keeper) DeleteSession(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessions(ctx sdk.Context) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Session
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateSessions(ctx sdk.Context, fn func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var session types.Session
		k.cdc.MustUnmarshal(iter.Value(), &session)

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSessionForAccountKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForAccountKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("session for account key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSessionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForNode(ctx sdk.Context, addr hubtypes.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessionsForNode(ctx sdk.Context, addr hubtypes.NodeAddress) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSessionForNodeKeyPrefix(addr))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForNodeKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("session for node key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForSubscriptionKey(subscriptionID, sessionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForSubscriptionKey(subscriptionID, sessionID)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessionsForSubscription(ctx sdk.Context, id uint64) (items types.Sessions) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetSessionForSubscriptionKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForSubscriptionKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSessionForQuota(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForQuotaKey(subscriptionID, addr, sessionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForQuota(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForQuotaKey(subscriptionID, addr, sessionID)
	)

	store.Delete(key)
}

func (k *Keeper) IterateSessionsForQuota(ctx sdk.Context, id uint64, addr sdk.AccAddress, fn func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForQuotaKeyPrefix(id, addr))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		session, found := k.GetSession(ctx, types.IDFromSessionForQuotaKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription quota key %X does not exist", iter.Key()))
		}

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSessionForExpiryAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SessionForExpiryAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForExpiryAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SessionForExpiryAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateSessionsForExpiryAt(ctx sdk.Context, end time.Time, fn func(index int, item types.Session) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.SessionForExpiryAtKeyPrefix, sdk.PrefixEndBytes(types.GetSessionForExpiryAtKeyPrefix(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		session, found := k.GetSession(ctx, types.IDFromSessionForExpiryAtKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("session for expiry at key %X does not exist", iter.Key()))
		}

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}
