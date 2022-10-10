package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) Deposit(ctx sdk.Context) (deposit sdk.Coin) {
	k.params.Get(ctx, types.KeyDeposit, &deposit)
	return
}

func (k *Keeper) InactiveDuration(ctx sdk.Context) (duration time.Duration) {
	k.params.Get(ctx, types.KeyInactiveDuration, &duration)
	return
}

func (k *Keeper) MaxPrice(ctx sdk.Context) (price sdk.Coins) {
	k.params.Get(ctx, types.KeyMaxPrice, &price)
	return
}

func (k *Keeper) MinPrice(ctx sdk.Context) (price sdk.Coins) {
	k.params.Get(ctx, types.KeyMinPrice, &price)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Deposit(ctx),
		k.InactiveDuration(ctx),
		k.MaxPrice(ctx),
		k.MinPrice(ctx),
	)
}

func (k *Keeper) IsValidPrice(ctx sdk.Context, price sdk.Coins) bool {
	var (
		maxPrice = k.MaxPrice(ctx)
		minPrice = k.MinPrice(ctx)
	)

	return price.IsAllLTE(maxPrice) &&
		price.IsAllGTE(minPrice)
}

func (k *Keeper) IsMaxPriceModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMaxPrice)
}

func (k *Keeper) IsMinPriceModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMinPrice)
}

func (k *Keeper) DeleteTransientKeyMaxPrice(ctx sdk.Context) {
	var (
		tkey  = sdk.NewTransientStoreKey(paramstypes.TStoreKey)
		store = ctx.TransientStore(tkey)
	)

	store.Delete(types.KeyMaxPrice)
}

func (k *Keeper) DeleteTransientKeyMinPrice(ctx sdk.Context) {
	var (
		tkey  = sdk.NewTransientStoreKey(paramstypes.TStoreKey)
		store = ctx.TransientStore(tkey)
	)

	store.Delete(types.KeyMinPrice)
}
