package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper Wrapper
func (k *Keeper) GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	return k.account.GetAccount(ctx, address)
}

func (k *Keeper) FundCommunityPool(ctx sdk.Context, from sdk.AccAddress, coin sdk.Coin) error {
	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), from)
}
