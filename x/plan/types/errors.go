// DO NOT COVER

package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorNodeNotFound     = sdkerrors.Register(ModuleName, 201, "node not found")
	ErrorPlanNotFound     = sdkerrors.Register(ModuleName, 202, "plan not found")
	ErrorProviderNotFound = sdkerrors.Register(ModuleName, 203, "provider not found")
	ErrorUnauthorized     = sdkerrors.Register(ModuleName, 204, "unauthorised")
)

func NewErrorNodeNotFound(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorProviderNotFound(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorUnauthorized(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
