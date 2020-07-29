package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/node/types"
)

func (k Keeper) SetNode(ctx sdk.Context, node types.Node) {
	key := types.NodeKey(node.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(node)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) HasNode(ctx sdk.Context, address hub.NodeAddress) bool {
	store := k.Store(ctx)

	key := types.NodeKey(address)
	return store.Has(key)
}

func (k Keeper) GetNode(ctx sdk.Context, address hub.NodeAddress) (node types.Node, found bool) {
	store := k.Store(ctx)

	key := types.NodeKey(address)
	value := store.Get(key)
	if value == nil {
		return node, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &node)
	return node, true
}

func (k Keeper) SetNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	key := types.NodeForProviderKey(p, n)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(n)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	store := k.Store(ctx)

	key := types.NodeForProviderKey(p, n)
	store.Delete(key)
}

func (k Keeper) GetNodes(ctx sdk.Context) (items types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) GetNodesForProvider(ctx sdk.Context, address hub.ProvAddress) (items types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeForProviderByProviderKey(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	}

	return items
}

func (k Keeper) IterateNodes(ctx sdk.Context, f func(index int, item types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var node types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &node)

		if stop := f(i, node); stop {
			break
		}
		i++
	}
}

func (k Keeper) SetActiveNodeAt(ctx sdk.Context, at time.Time, address hub.NodeAddress) {
	key := types.ActiveNodeAtKey(at, address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(address)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteActiveNodeAt(ctx sdk.Context, at time.Time, address hub.NodeAddress) {
	key := types.ActiveNodeAtKey(at, address)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) IterateActiveNodes(ctx sdk.Context, end time.Time, f func(index int, item types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.ActiveNodeAtKeyPrefix, sdk.PrefixEndBytes(types.ActiveNodeAtByTimeKey(end)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		node, _ := k.GetNode(ctx, address)
		if stop := f(i, node); stop {
			break
		}
		i++
	}
}
