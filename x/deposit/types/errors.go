// DO NOT COVER

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInsufficientFunds = errors.Register(ModuleName, 201, "insufficient funds")
	ErrorNotFound          = errors.Register(ModuleName, 202, "not found")
)

func NewErrorInsufficientFunds(addr sdk.AccAddress, coins sdk.Coins) error {
	return errors.Wrapf(ErrorInsufficientFunds, "insufficient deposit funds %s for address %s", coins, addr)
}

func NewErrorDepositNotFound(addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorNotFound, "deposit for address %s does not exist", addr)
}
