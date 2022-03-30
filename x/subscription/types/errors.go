package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField   = errors.Register(ModuleName, 101, "invalid field")
	ErrorInvalidFrom    = errors.Register(ModuleName, 102, "invalid from")
	ErrorInvalidAddress = errors.Register(ModuleName, 103, "invalid address")
	ErrorInvalidDeposit = errors.Register(ModuleName, 104, "invalid deposit")
	ErrorInvalidId      = errors.Register(ModuleName, 105, "invalid id")
	ErrorInvalidDenom   = errors.Register(ModuleName, 106, "invalid denom")
	ErrorInvalidBytes   = errors.Register(ModuleName, 107, "invalid bytes")
)

var (
	ErrorPlanDoesNotExist          = errors.Register(ModuleName, 201, "plan does not exist")
	ErrorNodeDoesNotExist          = errors.Register(ModuleName, 202, "node does not exist")
	ErrorUnauthorized              = errors.Register(ModuleName, 203, "unauthorized")
	ErrorInvalidPlanStatus         = errors.Register(ModuleName, 204, "invalid plan status")
	ErrorPriceDoesNotExist         = errors.Register(ModuleName, 205, "price does not exist")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 206, "invalid node status")
	ErrorSubscriptionDoesNotExist  = errors.Register(ModuleName, 207, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 208, "invalid subscription status")
	ErrorCanNotSubscribe           = errors.Register(ModuleName, 209, "can not subscribe")
	ErrorInvalidQuota              = errors.Register(ModuleName, 210, "invalid quota")
	ErrorDuplicateQuota            = errors.Register(ModuleName, 211, "duplicate quota")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 212, "quota does not exist")
	ErrorCanNotAddQuota            = errors.Register(ModuleName, 213, "can not add quota")
	ErrorCanNotUpdateQuota         = errors.Register(ModuleName, 214, "can not update quota")
	ErrorCanNotCancel              = errors.Register(ModuleName, 215, "can not cancel")
)
