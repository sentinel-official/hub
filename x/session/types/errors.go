package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField              = errors.Register(ModuleName, 101, "invalid field")
	ErrorSubscriptionDoesNotExit   = errors.Register(ModuleName, 102, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 103, "invalid subscription status")
	ErrorUnauthorized              = errors.Register(ModuleName, 104, "unauthorized")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 105, "quota does not exist")
	ErrorFailedToVerifyProof       = errors.Register(ModuleName, 106, "failed to verify proof")
	ErrorNotEnoughQuota            = errors.Register(ModuleName, 107, "not enough quota")
	ErrorNodeDoesNotExist          = errors.Register(ModuleName, 108, "node does not exist")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 109, "invalid node status")
	ErrorSessionDoesNotExist       = errors.Register(ModuleName, 110, "session does not exist")
	ErrorInvalidSessionStatus      = errors.Register(ModuleName, 111, "invalid session status")
	ErrorAccountDoesNotExist       = errors.Register(ModuleName, 112, "account does not exist")
	ErrorPublicKeyDoesNotExist     = errors.Register(ModuleName, 113, "public key does not exist")
	ErrorInvalidSignature          = errors.Register(ModuleName, 114, "invalid signature")
)
