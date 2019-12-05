package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetFreeClient(ctx sdk.Context, freeClient types.FreeClient) {
	key := types.FreeClientKey(freeClient.NodeID, freeClient.Client)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(freeClient)

	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) GetFreeClient(ctx sdk.Context, id hub.SessionID) (session types.Session, found bool) {
	store := ctx.KVStore(k.sessionKey)

	key := types.SessionKey(id)
	value := store.Get(key)
	if value == nil {
		return session, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &session)
	return session, true
}

func (k Keeper) GetAllFreeClients(ctx sdk.Context) (freeClients []types.FreeClient) {
	store := ctx.KVStore(k.nodeKey)

	iter := sdk.KVStorePrefixIterator(store, types.FreeClientKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var freeClient types.FreeClient
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &freeClient)
		freeClients = append(freeClients, freeClient)
	}

	return freeClients
}

func (k Keeper) GetFreeClientsOfNode(ctx sdk.Context, nodeID hub.NodeID) (freeClients []types.FreeClient) {
	store := ctx.KVStore(k.nodeKey)

	key := types.FreeClientKey(nodeID, nil)
	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var freeClient types.FreeClient
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &freeClient)
		freeClients = append(freeClients, freeClient)
	}

	return freeClients
}

func (k Keeper) GetFreeNodesOfClient(ctx sdk.Context, client sdk.AccAddress) (freeClients []types.FreeClient) {
	store := ctx.KVStore(k.nodeKey)

	key := types.FreeClientKey(nil, nil)
	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var freeClient types.FreeClient
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &freeClient)
		if freeClient.Client.String() == client.String() {
			freeClients = append(freeClients, freeClient)
		}
	}

	return freeClients
}
