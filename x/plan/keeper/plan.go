package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func (k *Keeper) SetCount(ctx sdk.Context, count uint64) {
	key := types.CountKey
	value := k.cdc.MustMarshalBinaryBare(&protobuf.UInt64Value{Value: count})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetCount(ctx sdk.Context) uint64 {
	store := k.Store(ctx)

	key := types.CountKey
	value := store.Get(key)
	if value == nil {
		return 0
	}

	var count protobuf.UInt64Value
	k.cdc.MustUnmarshalBinaryBare(value, &count)

	return count.GetValue()
}

func (k *Keeper) SetPlan(ctx sdk.Context, plan types.Plan) {
	key := types.PlanKey(plan.Id)
	value := k.cdc.MustMarshalBinaryBare(&plan)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plan types.Plan, found bool) {
	store := k.Store(ctx)

	key := types.PlanKey(id)
	value := store.Get(key)
	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &plan)
	return plan, true
}

func (k *Keeper) GetPlans(ctx sdk.Context, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Plan
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

func (k *Keeper) SetActivePlan(ctx sdk.Context, id uint64) {
	key := types.ActivePlanKey(id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteActivePlan(ctx sdk.Context, id uint64) {
	key := types.ActivePlanKey(id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetActivePlans(ctx sdk.Context, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.ActivePlanKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetPlan(ctx, types.IDFromStatusPlanKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k *Keeper) SetInactivePlan(ctx sdk.Context, id uint64) {
	key := types.InactivePlanKey(id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteInactivePlan(ctx sdk.Context, id uint64) {
	key := types.InactivePlanKey(id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetInactivePlans(ctx sdk.Context, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.InactivePlanKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetPlan(ctx, types.IDFromStatusPlanKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k *Keeper) SetActivePlanForProvider(ctx sdk.Context, address hubtypes.ProvAddress, id uint64) {
	key := types.ActivePlanForProviderKey(address, id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteActivePlanForProvider(ctx sdk.Context, address hubtypes.ProvAddress, id uint64) {
	store := k.Store(ctx)

	key := types.ActivePlanForProviderKey(address, id)
	store.Delete(key)
}

func (k *Keeper) GetActivePlansForProvider(ctx sdk.Context, address hubtypes.ProvAddress, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActivePlanForProviderKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetPlan(ctx, types.IDFromStatusPlanForProviderKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k *Keeper) SetInactivePlanForProvider(ctx sdk.Context, address hubtypes.ProvAddress, id uint64) {
	key := types.InactivePlanForProviderKey(address, id)
	value := k.cdc.MustMarshalBinaryBare(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteInactivePlanForProvider(ctx sdk.Context, address hubtypes.ProvAddress, id uint64) {
	store := k.Store(ctx)

	key := types.InactivePlanForProviderKey(address, id)
	store.Delete(key)
}

func (k *Keeper) GetInactivePlansForProvider(ctx sdk.Context, address hubtypes.ProvAddress, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetInactivePlanForProviderKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetPlan(ctx, types.IDFromStatusPlanForProviderKey(iter.Key()))
		items = append(items, item)
	})

	return items
}

func (k *Keeper) GetPlansForProvider(ctx sdk.Context, address hubtypes.ProvAddress, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetActivePlanForProviderKeyPrefix(address)),
			sdk.KVStorePrefixIterator(store, types.GetInactivePlanForProviderKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		item, _ := k.GetPlan(ctx, types.IDFromStatusPlanForProviderKey(iter.Key()))
		items = append(items, item)
	})

	return items
}
