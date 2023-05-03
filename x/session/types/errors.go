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
	ErrorDuplicateSession        = errors.Register(ModuleName, 201, "duplicate session")
	ErrorInvalidQuota            = errors.Register(ModuleName, 202, "invalid quota")
	ErrorInvalidSignature        = errors.Register(ModuleName, 203, "invalid signature")
	ErrorInvalidStatus           = errors.Register(ModuleName, 204, "invalid status")
	ErrorNodeNotFound            = errors.Register(ModuleName, 205, "node does not exist")
	ErrorQuotaNotFound           = errors.Register(ModuleName, 206, "quota not found")
	ErrorSessionNotFound         = errors.Register(ModuleName, 207, "session not found")
	ErrorSubscriptionNotFound    = errors.Register(ModuleName, 208, "subscription not found")
	ErrorUnauthorized            = errors.Register(ModuleName, 209, "unauthorized")
	ErrorUnexpectedNode          = errors.Register(ModuleName, 210, "unexpected node")
	ErrorInvalidSubscriptionType = errors.Register(ModuleName, 211, "invalid subscription type")
)

func NewErrorSubscriptionNotFound(id uint64) error {
	return errors.Wrapf(ErrorSubscriptionNotFound, "subscription %d does not exit", id)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorNodeNotFound(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorInvalidNodeStatus(addr hubtypes.NodeAddress, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for node %s", status, addr)
}

func NewErrorQuotaNotFound(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorQuotaNotFound, "subscription quota %d/%s does not exist", id, addr)
}

func NewErrorDuplicateSession(subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) error {
	return errors.Wrapf(ErrorDuplicateSession, "session %d already exist for subscription quota %d/%s", sessionID, subscriptionID, addr)
}

func NewErrorInvalidQuota(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorInvalidQuota, "invalid subscription quota %d/%s", id, addr)
}

func NewErrorSessionNotFound(id uint64) error {
	return errors.Wrapf(ErrorSessionNotFound, "session %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}

func NewErrorInvalidSignature(signature []byte) error {
	return errors.Wrapf(ErrorInvalidSignature, "invalid bandwidth proof signature %X", signature)
}

func NewErrorInvalidSessionStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for session %d", status, id)
}

func NewErrorUnexpectedNode(addr string) error {
	return errors.Wrapf(ErrorUnexpectedNode, "node %s is not expected", addr)
}

func NewErrorInvalidSubscriptionType(id uint64, t string) error {
	return errors.Wrapf(ErrorInvalidSubscriptionType, "invalid type %s for subscription %d", t, id)
}
