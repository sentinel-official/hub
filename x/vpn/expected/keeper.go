package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) exported.Account
}
