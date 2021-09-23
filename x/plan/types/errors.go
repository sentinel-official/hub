package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField    = errors.Register(ModuleName, 101, "invalid field")
	ErrorInvalidFrom     = errors.Register(ModuleName, 102, "invalid from")
	ErrorInvalidPrice    = errors.Register(ModuleName, 103, "invalid price")
	ErrorInvalidValidity = errors.Register(ModuleName, 104, "invalid validity")
	ErrorInvalidBytes    = errors.Register(ModuleName, 105, "invalid bytes")
	ErrorInvalidId       = errors.Register(ModuleName, 106, "invalid id")
	ErrorInvalidStatus   = errors.Register(ModuleName, 107, "invalid status")
	ErrorInvalidAddress  = errors.Register(ModuleName, 108, "invalid address")
)

var (
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 201, "provider does not exist")
	ErrorPlanDoesNotExist     = errors.Register(ModuleName, 202, "plan does not exist")
	ErrorNodeDoesNotExist     = errors.Register(ModuleName, 203, "node does not exist")
	ErrorUnauthorized         = errors.Register(ModuleName, 204, "unauthorized")
	DuplicateNodeForPlan      = errors.Register(ModuleName, 205, "duplicate node for plan")
)
