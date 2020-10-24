package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
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

func (k Keeper) SetActiveNode(ctx sdk.Context, address hub.NodeAddress) {
	key := types.ActiveNodeKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(address)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteActiveNode(ctx sdk.Context, address hub.NodeAddress) {
	key := types.ActiveNodeKey(address)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) GetActiveNodes(ctx sdk.Context) (items types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.ActiveNodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetInActiveNode(ctx sdk.Context, address hub.NodeAddress) {
	key := types.InActiveNodeKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(address)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInActiveNode(ctx sdk.Context, address hub.NodeAddress) {
	key := types.InActiveNodeKey(address)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) GetInActiveNodes(ctx sdk.Context) (items types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.InActiveNodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetActiveNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	key := types.ActiveNodeForProviderKey(p, n)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(n)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteActiveNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	store := k.Store(ctx)

	key := types.ActiveNodeForProviderKey(p, n)
	store.Delete(key)
}

func (k Keeper) GetActiveNodesForProvider(ctx sdk.Context, address hub.ProvAddress) (items types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetActiveNodeForProviderKeyPrefix(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetInActiveNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	key := types.InActiveNodeForProviderKey(p, n)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(n)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInActiveNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	store := k.Store(ctx)

	key := types.InActiveNodeForProviderKey(p, n)
	store.Delete(key)
}

func (k Keeper) GetInActiveNodesForProvider(ctx sdk.Context, address hub.ProvAddress) (items types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetInActiveNodeForProviderKeyPrefix(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetInActiveNodeAt(ctx sdk.Context, at time.Time, address hub.NodeAddress) {
	key := types.InActiveNodeAtKey(at, address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(address)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInActiveNodeAt(ctx sdk.Context, at time.Time, address hub.NodeAddress) {
	key := types.InActiveNodeAtKey(at, address)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) IterateInActiveNodesAt(ctx sdk.Context, at time.Time, fn func(index int, item types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.InActiveNodeAtKeyPrefix, sdk.PrefixEndBytes(types.GetInActiveNodeAtKeyPrefix(at)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		node, _ := k.GetNode(ctx, address)
		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}
