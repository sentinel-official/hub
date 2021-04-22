package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/swap/types"
)

func (k *Keeper) SetSwap(ctx sdk.Context, swap types.Swap) {
	key := types.SwapKey(swap.GetTxHash())
	value := k.cdc.MustMarshalBinaryBare(&swap)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) GetSwap(ctx sdk.Context, txHash types.EthereumHash) (swap types.Swap, found bool) {
	store := k.Store(ctx)

	key := types.SwapKey(txHash)
	value := store.Get(key)
	if value == nil {
		return swap, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &swap)
	return swap, true
}

func (k *Keeper) HasSwap(ctx sdk.Context, txHash types.EthereumHash) bool {
	key := types.SwapKey(txHash)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k *Keeper) GetSwaps(ctx sdk.Context, skip, limit int64) (items types.Swaps) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.SwapKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Swap
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}
