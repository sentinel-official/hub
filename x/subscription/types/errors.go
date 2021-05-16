package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField              = errors.Register(ModuleName, 101, "invalid field")
	ErrorPlanDoesNotExist          = errors.Register(ModuleName, 102, "plan does not exist")
	ErrorNodeDoesNotExist          = errors.Register(ModuleName, 103, "node does not exist")
	ErrorUnauthorized              = errors.Register(ModuleName, 104, "unauthorized")
	ErrorInvalidPlanStatus         = errors.Register(ModuleName, 105, "invalid plan status")
	ErrorPriceDoesNotExist         = errors.Register(ModuleName, 106, "price does not exist")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 107, "invalid node status")
	ErrorSubscriptionDoesNotExist  = errors.Register(ModuleName, 108, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 109, "invalid subscription status")
	ErrorCanNotSubscribe           = errors.Register(ModuleName, 110, "can not subscribe")
	ErrorInvalidQuota              = errors.Register(ModuleName, 111, "invalid quota")
	ErrorDuplicateQuota            = errors.Register(ModuleName, 112, "duplicate quota")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 113, "quota does not exist")
	ErrorCanNotAddQuota            = errors.Register(ModuleName, 114, "can not add quota")
)
