package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField     = errors.Register(ModuleName, 101, "invalid field")
	ErrorInvalidFrom      = errors.Register(ModuleName, 102, "invalid from")
	ErrorInvalidProvider  = errors.Register(ModuleName, 103, "invalid provider")
	ErrorInvalidPrice     = errors.Register(ModuleName, 104, "invalid price")
	ErrorInvalidRemoteURL = errors.Register(ModuleName, 105, "invalid remote_url")
	ErrorInvalidStatus    = errors.Register(ModuleName, 106, "invalid status")
)

var (
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 201, "provider does not exist")
	ErrorDuplicateNode        = errors.Register(ModuleName, 202, "duplicate node")
	ErrorNodeDoesNotExist     = errors.Register(ModuleName, 203, "node does not exist")
	ErrorInvalidPlanCount     = errors.Register(ModuleName, 204, "invalid plan count")
)
