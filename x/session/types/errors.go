// DO NOT COVER

package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")

	ErrorAllocationNotFound             = errors.Register(ModuleName, 201, "allocation not found")
	ErrorDuplicateActiveSession         = errors.Register(ModuleName, 202, "duplicate active session")
	ErrorInvalidAllocation              = errors.Register(ModuleName, 203, "invalid allocation")
	ErrorInvalidNode                    = errors.Register(ModuleName, 204, "invalid node")
	ErrorInvalidNodeStatus              = errors.Register(ModuleName, 205, "invalid node status")
	ErrorInvalidSessionStatus           = errors.Register(ModuleName, 206, "invalid session status")
	ErrorInvalidSignature               = errors.Register(ModuleName, 207, "invalid signature")
	ErrorInvalidSubscription            = errors.Register(ModuleName, 208, "invalid subscription")
	ErrorInvalidSubscriptionStatus      = errors.Register(ModuleName, 209, "invalid subscription status")
	ErrorNodeNotFound                   = errors.Register(ModuleName, 210, "node not found")
	ErrorPayoutForAddressByNodeNotFound = errors.Register(ModuleName, 211, "payout for address by node not found")
	ErrorPlanNotFound                   = errors.Register(ModuleName, 212, "plan not found")
	ErrorSessionNotFound                = errors.Register(ModuleName, 213, "session not found")
	ErrorSubscriptionNotFound           = errors.Register(ModuleName, 214, "subscription not found")
	ErrorUnauthorized                   = errors.Register(ModuleName, 215, "unauthorised")
)

func NewErrorAllocationNotFound(id uint64, addr interface{}) error {
	return errors.Wrapf(ErrorAllocationNotFound, "allocation %d/%s does not exist", id, addr)
}

func NewErrorDuplicateActiveSession(id uint64) error {
	return errors.Wrapf(ErrorDuplicateActiveSession, "active session %d already exist", id)
}

func NewErrorInvalidAllocation(id uint64, addr interface{}) error {
	return errors.Wrapf(ErrorInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

func NewErrorInvalidNode(addr interface{}) error {
	return errors.Wrapf(ErrorInvalidNode, "invalid node %s ", addr)
}

func NewErrorInvalidNodeStatus(addr interface{}, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %s ", status, addr)
}

func NewErrorInvalidSessionStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidSessionStatus, "invalid status %s for session %d", status, id)
}

func NewErrorInvalidSignature(signature []byte) error {
	return errors.Wrapf(ErrorInvalidSignature, "invalid signature %X", signature)
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

func NewErrorPayoutForAddressByNodeNotFound(accAddr, nodeAddr interface{}) error {
	return errors.Wrapf(ErrorPayoutForAddressByNodeNotFound, "payout for address %s by node %s not found", accAddr, nodeAddr)
}

func NewErrorPlanNotFound(id uint64) error {
	return errors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorSessionNotFound(id uint64) error {
	return errors.Wrapf(ErrorSessionNotFound, "session %d does not exist", id)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return errors.Wrapf(ErrorSubscriptionNotFound, "subscription %s does not exist", id)
}

func NewErrorUnauthorized(addr interface{}) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
