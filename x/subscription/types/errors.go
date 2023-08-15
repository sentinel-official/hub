// DO NOT COVER

package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")

	ErrorAllocationNotFound        = errors.Register(ModuleName, 201, "allocation not found")
	ErrorInsufficientBytes         = errors.Register(ModuleName, 202, "insufficient bytes")
	ErrorInvalidAllocation         = errors.Register(ModuleName, 203, "invalid allocation")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 204, "invalid node status")
	ErrorInvalidPlanStatus         = errors.Register(ModuleName, 205, "invalid plan status")
	ErrorInvalidSubscription       = errors.Register(ModuleName, 206, "invalid subscription")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 207, "invalid subscription status")
	ErrorNodeNotFound              = errors.Register(ModuleName, 208, "node not found")
	ErrorPayoutNotFound            = errors.Register(ModuleName, 209, "payout not found")
	ErrorPlanNotFound              = errors.Register(ModuleName, 210, "plan not found")
	ErrorPriceNotFound             = errors.Register(ModuleName, 211, "price does not exist")
	ErrorSubscriptionNotFound      = errors.Register(ModuleName, 212, "subscription not found")
	ErrorUnauthorized              = errors.Register(ModuleName, 213, "unauthorised")
)

func NewErrorAllocationNotFound(id uint64, addr interface{}) error {
	return errors.Wrapf(ErrorAllocationNotFound, "allocation %d/%s does not exist", id, addr)
}

func NewErrorInsufficientBytes(id uint64, addr interface{}) error {
	return errors.Wrapf(ErrorInsufficientBytes, "insufficient bytes for allocation %d/%s", id, addr)
}

func NewErrorInvalidAllocation(id uint64, addr interface{}) error {
	return errors.Wrapf(ErrorInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

func NewErrorInvalidNodeStatus(addr interface{}, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

func NewErrorInvalidPlanStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidPlanStatus, "invalid status %s for plan %d", status, id)
}

func NewErrorInvalidSubscription(id uint64) error {
	return errors.Wrapf(ErrorInvalidSubscription, "invalid subscription %d", id)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidSubscriptionStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorNodeNotFound(addr interface{}) error {
	return errors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPayoutNotFound(id uint64) error {
	return errors.Wrapf(ErrorPayoutNotFound, "payout %d does not exist", id)
}

func NewErrorPlanNotFound(id uint64) error {
	return errors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorPriceNotFound(denom string) error {
	return errors.Wrapf(ErrorPriceNotFound, "price for denom %s does not exist", denom)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return errors.Wrapf(ErrorSubscriptionNotFound, "subscription %d does not exist", id)
}

func NewErrorUnauthorized(addr interface{}) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
