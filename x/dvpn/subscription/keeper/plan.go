package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func (k Keeper) SetPlansCount(ctx sdk.Context, address hub.ProvAddress, count uint64) {
	key := types.PlansCountKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.PlanStore(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlansCount(ctx sdk.Context, address hub.ProvAddress) (count uint64) {
	store := k.PlanStore(ctx)

	key := types.PlansCountKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetPlan(ctx sdk.Context, plan types.Plan) {
	key := types.PlanKey(plan.Provider, plan.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(plan)

	store := k.PlanStore(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlan(ctx sdk.Context, address hub.ProvAddress, i uint64) (plan types.Plan, found bool) {
	store := k.PlanStore(ctx)

	key := types.PlanKey(address, i)
	value := store.Get(key)
	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &plan)
	return plan, true
}

func (k Keeper) GetPlans(ctx sdk.Context) (plans types.Plans) {
	store := k.PlanStore(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var plan types.Plan
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &plan)
		plans = append(plans, plan)
	}

	return plans
}

func (k Keeper) GetPlansOfProvider(ctx sdk.Context, address hub.ProvAddress) (plans types.Plans) {
	store := k.PlanStore(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.PlanForProviderKeyPrefix(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var plan types.Plan
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &plan)
		plans = append(plans, plan)
	}

	return plans
}

func (k Keeper) SetNodeAddressForPlan(ctx sdk.Context, pa hub.ProvAddress, i uint64, na hub.NodeAddress) {
	key := types.NodeAddressKey(pa, i, na)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(na)

	store := k.PlanStore(ctx)
	store.Set(key, value)
}

func (k Keeper) HasNodeAddressForPlan(ctx sdk.Context, pa hub.ProvAddress, i uint64, na hub.NodeAddress) bool {
	store := k.PlanStore(ctx)

	key := types.NodeAddressKey(pa, i, na)
	value := store.Get(key)

	return value != nil
}

func (k Keeper) DeleteNodeAddressForPlan(ctx sdk.Context, pa hub.ProvAddress, i uint64, na hub.NodeAddress) {
	store := k.PlanStore(ctx)

	key := types.NodeAddressKey(pa, i, na)
	store.Delete(key)
}

func (k Keeper) GetNodesForPlan(ctx sdk.Context, address hub.ProvAddress, i uint64) (nodes node.Nodes) {
	store := k.PlanStore(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeAddressForPlanKeyPrefix(address, i))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var na hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &na)

		n, _ := k.GetNode(ctx, na)
		nodes = append(nodes, n)
	}

	return nodes
}
