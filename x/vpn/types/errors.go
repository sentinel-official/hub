package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorUnknownMsgType = errors.Register(ModuleName, 301, "unknown message type")
)
