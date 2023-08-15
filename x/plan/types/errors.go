// DO NOT COVER

package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")

	ErrorNodeNotFound     = errors.Register(ModuleName, 201, "node not found")
	ErrorPlanNotFound     = errors.Register(ModuleName, 202, "plan not found")
	ErrorProviderNotFound = errors.Register(ModuleName, 203, "provider not found")
	ErrorUnauthorized     = errors.Register(ModuleName, 204, "unauthorised")
)

func NewErrorNodeNotFound(addr interface{}) error {
	return errors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return errors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorProviderNotFound(addr interface{}) error {
	return errors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorUnauthorized(addr interface{}) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
