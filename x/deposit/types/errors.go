package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField = errors.Register(ModuleName, 101, "invalid field")
)

var (
	ErrorInsufficientDepositFunds = errors.Register(ModuleName, 201, "insufficient deposit funds")
	ErrorDepositDoesNotExist      = errors.Register(ModuleName, 202, "deposit does not exist")
)
