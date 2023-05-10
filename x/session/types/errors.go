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
	ErrorDuplicate        = errors.Register(ModuleName, 201, "duplicate")
	ErrorInvalidQuota     = errors.Register(ModuleName, 202, "invalid quota")
	ErrorInvalidSignature = errors.Register(ModuleName, 203, "invalid signature")
	ErrorInvalidStatus    = errors.Register(ModuleName, 204, "invalid status")
	ErrorNotFound         = errors.Register(ModuleName, 205, "not found")
	ErrorUnauthorized     = errors.Register(ModuleName, 209, "unauthorized")
	ErrorInvalidNode      = errors.Register(ModuleName, 210, "invalid node")
	ErrorInvalidType      = errors.Register(ModuleName, 211, "invalid type")
)

func NewErrorSubscriptionNotFound(id uint64) error {
	return errors.Wrapf(ErrorNotFound, "subscription %d does not exit", id)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorNodeNotFound(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorNotFound, "node %s does not exist", addr)
}

func NewErrorInvalidNodeStatus(addr hubtypes.NodeAddress, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for node %s", status, addr)
}

func NewErrorQuotaNotFound(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorNotFound, "subscription quota %d/%s does not exist", id, addr)
}

func NewErrorDuplicateSession(subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) error {
	return errors.Wrapf(ErrorDuplicate, "session %d already exist for subscription quota %d/%s", sessionID, subscriptionID, addr)
}

func NewErrorInvalidQuota(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorInvalidQuota, "invalid subscription quota %d/%s", id, addr)
}

func NewErrorSessionNotFound(id uint64) error {
	return errors.Wrapf(ErrorNotFound, "session %d does not exist", id)
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

func NewErrorInvalidNode(addr string) error {
	return errors.Wrapf(ErrorInvalidNode, "invalid node %s", addr)
}

func NewErrorInvalidSubscriptionType(id uint64, t string) error {
	return errors.Wrapf(ErrorInvalidType, "invalid type %s for subscription %d", t, id)
}
