package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) Deposit(ctx sdk.Context) (v sdk.Coin) {
	k.params.Get(ctx, types.KeyDeposit, &v)
	return
}

func (k *Keeper) InactiveDuration(ctx sdk.Context) (v time.Duration) {
	k.params.Get(ctx, types.KeyInactiveDuration, &v)
	return
}

func (k *Keeper) MaxPrice(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMaxPrice, &v)
	return
}

func (k *Keeper) MinPrice(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMinPrice, &v)
	return
}

func (k *Keeper) StakingShare(ctx sdk.Context) (v sdk.Dec) {
	k.params.Get(ctx, types.KeyStakingShare, &v)
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
		k.StakingShare(ctx),
	)
}

func (k *Keeper) IsValidPrice(ctx sdk.Context, price sdk.Coins) bool {
	maxPrice := k.MaxPrice(ctx)
	for _, coin := range maxPrice {
		amount := price.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrice := k.MinPrice(ctx)
	for _, coin := range minPrice {
		amount := price.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsMaxPriceModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMaxPrice)
}

func (k *Keeper) IsMinPriceModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMinPrice)
}
