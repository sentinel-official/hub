package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SupplyKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, address sdk.AccAddress, module string, coins sdk.Coins) sdk.Error
	SendCoinsFromModuleToAccount(ctx sdk.Context, module string, address sdk.AccAddress, coins sdk.Coins) sdk.Error
}
