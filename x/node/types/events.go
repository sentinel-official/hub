package types

import (
	"fmt"
)

const (
	AttributeKeyProvider = "provider"
	AttributeKeyAddress  = "address"
	AttributeKeyStatus   = "status"
)

var (
	EventTypeSet       = fmt.Sprintf("%s:set", ModuleName)
	EventTypeUpdate    = fmt.Sprintf("%s:update", ModuleName)
	EventTypeSetStatus = fmt.Sprintf("%s:set_status", ModuleName)
)
