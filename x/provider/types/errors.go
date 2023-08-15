// DO NOT COVER

package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateProvider = errors.Register(ModuleName, 201, "duplicate provider")
	ErrorProviderNotFound  = errors.Register(ModuleName, 202, "provider not found")
)

func NewErrorDuplicateProvider(addr interface{}) error {
	return errors.Wrapf(ErrorDuplicateProvider, "provider %s already exist", addr)
}

func NewErrorProviderNotFound(addr interface{}) error {
	return errors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}
