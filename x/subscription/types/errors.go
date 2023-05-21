// DO NOT COVER

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorInsufficientBytes = errors.Register(ModuleName, 201, "insufficient bytes")
	ErrorInvalidAllocation = errors.Register(ModuleName, 202, "invalid allocation")
	ErrorInvalidStatus     = errors.Register(ModuleName, 203, "invalid status")
	ErrorInvalidType       = errors.Register(ModuleName, 204, "invalid type")
	ErrorNotFound          = errors.Register(ModuleName, 205, "not found")
	ErrorUnauthorized      = errors.Register(ModuleName, 207, "unauthorized")
)

func NewErrorInsufficientBytes(id uint64, bytes sdk.Int) error {
	return errors.Wrapf(ErrorInsufficientBytes, "insufficient bytes %d for subscription %d", bytes, id)
}

func NewErrorInvalidAllocation(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorInvalidSubscriptionType(id uint64, t SubscriptionType) error {
	return errors.Wrapf(ErrorInvalidType, "invalid type %s for subscription %d", t, id)
}

func NewErrorAllocationNotFound(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorNotFound, "subscription allocation %d/%s does not exist", id, addr)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return errors.Wrapf(ErrorNotFound, "subscription %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return errors.Wrapf(ErrorNotFound, "plan %d does not exist", id)
}

func NewErrorInvalidPlanStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for plan %d", status, id)
}

func NewErrorPriceNotFound(denom string) error {
	return errors.Wrapf(ErrorNotFound, "price for denom %s does not exist", denom)
}

func NewErrorNodeNotFound(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorNotFound, "node %s does not exist", addr)
}

func NewErrorInvalidNodeStatus(addr hubtypes.NodeAddress, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for node %s", status, addr)
}
