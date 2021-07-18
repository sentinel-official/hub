package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField    = errors.Register(ModuleName, 101, "invalid field")
	ErrorInvalidFrom     = errors.Register(ModuleName, 102, "invalid from")
	ErrorInvalidReceiver = errors.Register(ModuleName, 103, "invalid receiver")
	ErrorInvalidTxHash   = errors.Register(ModuleName, 104, "invalid tx_hash")
	ErrorInvalidAmount   = errors.Register(ModuleName, 105, "invalid amount")
)

var (
	ErrorSwapIsDisabled = errors.Register(ModuleName, 201, "swap is disabled")
	ErrorUnauthorized   = errors.Register(ModuleName, 202, "unauthorized")
	ErrorDuplicateSwap  = errors.Register(ModuleName, 203, "duplicate swap")
)

var (
	ErrorUnknownMsgType = errors.Register(ModuleName, 301, "unknown message type")
)
