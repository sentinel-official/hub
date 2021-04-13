package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/swap/types"
)

func (k Keeper) MintCoin(ctx sdk.Context, coin sdk.Coin) error {
	return k.bank.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
}

func (k Keeper) SendCoinFromModuleToAccount(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) error {
	return k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(coin))
}
