package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/swap/types"
)

func (k Keeper) SetSwap(ctx sdk.Context, swap types.Swap) {
	key := types.SwapKey(swap.TxHash)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(swap)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetSwap(ctx sdk.Context, txHash types.EthereumHash) (swap types.Swap, found bool) {
	store := k.Store(ctx)

	key := types.SwapKey(txHash)
	value := store.Get(key)
	if value == nil {
		return swap, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &swap)
	return swap, true
}

func (k Keeper) HasSwap(ctx sdk.Context, txHash types.EthereumHash) bool {
	key := types.SwapKey(txHash)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k Keeper) GetSwaps(ctx sdk.Context, skip, limit int) (items types.Swaps) {
	var (
		store = k.Store(ctx)
		iter  = hub.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.SwapKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Swap
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}
