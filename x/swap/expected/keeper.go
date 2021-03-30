package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) exported.Account
}

type SupplyKeeper interface {
	MintCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, name string, address sdk.AccAddress, coins sdk.Coins) error
}
