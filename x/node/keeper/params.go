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

func (k *Keeper) MaxGigabytePrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMaxGigabytePrices, &v)
	return
}

func (k *Keeper) MinGigabytePrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMinGigabytePrices, &v)
	return
}

func (k *Keeper) MaxHourlyPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMaxHourlyPrices, &v)
	return
}

func (k *Keeper) MinHourlyPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMinHourlyPrices, &v)
	return
}

func (k *Keeper) MaxLeaseHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyMaxLeaseHours, &v)
	return
}

func (k *Keeper) MinLeaseHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyMinLeaseHours, &v)
	return
}

func (k *Keeper) MaxLeaseGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyMaxLeaseGigabytes, &v)
	return
}

func (k *Keeper) MinLeaseGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, types.KeyMinLeaseGigabytes, &v)
	return
}

func (k *Keeper) RevenueShare(ctx sdk.Context) (v sdk.Dec) {
	k.params.Get(ctx, types.KeyRevenueShare, &v)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Deposit(ctx),
		k.InactiveDuration(ctx),
		k.MaxGigabytePrices(ctx),
		k.MinGigabytePrices(ctx),
		k.MaxHourlyPrices(ctx),
		k.MinHourlyPrices(ctx),
		k.MaxLeaseHours(ctx),
		k.MinLeaseHours(ctx),
		k.MaxLeaseGigabytes(ctx),
		k.MinLeaseGigabytes(ctx),
		k.RevenueShare(ctx),
	)
}

func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices sdk.Coins) bool {
	maxPrices := k.MaxGigabytePrices(ctx)
	for _, coin := range maxPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrices := k.MinGigabytePrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidHourlyPrices(ctx sdk.Context, prices sdk.Coins) bool {
	maxPrices := k.MaxHourlyPrices(ctx)
	for _, coin := range maxPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrices := k.MinHourlyPrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidLeaseHours(ctx sdk.Context, hours int64) bool {
	maxHours := k.MaxLeaseHours(ctx)
	if maxHours > 0 && hours > maxHours {
		return false
	}

	minHours := k.MinLeaseHours(ctx)
	if minHours > 0 && hours < minHours {
		return false
	}

	return true
}

func (k *Keeper) IsValidLeaseGigabytes(ctx sdk.Context, gigabytes int64) bool {
	maxGigabytes := k.MaxLeaseGigabytes(ctx)
	if maxGigabytes > 0 && gigabytes > maxGigabytes {
		return false
	}

	minGigabytes := k.MinLeaseGigabytes(ctx)
	if minGigabytes > 0 && gigabytes < minGigabytes {
		return false
	}

	return true
}

func (k *Keeper) IsMaxGigabytePricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMaxGigabytePrices)
}

func (k *Keeper) IsMinGigabytePricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMinGigabytePrices)
}

func (k *Keeper) IsMaxHourlyPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMaxHourlyPrices)
}

func (k *Keeper) IsMinHourlyPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMinHourlyPrices)
}
