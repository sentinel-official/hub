package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	SetAccount(ctx sdk.Context, account authtypes.AccountI)
	IterateAccounts(ctx sdk.Context, fn func(account authtypes.AccountI) bool)
}

type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}
