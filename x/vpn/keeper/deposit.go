package keeper

import (
	csdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddDeposit(ctx csdk.Context, address csdk.AccAddress,
	coin csdk.Coin) (tags csdk.Tags, err csdk.Error) {

	return k.depositKeeper.Add(ctx, address, csdk.Coins{coin})
}

func (k Keeper) SubtractDeposit(ctx csdk.Context, address csdk.AccAddress,
	coin csdk.Coin) (tags csdk.Tags, err csdk.Error) {

	return k.depositKeeper.Subtract(ctx, address, csdk.Coins{coin})
}

func (k Keeper) SendDeposit(ctx csdk.Context, from, toAddress csdk.AccAddress,
	coin csdk.Coin) (tags csdk.Tags, err csdk.Error) {

	return k.depositKeeper.Send(ctx, from, toAddress, csdk.Coins{coin})
}
