package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func (k Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) {
	key := types.NodeForPlanKey(id, address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(address)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) bool {
	store := k.Store(ctx)

	key := types.NodeForPlanKey(id, address)
	return store.Has(key)
}

func (k Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) {
	store := k.Store(ctx)

	key := types.NodeForPlanKey(id, address)
	store.Delete(key)
}

func (k Keeper) GetNodesForPlan(ctx sdk.Context, id uint64) (items node.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetNodeForPlanKeyPrefix(id))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	}

	return items
}
