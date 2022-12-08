package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, from sdk.AccAddress, to string, coins sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, from string, to sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, from, to string, coins sdk.Coins) error
}
