package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddDeposit(ctx csdkTypes.Context, address csdkTypes.AccAddress,
	coin csdkTypes.Coin) (tags csdkTypes.Tags, err csdkTypes.Error) {

	return k.depositKeeper.Add(ctx, address, csdkTypes.Coins{coin})
}

func (k Keeper) SubtractDeposit(ctx csdkTypes.Context, address csdkTypes.AccAddress,
	coin csdkTypes.Coin) (tags csdkTypes.Tags, err csdkTypes.Error) {

	return k.depositKeeper.Subtract(ctx, address, csdkTypes.Coins{coin})
}

func (k Keeper) SendDepositTo(ctx csdkTypes.Context, from, toAddress csdkTypes.AccAddress,
	coin csdkTypes.Coin) (tags csdkTypes.Tags, err csdkTypes.Error) {

	return k.depositKeeper.SendTo(ctx, from, toAddress, csdkTypes.Coins{coin})
}
