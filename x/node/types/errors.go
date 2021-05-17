package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidFieldFrom            = errors.Register(ModuleName, 101, "invalid value for field from; expected a valid address")
	ErrorProviderDoesNotExist        = errors.Register(ModuleName, 102, "provider does not exist")
	ErrorDuplicateNode               = errors.Register(ModuleName, 103, "duplicate node")
	ErrorNodeDoesNotExist            = errors.Register(ModuleName, 104, "node does not exist")
	ErrorInvalidFieldProviderOrPrice = errors.Register(ModuleName, 105, "invalid value for provider or price; expected either price or provider to be empty")
	ErrorInvalidFieldProvider        = errors.Register(ModuleName, 106, "invalid value for field provider; expected a valid provider address")
	ErrorInvalidFieldPrice           = errors.Register(ModuleName, 107, "invalid value for field price; expected a valid price")
	ErrorInvalidFieldRemoteURL       = errors.Register(ModuleName, 108, "invalid value for field remote_url; expected length between 1 and 64")
	ErrorInvalidFieldStatus          = errors.Register(ModuleName, 109, "invalid value for field status; expected either Active or Inactive")
	ErrorInvalidPlanCount            = errors.Register(ModuleName, 110, "invalid plan count")
)
