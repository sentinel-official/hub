package types

import (
	"fmt"
)

const (
	AttributeKeyAddress = "address"
	AttributeKeyID      = "id"
	AttributeKeyCount   = "count"
	AttributeKeyStatus  = "status"
)

var (
	EventTypeSetCount   = fmt.Sprintf("%s:set_count", ModuleName)
	EventTypeSet        = fmt.Sprintf("%s:set", ModuleName)
	EventTypeSetStatus  = fmt.Sprintf("%s:set_status", ModuleName)
	EventTypeAddNode    = fmt.Sprintf("%s:add_node", ModuleName)
	EventTypeRemoveNode = fmt.Sprintf("%s:remove_node", ModuleName)
)
