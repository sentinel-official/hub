package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sentinel-official/hub/v1/x/swap/types"
)

func (k *Keeper) GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	return k.account.GetAccount(ctx, address)
}

func (k *Keeper) MintCoin(ctx sdk.Context, coin sdk.Coin) error {
	return k.bank.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromModuleToAccount(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) error {
	return k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(coin))
}
