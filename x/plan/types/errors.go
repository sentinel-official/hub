package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorProviderNotFound = errors.Register(ModuleName, 201, "provider not found")
	ErrorPlanNotFound     = errors.Register(ModuleName, 202, "plan not found")
	ErrorUnauthorized     = errors.Register(ModuleName, 203, "unauthorized")
)

func NewErrorProviderNotFound(addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return errors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorised", addr)
}
