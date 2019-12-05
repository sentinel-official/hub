package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetFreeClient(ctx sdk.Context, freeClient types.FreeClient) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(freeClient)

	store := ctx.KVStore(k.freeClientKey)
	store.Set(types.FreeClientsKeyPrefix, value)
}

func (k Keeper) SetFreeNodesOfClient(ctx sdk.Context, freeClient types.FreeClient) {
	key := types.FreeNodesOfClientKey(freeClient.Client, freeClient.NodeID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(freeClient.NodeID)

	store := ctx.KVStore(k.freeClientKey)
	store.Set(key, value)
}

func (k Keeper) SetFreeClientOfNode(ctx sdk.Context, freeClient types.FreeClient) {
	key := types.FreeClientOfNodeKey(freeClient.NodeID, freeClient.Client)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(freeClient.Client)

	store := ctx.KVStore(k.freeClientKey)
	store.Set(key, value)
}

func (k Keeper) GetAllFreeClients(ctx sdk.Context) (freeClients []types.FreeClient) {
	store := ctx.KVStore(k.freeClientKey)

	iter := sdk.KVStorePrefixIterator(store, types.FreeClientsKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var freeClient types.FreeClient
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &freeClient)
		freeClients = append(freeClients, freeClient)
	}

	return freeClients
}

func (k Keeper) GetFreeClientsOfNode(ctx sdk.Context, nodeID hub.NodeID) (freeClients []sdk.AccAddress) {
	store := ctx.KVStore(k.freeClientKey)

	key := types.FreeClientOfNodeKey(nodeID, nil)
	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var freeClient sdk.AccAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &freeClient)
		freeClients = append(freeClients, freeClient)
	}

	return freeClients
}

func (k Keeper) GetFreeNodesOfClient(ctx sdk.Context, client sdk.AccAddress) (freeNodes []hub.NodeID) {
	store := ctx.KVStore(k.freeClientKey)

	key := types.FreeNodesOfClientKey(client, nil)
	iter := sdk.KVStorePrefixIterator(store, key)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var freeNode hub.NodeID
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &freeNode)
		freeNodes = append(freeNodes, freeNode)
	}

	return freeNodes
}
