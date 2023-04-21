package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
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
