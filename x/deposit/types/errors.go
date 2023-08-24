// DO NOT COVER

package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrorDepositNotFound     = sdkerrors.Register(ModuleName, 201, "deposit not found")
	ErrorInsufficientDeposit = sdkerrors.Register(ModuleName, 202, "insufficient deposit")
	ErrorInsufficientFunds   = sdkerrors.Register(ModuleName, 203, "insufficient funds")
)

func NewErrorDepositNotFound(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorDepositNotFound, "deposit for address %s does not exist", addr)
}

func NewErrorInsufficientDeposit(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInsufficientDeposit, "insufficient deposit for address %s", addr)
}

func NewErrorInsufficientFunds(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInsufficientFunds, "insufficient funds for address %s", addr)
}
