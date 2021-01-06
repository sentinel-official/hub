package types

import (
	"fmt"
)

const (
	AttributeKeyAddress = "address"
)

var (
	EventTypeSet    = fmt.Sprintf("%s:set", ModuleName)
	EventTypeUpdate = fmt.Sprintf("%s:update", ModuleName)
)
