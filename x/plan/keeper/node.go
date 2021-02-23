package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func (k Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress) {
	key := types.NodeForPlanKey(id, address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(true)

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

func (k Keeper) GetNodesForPlan(ctx sdk.Context, id uint64, skip, limit int) (items node.Nodes) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetNodeForPlanKeyPrefix(id)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetNode(ctx, types.AddressFromNodeForPlanKey(iter.Key()))
		items = append(items, item)
	})

	return items
}
