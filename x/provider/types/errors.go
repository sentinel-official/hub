// DO NOT COVER

package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateProvider = sdkerrors.Register(ModuleName, 201, "duplicate provider")
	ErrorProviderNotFound  = sdkerrors.Register(ModuleName, 202, "provider not found")
)

func NewErrorDuplicateProvider(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorDuplicateProvider, "provider %s already exist", addr)
}

func NewErrorProviderNotFound(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}
