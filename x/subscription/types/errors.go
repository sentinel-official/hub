// DO NOT COVER

package types

import (
	sdkerrors "cosmossdk.io/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorAllocationNotFound        = sdkerrors.Register(ModuleName, 201, "allocation not found")
	ErrorInsufficientBytes         = sdkerrors.Register(ModuleName, 202, "insufficient bytes")
	ErrorInvalidAllocation         = sdkerrors.Register(ModuleName, 203, "invalid allocation")
	ErrorInvalidNodeStatus         = sdkerrors.Register(ModuleName, 204, "invalid node status")
	ErrorInvalidPlanStatus         = sdkerrors.Register(ModuleName, 205, "invalid plan status")
	ErrorInvalidSubscription       = sdkerrors.Register(ModuleName, 206, "invalid subscription")
	ErrorInvalidSubscriptionStatus = sdkerrors.Register(ModuleName, 207, "invalid subscription status")
	ErrorNodeNotFound              = sdkerrors.Register(ModuleName, 208, "node not found")
	ErrorPayoutNotFound            = sdkerrors.Register(ModuleName, 209, "payout not found")
	ErrorPlanNotFound              = sdkerrors.Register(ModuleName, 210, "plan not found")
	ErrorPriceNotFound             = sdkerrors.Register(ModuleName, 211, "price does not exist")
	ErrorSubscriptionNotFound      = sdkerrors.Register(ModuleName, 212, "subscription not found")
	ErrorUnauthorized              = sdkerrors.Register(ModuleName, 213, "unauthorised")
)

func NewErrorAllocationNotFound(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorAllocationNotFound, "allocation %d/%s does not exist", id, addr)
}

func NewErrorInsufficientBytes(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInsufficientBytes, "insufficient bytes for allocation %d/%s", id, addr)
}

func NewErrorInvalidAllocation(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

func NewErrorInvalidNodeStatus(addr interface{}, status hubtypes.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

func NewErrorInvalidPlanStatus(id uint64, status hubtypes.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidPlanStatus, "invalid status %s for plan %d", status, id)
}

func NewErrorInvalidSubscription(id uint64) error {
	return sdkerrors.Wrapf(ErrorInvalidSubscription, "invalid subscription %d", id)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status hubtypes.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidSubscriptionStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorNodeNotFound(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPayoutNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorPayoutNotFound, "payout %d does not exist", id)
}

func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorPriceNotFound, "price for denom %s does not exist", denom)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorSubscriptionNotFound, "subscription %d does not exist", id)
}

func NewErrorUnauthorized(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
