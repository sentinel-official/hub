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

func (k Keeper) GetNodes(ctx sdk.Context, skip, limit int) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	})

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

func (k Keeper) GetActiveNodes(ctx sdk.Context, skip, limit int) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.ActiveNodeKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetInactiveNode(ctx sdk.Context, address hub.NodeAddress) {
	key := types.InactiveNodeKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(address)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInactiveNode(ctx sdk.Context, address hub.NodeAddress) {
	key := types.InactiveNodeKey(address)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) GetInactiveNodes(ctx sdk.Context, skip, limit int) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.InactiveNodeKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	})

	return items
}

func (k Keeper) GetNodesForProvider(ctx sdk.Context, address hub.ProvAddress, skip, limit int) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActiveNodeForProviderKeyPrefix(address)),
			sdk.KVStorePrefixIterator(store, types.GetInactiveNodeForProviderKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	})

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

func (k Keeper) GetActiveNodesForProvider(ctx sdk.Context, address hub.ProvAddress, skip, limit int) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActiveNodeForProviderKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetInactiveNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	key := types.InactiveNodeForProviderKey(p, n)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(n)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInactiveNodeForProvider(ctx sdk.Context, p hub.ProvAddress, n hub.NodeAddress) {
	store := k.Store(ctx)

	key := types.InactiveNodeForProviderKey(p, n)
	store.Delete(key)
}

func (k Keeper) GetInactiveNodesForProvider(ctx sdk.Context, address hub.ProvAddress, skip, limit int) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetInactiveNodeForProviderKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetInactiveNodeAt(ctx sdk.Context, at time.Time, address hub.NodeAddress) {
	key := types.InactiveNodeAtKey(at, address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(address)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) DeleteInactiveNodeAt(ctx sdk.Context, at time.Time, address hub.NodeAddress) {
	key := types.InactiveNodeAtKey(at, address)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k Keeper) IterateInactiveNodesAt(ctx sdk.Context, at time.Time, fn func(index int, item types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.InactiveNodeAtKeyPrefix, sdk.PrefixEndBytes(types.GetInactiveNodeAtKeyPrefix(at)))
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
