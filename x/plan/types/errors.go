package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 201, "provider does not exist")
	ErrorPlanDoesNotExist     = errors.Register(ModuleName, 202, "plan does not exist")
	ErrorNodeDoesNotExist     = errors.Register(ModuleName, 203, "node does not exist")
	ErrorUnauthorized         = errors.Register(ModuleName, 204, "unauthorized")
	DuplicateNodeForPlan      = errors.Register(ModuleName, 205, "duplicate node for plan")
)
