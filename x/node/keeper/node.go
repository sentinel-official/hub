package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/v1/types"
	"github.com/sentinel-official/hub/v1/x/node/types"
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
		panic(fmt.Errorf("failed to set the node %v", node))
	}
}

func (k *Keeper) HasNode(ctx sdk.Context, addr hubtypes.NodeAddress) bool {
	return k.HasActiveNode(ctx, addr) || k.HasInactiveNode(ctx, addr)
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

func (k *Keeper) SetNodeForInactiveAt(ctx sdk.Context, at time.Time, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForInactiveAtKey(at, addr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteNodeForInactiveAt(ctx sdk.Context, at time.Time, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForInactiveAtKey(at, addr)
	)

	store.Delete(key)
}

func (k *Keeper) IterateNodesForInactiveAt(ctx sdk.Context, at time.Time, fn func(index int, item types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.NodeForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetNodeForInactiveAtKeyPrefix(at)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		node, found := k.GetNode(ctx, types.AddressFromNodeForInactiveAtKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("node for inactive at key %X does not exist", iter.Key()))
		}

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForPlanKey(id, addr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.NodeForPlanKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, addr hubtypes.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForPlanKey(id, addr)
	)

	store.Delete(key)
}

func (k *Keeper) GetNodesForPlan(ctx sdk.Context, id uint64) (items types.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetNodeForPlanKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		item, found := k.GetNode(ctx, types.AddressFromNodeForPlanKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("node for plan key %X does not exist", iter.Key()))
		}

		items = append(items, item)
	}

	return items
}
