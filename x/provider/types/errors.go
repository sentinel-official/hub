package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorDuplicate     = errors.Register(ModuleName, 201, "duplicate")
	ErrorNotFound      = errors.Register(ModuleName, 202, "not found")
	ErrorInvalidStatus = errors.Register(ModuleName, 203, "invalid status")
)

func NewErrorDuplicateProvider(addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorDuplicate, "provider %s already exists", addr)
}

func NewErrorProviderNotFound(addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorNotFound, "provider %s does not exist", addr)
}

func NewErrorInvalidProviderStatus(addr hubtypes.ProvAddress, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for the provider %s", status, addr)
}
