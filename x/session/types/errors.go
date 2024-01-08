// DO NOT COVER

package types

import (
	sdkerrors "cosmossdk.io/errors"

	hubtypes "github.com/sentinel-official/hub/v12/types"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorAllocationNotFound             = sdkerrors.Register(ModuleName, 201, "allocation not found")
	ErrorDuplicateActiveSession         = sdkerrors.Register(ModuleName, 202, "duplicate active session")
	ErrorInvalidAllocation              = sdkerrors.Register(ModuleName, 203, "invalid allocation")
	ErrorInvalidNode                    = sdkerrors.Register(ModuleName, 204, "invalid node")
	ErrorInvalidNodeStatus              = sdkerrors.Register(ModuleName, 205, "invalid node status")
	ErrorInvalidSessionStatus           = sdkerrors.Register(ModuleName, 206, "invalid session status")
	ErrorInvalidSignature               = sdkerrors.Register(ModuleName, 207, "invalid signature")
	ErrorInvalidSubscription            = sdkerrors.Register(ModuleName, 208, "invalid subscription")
	ErrorInvalidSubscriptionStatus      = sdkerrors.Register(ModuleName, 209, "invalid subscription status")
	ErrorNodeNotFound                   = sdkerrors.Register(ModuleName, 210, "node not found")
	ErrorPayoutForAddressByNodeNotFound = sdkerrors.Register(ModuleName, 211, "payout for address by node not found")
	ErrorPlanNotFound                   = sdkerrors.Register(ModuleName, 212, "plan not found")
	ErrorSessionNotFound                = sdkerrors.Register(ModuleName, 213, "session not found")
	ErrorSubscriptionNotFound           = sdkerrors.Register(ModuleName, 214, "subscription not found")
	ErrorUnauthorized                   = sdkerrors.Register(ModuleName, 215, "unauthorised")
)

func NewErrorAllocationNotFound(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorAllocationNotFound, "allocation %d/%s does not exist", id, addr)
}

func NewErrorDuplicateActiveSession(id uint64) error {
	return sdkerrors.Wrapf(ErrorDuplicateActiveSession, "active session %d already exist", id)
}

func NewErrorInvalidAllocation(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

func NewErrorInvalidNode(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInvalidNode, "invalid node %s", addr)
}

func NewErrorInvalidNodeStatus(addr interface{}, status hubtypes.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

func NewErrorInvalidSessionStatus(id uint64, status hubtypes.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidSessionStatus, "invalid status %s for session %d", status, id)
}

func NewErrorInvalidSignature(signature []byte) error {
	return sdkerrors.Wrapf(ErrorInvalidSignature, "invalid signature %X", signature)
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

func NewErrorPayoutForAddressByNodeNotFound(accAddr, nodeAddr interface{}) error {
	return sdkerrors.Wrapf(ErrorPayoutForAddressByNodeNotFound, "payout for address %s by node %s not found", accAddr, nodeAddr)
}

func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorSessionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorSessionNotFound, "session %d does not exist", id)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorSubscriptionNotFound, "subscription %d does not exist", id)
}

func NewErrorUnauthorized(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
