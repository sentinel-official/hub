package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func (k Keeper) SetFreeNodeOfClient(ctx sdk.Context, client sdk.AccAddress, id hub.NodeID) {
	key := types.FreeNodesOfClientKey(client, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) SetFreeClientOfNode(ctx sdk.Context, id hub.NodeID, client sdk.AccAddress) {
	key := types.FreeClientOfNodeKey(id, client)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(client)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) GetFreeClientsOfNode(ctx sdk.Context, nodeID hub.NodeID) (freeClients []sdk.AccAddress) {
	store := ctx.KVStore(k.nodeKey)
	
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

func (k Keeper) GetFreeClientOfNode(ctx sdk.Context, nodeID hub.NodeID, client sdk.AccAddress) (address sdk.AccAddress, found bool) {
	store := ctx.KVStore(k.nodeKey)
	
	key := types.FreeClientOfNodeKey(nodeID, client)
	
	value := store.Get(key)
	if value == nil {
		return nil, false
	}
	
	var freeClient sdk.AccAddress
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &freeClient)
	
	return freeClient, true
}

func (k Keeper) GetFreeNodesOfClient(ctx sdk.Context, client sdk.AccAddress) (freeNodes []hub.NodeID) {
	store := ctx.KVStore(k.nodeKey)
	
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

func (k Keeper) GetFreeNodeOfClient(ctx sdk.Context, client sdk.AccAddress, nodeID hub.NodeID) (id hub.NodeID, found bool) {
	store := ctx.KVStore(k.nodeKey)
	
	key := types.FreeNodesOfClientKey(client, nodeID)
	
	value := store.Get(key)
	if value == nil {
		return nil, false
	}
	
	var freeNode hub.NodeID
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &freeNode)
	
	return freeNode, true
}

func (k Keeper) RemoveFreeClientOfNode(ctx sdk.Context, nodeID hub.NodeID, client sdk.AccAddress) {
	clientKey := types.FreeNodesOfClientKey(client, nodeID)
	nodeKey := types.FreeClientOfNodeKey(nodeID, client)
	
	store := ctx.KVStore(k.nodeKey)
	store.Delete(clientKey)
	store.Delete(nodeKey)
}

func (k Keeper) SetFreeClient(ctx sdk.Context, freeClient types.FreeClient) {
	key := types.GetFreeClientKey(freeClient.NodeID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(freeClient)
	
	store := ctx.KVStore(k.nodeKey)
	store.Set(key, value)
}

func (k Keeper) GetFreeClients(ctx sdk.Context) (freeClients []types.FreeClient) {
	store := ctx.KVStore(k.nodeKey)
	
	iter := sdk.KVStorePrefixIterator(store, types.FreeClientKey)
	defer iter.Close()
	
	for ; iter.Valid(); iter.Next() {
		var freeClient types.FreeClient
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &freeClient)
		freeClients = append(freeClients, freeClient)
	}
	
	return freeClients
}

func (k Keeper) RemoveFreeClient(ctx sdk.Context, nodeID hub.NodeID) {
	freeClientKey := types.GetFreeClientKey(nodeID)
	
	store := ctx.KVStore(k.nodeKey)
	store.Delete(freeClientKey)
}
