package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorMarshal                     = errors.Register(ModuleName, 101, "error occurred while marshalling")
	ErrorUnmarshal                   = errors.Register(ModuleName, 102, "error occurred while unmarshalling")
	ErrorUnknownMsgType              = errors.Register(ModuleName, 103, "unknown message type")
	ErrorUnknownQueryType            = errors.Register(ModuleName, 104, "unknown query type")
	ErrorInvalidFieldFrom            = errors.Register(ModuleName, 105, "invalid value for field from; expected a valid address")
	ErrorProviderDoesNotExist        = errors.Register(ModuleName, 106, "provider does not exist")
	ErrorDuplicateNode               = errors.Register(ModuleName, 107, "duplicate node")
	ErrorNodeDoesNotExist            = errors.Register(ModuleName, 108, "node does not exist")
	ErrorInvalidFieldProviderOrPrice = errors.Register(ModuleName, 109, "invalid value for provider or price; expcted either price or provider to be empty")
	ErrorInvalidFieldProvider        = errors.Register(ModuleName, 110, "invalid value for field provider; expected a valid provider address")
	ErrorInvalidFieldPrice           = errors.Register(ModuleName, 111, "invalid value for field price; expected a valid price")
	ErrorInvalidFieldRemoteURL       = errors.Register(ModuleName, 112, "invalid value for field remote_url; expcted length between 1 and 64")
	ErrorInvalidFieldStatus          = errors.Register(ModuleName, 113, "invalid value for field status; expcted either Active or Inactive")
)
