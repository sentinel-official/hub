package types

import (
	"fmt"
)

const (
	AttributeKeyAddress = "address"
	AttributeKeyDeposit = "deposit"
)

var (
	EventTypeSet    = fmt.Sprintf("%s:set", ModuleName)
	EventTypeUpdate = fmt.Sprintf("%s:update", ModuleName)
)
