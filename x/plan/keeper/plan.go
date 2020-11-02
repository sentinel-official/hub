package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func (k Keeper) SetPlansCount(ctx sdk.Context, count uint64) {
	key := types.CountKey
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlansCount(ctx sdk.Context) (count uint64) {
	store := k.Store(ctx)

	key := types.CountKey
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

func (k Keeper) GetPlans(ctx sdk.Context, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Plan
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

func (k Keeper) SetPlanForProvider(ctx sdk.Context, address hub.ProvAddress, id uint64) {
	key := types.PlanForProviderKey(address, id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(id)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress, skip, limit int) (items types.Plans) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetPlanForProviderKeyPrefix(address)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var id uint64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &id)

		item, _ := k.GetPlan(ctx, id)
		items = append(items, item)
	})

	return items
}
