package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField          = errors.Register(ModuleName, 101, "invalid field")
	ErrorInvalidFrom           = errors.Register(ModuleName, 102, "invalid from")
	ErrorInvalidId             = errors.Register(ModuleName, 103, "invalid id")
	ErrorInvalidNode           = errors.Register(ModuleName, 104, "invalid node")
	ErrorInvalidProofId        = errors.Register(ModuleName, 105, "invalid proof->id")
	ErrorInvalidProofDuration  = errors.Register(ModuleName, 106, "invalid proof->duration")
	ErrorInvalidProofBandwidth = errors.Register(ModuleName, 107, "invalid proof->bandwidth")
	ErrorInvalidSignature      = errors.Register(ModuleName, 108, "invalid signature")
	ErrorInvalidRating         = errors.Register(ModuleName, 109, "invalid rating")
)

var (
	ErrorSubscriptionDoesNotExit   = errors.Register(ModuleName, 201, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 202, "invalid subscription status")
	ErrorUnauthorized              = errors.Register(ModuleName, 203, "unauthorized")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 204, "quota does not exist")
	ErrorFailedToVerifyProof       = errors.Register(ModuleName, 205, "failed to verify proof")
	ErrorNotEnoughQuota            = errors.Register(ModuleName, 206, "not enough quota")
	ErrorNodeDoesNotExist          = errors.Register(ModuleName, 207, "node does not exist")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 208, "invalid node status")
	ErrorSessionDoesNotExist       = errors.Register(ModuleName, 209, "session does not exist")
	ErrorInvalidSessionStatus      = errors.Register(ModuleName, 210, "invalid session status")
	ErrorNodeAddressMismatch       = errors.Register(ModuleName, 211, "node address mismatch")
	ErrorNodeDoesNotExistForPlan   = errors.Register(ModuleName, 212, "node does not exist for plan")
	ErrorDuplicateSession          = errors.Register(ModuleName, 213, "duplicate session")
)
