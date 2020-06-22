package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func (k Keeper) SetPlansCountForProvider(ctx sdk.Context, address hub.ProvAddress, count uint64) {
	key := types.PlansCountForProviderKey(address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.PlanStore(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlansCountForProvider(ctx sdk.Context, address hub.ProvAddress) (count uint64) {
	store := k.PlanStore(ctx)

	key := types.PlansCountForProviderKey(address)
	value := store.Get(key)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetPlan(ctx sdk.Context, plan types.Plan) {
	key := types.PlanForProviderKey(plan.Provider, plan.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(plan)

	store := k.PlanStore(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlan(ctx sdk.Context, address hub.ProvAddress, i uint64) (plan types.Plan, found bool) {
	store := k.PlanStore(ctx)

	key := types.PlanForProviderKey(address, i)
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
