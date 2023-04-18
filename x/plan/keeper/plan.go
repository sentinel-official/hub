package keeper

import (
	"fmt"
	protobuf "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func (k *Keeper) SetActivePlan(ctx sdk.Context, plan types.Plan) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(plan.ID)
		value = k.cdc.MustMarshal(&plan)
	)

	store.Set(key, value)
}

func (k *Keeper) HasActivePlan(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetActivePlan(ctx sdk.Context, id uint64) (plan types.Plan, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshal(value, &plan)
	return plan, true
}

func (k *Keeper) DeleteActivePlan(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactivePlan(ctx sdk.Context, plan types.Plan) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(plan.ID)
		value = k.cdc.MustMarshal(&plan)
	)

	store.Set(key, value)
}

func (k *Keeper) HasInactivePlan(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetInactivePlan(ctx sdk.Context, id uint64) (plan types.Plan, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshal(value, &plan)
	return plan, true
}

func (k *Keeper) DeleteInactivePlan(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) SetPlan(ctx sdk.Context, plan types.Plan) {
	switch plan.Status {
	case hubtypes.StatusActive:
		k.SetActivePlan(ctx, plan)
	case hubtypes.StatusInactive:
		k.SetInactivePlan(ctx, plan)
	default:
		panic(fmt.Errorf("invalid status for the plan %v", plan))
	}
}

func (k *Keeper) HasPlan(ctx sdk.Context, id uint64) bool {
	return k.HasActivePlan(ctx, id) ||
		k.HasInactivePlan(ctx, id)
}

func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plan types.Plan, found bool) {
	plan, found = k.GetActivePlan(ctx, id)
	if found {
		return
	}

	plan, found = k.GetInactivePlan(ctx, id)
	if found {
		return
	}

	return plan, false
}

func (k *Keeper) GetPlans(ctx sdk.Context) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Plan
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetActivePlanForProvider(ctx sdk.Context, addr hubtypes.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanForProviderKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteActivePlanForProvider(ctx sdk.Context, addr hubtypes.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.ActivePlanForProviderKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactivePlanForProvider(ctx sdk.Context, addr hubtypes.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanForProviderKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteInactivePlanForProvider(ctx sdk.Context, addr hubtypes.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.InactivePlanForProviderKey(addr, id)
	)

	store.Delete(key)
}
