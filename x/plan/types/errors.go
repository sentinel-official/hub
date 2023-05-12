// DO NOT COVER

package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorNotFound      = errors.Register(ModuleName, 201, "not found")
	ErrorUnauthorized  = errors.Register(ModuleName, 202, "unauthorized")
	ErrorInvalidStatus = errors.Register(ModuleName, 203, "invalid status")
)

func NewErrorProviderNotFound(addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorNotFound, "provider %s does not exist", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return errors.Wrapf(ErrorNotFound, "plan %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorised", addr)
}

func NewErrorNodeNotFound(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorNotFound, "node %s does not exist", addr)
}

func NewErrorInvalidPlanStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for plan %d", status, id)
}

func NewErrorPriceNotFound(denom string) error {
	return errors.Wrapf(ErrorNotFound, "price for denom %s does not exist", denom)
}
