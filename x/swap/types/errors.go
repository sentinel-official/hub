package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorMarshal              = errors.Register(ModuleName, 101, "error occurred while marshalling")
	ErrorUnmarshal            = errors.Register(ModuleName, 102, "error occurred while unmarshalling")
	ErrorUnknownMsgType       = errors.Register(ModuleName, 103, "unknown message type")
	ErrorUnknownQueryType     = errors.Register(ModuleName, 104, "unknown query type")
	ErrorInvalidFieldFrom     = errors.Register(ModuleName, 105, "invalid value for field from; expected a valid address")
	ErrorSwapIsDisabled       = errors.Register(ModuleName, 106, "swap is disabled")
	ErrorUnauthorized         = errors.Register(ModuleName, 107, "unauthorized")
	ErrorDuplicateSwap        = errors.Register(ModuleName, 108, "duplicate swap")
	ErrorInvalidFieldReceiver = errors.Register(ModuleName, 109, "invalid value for field receiver; expected a valid address")
	ErrorInvalidFieldAmount   = errors.Register(ModuleName, 110, "invalid value for field amount; expected a value greater than or equal to 100")
)
