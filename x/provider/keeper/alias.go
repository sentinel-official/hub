package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) FundCommunityPool(ctx sdk.Context, from sdk.AccAddress, coin sdk.Coin) error {
	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), from)
}
