package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v1/x/swap/types"
)

func (k *Keeper) SetSwap(ctx sdk.Context, swap types.Swap) {
	key := types.SwapKey(swap.GetTxHash())
	value := k.cdc.MustMarshal(&swap)

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

	k.cdc.MustUnmarshal(value, &swap)
	return swap, true
}

func (k *Keeper) HasSwap(ctx sdk.Context, txHash types.EthereumHash) bool {
	key := types.SwapKey(txHash)

	store := k.Store(ctx)
	return store.Has(key)
}

func (k *Keeper) GetSwaps(ctx sdk.Context) (items types.Swaps) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.SwapKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Swap
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}
