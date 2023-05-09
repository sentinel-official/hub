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
	ErrorInvalidQuota      = errors.Register(ModuleName, 202, "invalid quota")
	ErrorInvalidStatus     = errors.Register(ModuleName, 203, "invalid status")
	ErrorInvalidType       = errors.Register(ModuleName, 204, "invalid type")
	ErrorNotFound          = errors.Register(ModuleName, 205, "not found")
	ErrorUnauthorized      = errors.Register(ModuleName, 207, "unauthorized")
)

func NewErrorInsufficientBytes(id uint64, bytes sdk.Int) error {
	return errors.Wrapf(ErrorInsufficientBytes, "insufficient bytes %d for subscription %d", bytes, id)
}

func NewErrorInvalidQuota(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorInvalidQuota, "invalid quota %d/%s", id, addr)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status hubtypes.Status) error {
	return errors.Wrapf(ErrorInvalidStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorInvalidSubscriptionType(id uint64, t SubscriptionType) error {
	return errors.Wrapf(ErrorInvalidType, "invalid type %s for subscription %d", t, id)
}

func NewErrorQuotaNotFound(id uint64, addr sdk.AccAddress) error {
	return errors.Wrapf(ErrorNotFound, "subscription quota %d/%s does not exist", id, addr)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return errors.Wrapf(ErrorNotFound, "subscription %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return errors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
