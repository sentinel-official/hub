package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) SetQuota(ctx sdk.Context, id uint64, quota types.Quota) {
	key := types.QuotaKey(id, quota.GetAddress())
	value := k.cdc.MustMarshalBinaryBare(&quota)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) (quota types.Quota, found bool) {
	store := k.Store(ctx)

	key := types.QuotaKey(id, address)
	value := store.Get(key)
	if value == nil {
		return quota, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &quota)
	return quota, true
}

func (k *Keeper) HasQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) bool {
	key := types.QuotaKey(id, address)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k *Keeper) DeleteQuota(ctx sdk.Context, id uint64, address sdk.AccAddress) {
	key := types.QuotaKey(id, address)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetQuotas(ctx sdk.Context, id uint64, skip, limit int64) (items types.Quotas) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.GetQuotaKeyPrefix(id)),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Quota
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

func (k *Keeper) IterateQuotas(ctx sdk.Context, id uint64, fn func(index int, item types.Quota) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetQuotaKeyPrefix(id))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var quota types.Quota
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &quota)

		if stop := fn(i, quota); stop {
			break
		}
		i++
	}
}
