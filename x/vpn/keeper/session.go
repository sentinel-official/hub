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

func (k Keeper) SetActiveSessionIDsAtHeight(ctx csdkTypes.Context, height int64, ids sdkTypes.IDs) {
	ids = ids.Sort()

	key := types.ActiveSessionIDsAtHeightKey(height)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(ids)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetActiveSessionIDsAtHeight(ctx csdkTypes.Context, height int64) (ids sdkTypes.IDs) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.ActiveSessionIDsAtHeightKey(height)
	value := store.Get(key)
	if value == nil {
		return ids
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &ids)
	return ids
}

func (k Keeper) SetSessionsCount(ctx csdkTypes.Context, address csdkTypes.AccAddress, count uint64) {
	key := types.SessionsCountKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.sessionStoreKey)
	store.Set(key, value)
}

func (k Keeper) GetSessionsCount(ctx csdkTypes.Context, address csdkTypes.AccAddress) (count uint64) {
	store := ctx.KVStore(k.sessionStoreKey)

	key := types.SessionsCountKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) GetSessionsOfAddress(ctx csdkTypes.Context, address csdkTypes.AccAddress) (sessions []types.Session) {
	count := k.GetSessionsCount(ctx, address)

	sessions = make([]types.Session, count)
	for index := uint64(0); index < count; index++ {
		id := sdkTypes.IDFromAddressAndCount(address, index)
		session, _ := k.GetSession(ctx, id)
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

func (k Keeper) AddSession(ctx csdkTypes.Context, session types.Session) (allTags csdkTypes.Tags, err csdkTypes.Error) {
	allTags = csdkTypes.EmptyTags()

	tags, err := k.AddDeposit(ctx, session.Client, session.Deposit)
	if err != nil {
		return nil, err
	}

	allTags = allTags.AppendTags(tags)

	session.ClientPubKey, err = k.accountKeeper.GetPubKey(ctx, session.Client)
	if err != nil {
		return nil, err
	}

	count := k.GetSessionsCount(ctx, session.Client)
	session.ID = sdkTypes.IDFromAddressAndCount(session.Client, count)

	k.SetSession(ctx, session)
	allTags = allTags.AppendTag(types.TagSessionID, session.ID.String())

	k.SetSessionsCount(ctx, session.Client, count+1)
	return allTags, nil
}

func (k Keeper) AddActiveSessionIDAtHeight(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveSessionIDsAtHeight(ctx, height)
	if ids.Search(id) != ids.Len() {
		return
	}

	ids = ids.Append(id)
	k.SetActiveSessionIDsAtHeight(ctx, height, ids)
}

func (k Keeper) RemoveActiveSessionIDAtHeight(ctx csdkTypes.Context, height int64, id sdkTypes.ID) {
	ids := k.GetActiveSessionIDsAtHeight(ctx, height)

	index := ids.Search(id)
	if index == ids.Len() {
		return
	}

	ids = ids.Delete(index)
	k.SetActiveSessionIDsAtHeight(ctx, height, ids)
}
