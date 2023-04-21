package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorDuplicateProvider = errors.Register(ModuleName, 201, "duplicate provider")
	ErrorProviderNotFound  = errors.Register(ModuleName, 202, "provider not found")
	ErrorInvalidStatus     = errors.Register(ModuleName, 203, "invalid status")
)

func NewErrorDuplicateProvider(addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorDuplicateProvider, "provider %s already exists", addr)
}

func NewErrorProviderNotFound(addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorInvalidStatus(status hubtypes.Status, addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for the provider %s", status, addr)
}
