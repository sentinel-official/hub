package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/mint/types"
)

func (k *Keeper) SetInflation(ctx sdk.Context, inflation types.Inflation) {
	var (
		store = k.Store(ctx)
		key   = types.InflationKey(inflation.Timestamp)
		value = k.cdc.MustMarshal(&inflation)
	)

	store.Set(key, value)
}

func (k *Keeper) GetInflation(ctx sdk.Context, t time.Time) (inflation types.Inflation, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.InflationKey(t)
		value = store.Get(key)
	)

	if value == nil {
		return inflation, false
	}

	k.cdc.MustUnmarshal(value, &inflation)
	return inflation, true
}

func (k *Keeper) DeleteInflation(ctx sdk.Context, t time.Time) {
	var (
		store = k.Store(ctx)
		key   = types.InflationKey(t)
	)

	store.Delete(key)
}

func (k *Keeper) GetInflations(ctx sdk.Context, skip, limit int64) (items []types.Inflation) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.InflationKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Inflation
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

func (k *Keeper) IterateInflations(ctx sdk.Context, fn func(index int, item types.Inflation) (stop bool)) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.InflationKeyPrefix)
	)

	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item types.Inflation
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}
