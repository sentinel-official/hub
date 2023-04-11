package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorInvalidGigabytePrices = errors.Register(ModuleName, 201, "invalid gigabyte prices")
	ErrorInvalidHourlyPrices   = errors.Register(ModuleName, 202, "invalid hourly prices")
	ErrorDuplicateNode         = errors.Register(ModuleName, 203, "duplicate node")
	ErrorNodeDoesNotExist      = errors.Register(ModuleName, 204, "node does not exist")
)
