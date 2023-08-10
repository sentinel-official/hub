// DO NOT COVER

package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorDepositNotFound     = errors.Register(ModuleName, 201, "deposit not found")
	ErrorInsufficientDeposit = errors.Register(ModuleName, 202, "insufficient deposit")
	ErrorInsufficientFunds   = errors.Register(ModuleName, 203, "insufficient funds")
)

func NewErrorDepositNotFound(addr interface{}) error {
	return errors.Wrapf(ErrorDepositNotFound, "deposit for address %s does not exist", addr)
}

func NewErrorInsufficientDeposit(addr interface{}) error {
	return errors.Wrapf(ErrorInsufficientDeposit, "insufficient deposit for address %s", addr)
}

func NewErrorInsufficientFunds(addr interface{}) error {
	return errors.Wrapf(ErrorInsufficientFunds, "insufficient funds for address %s", addr)
}
