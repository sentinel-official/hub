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
	ErrorSubscriptionDoesNotExit   = errors.Register(ModuleName, 106, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 107, "invalid subscription status")
	ErrorUnauthorized              = errors.Register(ModuleName, 108, "unauthorized")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 109, "quota does not exist")
	ErrorFailedToVerifyProof       = errors.Register(ModuleName, 110, "failed to verify proof")
	ErrorInvalidBandwidth          = errors.Register(ModuleName, 111, "invalid bandwidth")
)
