package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

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

func (k Keeper) GetFreeClientOfNode(ctx sdk.Context, nodeID hub.NodeID, client sdk.AccAddress) sdk.AccAddress {
	store := ctx.KVStore(k.freeClientKey)

	key := types.FreeClientOfNodeKey(nodeID, client)

	value := store.Get(key)

	var freeClient sdk.AccAddress
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &freeClient)

	return freeClient
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

func (k Keeper) GetFreeNodeOfClient(ctx sdk.Context, client sdk.AccAddress, nodeID hub.NodeID) hub.NodeID {
	store := ctx.KVStore(k.freeClientKey)

	key := types.FreeNodesOfClientKey(client, nodeID)

	value := store.Get(key)

	var freeNode hub.NodeID
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &freeNode)

	return freeNode
}

func (k Keeper) RemoveFreeClient(ctx sdk.Context, nodeID hub.NodeID, client sdk.AccAddress) {
	clientKey := types.FreeNodesOfClientKey(client, nodeID)
	nodeKey := types.FreeClientOfNodeKey(nodeID, client)

	store := ctx.KVStore(k.freeClientKey)
	store.Delete(clientKey)
	store.Delete(nodeKey)
}
