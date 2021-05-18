package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField         = errors.Register(ModuleName, 101, "invalid field")
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 102, "provider does not exist")
	ErrorPlanDoesNotExist     = errors.Register(ModuleName, 103, "plan does not exist")
	ErrorNodeDoesNotExist     = errors.Register(ModuleName, 104, "node does not exist")
	ErrorUnauthorized         = errors.Register(ModuleName, 105, "unauthorized")
	DuplicateNodeForPlan      = errors.Register(ModuleName, 106, "duplicate node for plan")
)
