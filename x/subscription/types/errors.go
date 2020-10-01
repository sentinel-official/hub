package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorMarshal                   = errors.Register(ModuleName, 101, "error occurred while marshalling")
	ErrorUnmarshal                 = errors.Register(ModuleName, 102, "error occurred while unmarshalling")
	ErrorUnknownMsgType            = errors.Register(ModuleName, 103, "unknown message type")
	ErrorUnknownQueryType          = errors.Register(ModuleName, 104, "unknown query type")
	ErrorInvalidField              = errors.Register(ModuleName, 105, "invalid field")
	ErrorPlanDoesNotExist          = errors.Register(ModuleName, 106, "plan does not exist")
	ErrorNodeDoesNotExist          = errors.Register(ModuleName, 107, "node does not exist")
	ErrorUnauthorized              = errors.Register(ModuleName, 108, "unauthorized")
	ErrorInvalidPlanStatus         = errors.Register(ModuleName, 109, "invalid plan status")
	ErrorPriceDoesNotExist         = errors.Register(ModuleName, 110, "price does not exist")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 111, "invalid node status")
	ErrorSubscriptionDoesNotExist  = errors.Register(ModuleName, 112, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 113, "invalid subscription status")
	ErrorCanNotSubscribe           = errors.Register(ModuleName, 114, "can not subscribe")
	ErrorInvalidQuota              = errors.Register(ModuleName, 115, "invalid quota")
	ErrorDuplicateQuota            = errors.Register(ModuleName, 116, "duplicate quota")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 117, "quota does not exist")
	ErrorCanNotAddQuota            = errors.Register(ModuleName, 118, "can not add quota")
)
