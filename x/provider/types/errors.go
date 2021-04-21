package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorMarshal                 = errors.Register(ModuleName, 101, "error occurred while marshalling")
	ErrorUnmarshal               = errors.Register(ModuleName, 102, "error occurred while unmarshalling")
	ErrorUnknownMsgType          = errors.Register(ModuleName, 103, "unknown message type")
	ErrorUnknownQueryType        = errors.Register(ModuleName, 104, "unknown query type")
	ErrorInvalidFieldName        = errors.Register(ModuleName, 105, "invalid field name; expected length between 1 and 64")
	ErrorDuplicateProvider       = errors.Register(ModuleName, 106, "duplicate provider")
	ErrorProviderDoesNotExist    = errors.Register(ModuleName, 107, "provider does not exist")
	ErrorInvalidFieldFrom        = errors.Register(ModuleName, 108, "invalid field from; expected a valid address")
	ErrorInvalidFieldIdentity    = errors.Register(ModuleName, 109, "invalid field identity; expected length between 0 and 64")
	ErrorInvalidFieldWebsite     = errors.Register(ModuleName, 110, "invalid field website; expected length between 0 and 64")
	ErrorInvalidFieldDescription = errors.Register(ModuleName, 111, "invalid field description; expected length between 0 and 256")
)
