package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k *Keeper) SetQuota(ctx sdk.Context, id uint64, quota types.Quota) {
	var (
		store = k.Store(ctx)
		key   = types.QuotaKey(id, quota.GetAccountAddress())
		value = k.cdc.MustMarshal(&quota)
	)

	store.Set(key, value)
}

func (k *Keeper) GetQuota(ctx sdk.Context, id uint64, addr sdk.AccAddress) (quota types.Quota, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.QuotaKey(id, addr)
		value = store.Get(key)
	)

	if value == nil {
		return quota, false
	}

	k.cdc.MustUnmarshal(value, &quota)
	return quota, true
}

func (k *Keeper) HasQuota(ctx sdk.Context, id uint64, addr sdk.AccAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.QuotaKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteQuota(ctx sdk.Context, id uint64, addr sdk.AccAddress) {
	var (
		store = k.Store(ctx)
		key   = types.QuotaKey(id, addr)
	)

	store.Delete(key)
}

func (k *Keeper) GetQuotas(ctx sdk.Context, id uint64) (items types.Quotas) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetQuotaKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Quota
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateQuotas(ctx sdk.Context, id uint64, fn func(index int, item types.Quota) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetQuotaKeyPrefix(id))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var quota types.Quota
		k.cdc.MustUnmarshal(iter.Value(), &quota)

		if stop := fn(i, quota); stop {
			break
		}
		i++
	}
}
