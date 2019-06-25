package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddDeposit(ctx sdk.Context, address sdk.AccAddress,
	coin sdk.Coin) (tags sdk.Tags, err sdk.Error) {

	return k.depositKeeper.Add(ctx, address, sdk.Coins{coin})
}

func (k Keeper) SubtractDeposit(ctx sdk.Context, address sdk.AccAddress,
	coin sdk.Coin) (tags sdk.Tags, err sdk.Error) {

	return k.depositKeeper.Subtract(ctx, address, sdk.Coins{coin})
}

func (k Keeper) SendDeposit(ctx sdk.Context, from, toAddress sdk.AccAddress,
	coin sdk.Coin) (tags sdk.Tags, err sdk.Error) {

	return k.depositKeeper.Send(ctx, from, toAddress, sdk.Coins{coin})
}
