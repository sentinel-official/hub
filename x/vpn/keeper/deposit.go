package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.deposit.Add(ctx, address, sdk.Coins{coin})
}

func (k Keeper) SubtractDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.deposit.Subtract(ctx, address, sdk.Coins{coin})
}

func (k Keeper) SendDeposit(ctx sdk.Context, from, toAddress sdk.AccAddress, coin sdk.Coin) sdk.Error {
	return k.deposit.SendCoinsFromDepositToAccount(ctx, from, toAddress, sdk.Coins{coin})
}
