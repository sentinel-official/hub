package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) SetActiveNode(ctx sdk.Context, node types.Node) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(node.GetAddress())
		value = k.cdc.MustMarshal(&node)
	)

	store.Set(key, value)
}

func (k *Keeper) HasActiveNode(ctx sdk.Context, addr hubtypes.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) GetActiveNode(ctx sdk.Context, addr hubtypes.NodeAddress) (v types.Node, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) DeleteActiveNode(ctx sdk.Context, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(addr)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactiveNode(ctx sdk.Context, node types.Node) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(node.GetAddress())
		value = k.cdc.MustMarshal(&node)
	)

	store.Set(key, value)
}

func (k *Keeper) HasInactiveNode(ctx sdk.Context, addr hubtypes.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) GetInactiveNode(ctx sdk.Context, addr hubtypes.NodeAddress) (v types.Node, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) DeleteInactiveNode(ctx sdk.Context, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(addr)
	)

	store.Delete(key)
}

func (k *Keeper) SetNode(ctx sdk.Context, node types.Node) {
	switch node.Status {
	case hubtypes.StatusActive:
		k.SetActiveNode(ctx, node)
	case hubtypes.StatusInactive:
		k.SetInactiveNode(ctx, node)
	default:
		panic(fmt.Errorf("invalid status for the node %v", node))
	}
}

func (k *Keeper) HasNode(ctx sdk.Context, addr hubtypes.NodeAddress) bool {
	return k.HasActiveNode(ctx, addr) ||
		k.HasInactiveNode(ctx, addr)
}

func (k *Keeper) GetNode(ctx sdk.Context, addr hubtypes.NodeAddress) (node types.Node, found bool) {
	node, found = k.GetActiveNode(ctx, addr)
	if found {
		return
	}

	node, found = k.GetInactiveNode(ctx, addr)
	if found {
		return
	}

	return node, false
}

func (k *Keeper) GetNodes(ctx sdk.Context) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Node
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateNodes(ctx sdk.Context, fn func(index int, item types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var node types.Node
		k.cdc.MustUnmarshal(iter.Value(), &node)

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetInactiveNodeAt(ctx sdk.Context, at time.Time, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeAtKey(at, addr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteInactiveNodeAt(ctx sdk.Context, at time.Time, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeAtKey(at, addr)
	)

	store.Delete(key)
}

func (k *Keeper) IterateInactiveNodesAt(ctx sdk.Context, at time.Time, fn func(index int, item types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.InactiveNodeAtKeyPrefix, sdk.PrefixEndBytes(types.GetInactiveNodeAtKeyPrefix(at)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var (
			key     = iter.Key()
			node, _ = k.GetNode(ctx, types.AddressFromInactiveNodeAtKey(key))
		)

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}
