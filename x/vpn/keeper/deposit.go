package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SubtractAndAddDeposit(ctx csdkTypes.Context, address csdkTypes.AccAddress,
	coin csdkTypes.Coin) (tags csdkTypes.Tags, err csdkTypes.Error) {

	return k.depositKeeper.SubtractAndAddDeposit(ctx, address, csdkTypes.Coins{coin})
}

func (k Keeper) AddAndSubtractDeposit(ctx csdkTypes.Context, address csdkTypes.AccAddress,
	coin csdkTypes.Coin) (tags csdkTypes.Tags, err csdkTypes.Error) {

	return k.depositKeeper.AddAndSubtractDeposit(ctx, address, csdkTypes.Coins{coin})
}
