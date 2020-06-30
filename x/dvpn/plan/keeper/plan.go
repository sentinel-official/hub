package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	"github.com/sentinel-official/hub/x/dvpn/plan/types"
)

func (k Keeper) SetPlansCount(ctx sdk.Context, count uint64) {
	key := types.PlansCountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlansCount(ctx sdk.Context) (count uint64) {
	store := k.Store(ctx)

	key := types.PlansCountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetPlan(ctx sdk.Context, plan types.Plan) {
	key := types.PlanKey(plan.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(plan)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlan(ctx sdk.Context, id uint64) (plan types.Plan, found bool) {
	store := k.Store(ctx)

	key := types.PlanKey(id)
	value := store.Get(key)
	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &plan)
	return plan, true
}

func (k Keeper) GetPlans(ctx sdk.Context) (items types.Plans) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Plan
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) SetPlanForProvider(ctx sdk.Context, address hub.ProvAddress, id uint64) {
	key := types.PlanForProviderKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress) (items types.Plans) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.PlanForProviderKeyPrefix(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &id)

		item, _ := k.GetPlan(ctx, id)
		items = append(items, item)
	}

	return items
}

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

	iter := sdk.KVStorePrefixIterator(store, types.NodeForPlanKeyPrefix(id))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var address hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &address)

		item, _ := k.GetNode(ctx, address)
		items = append(items, item)
	}

	return items
}
