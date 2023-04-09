package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 201, "provider does not exist")
	ErrorDuplicateNode        = errors.Register(ModuleName, 202, "duplicate node")
	ErrorNodeDoesNotExist     = errors.Register(ModuleName, 203, "node does not exist")
	ErrorInvalidPlanCount     = errors.Register(ModuleName, 204, "invalid plan count")
)
