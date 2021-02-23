package types

import (
	"fmt"
)

const (
	AttributeKeyOwner     = "owner"
	AttributeKeyAddress   = "address"
	AttributeKeyID        = "id"
	AttributeKeyStatus    = "status"
	AttributeKeyNode      = "node"
	AttributeKeyCount     = "count"
	AttributeKeyPlan      = "plan"
	AttributeKeyConsumed  = "consumed"
	AttributeKeyAllocated = "allocated"
)

var (
	EventTypeSetCount    = fmt.Sprintf("%s:set_count", ModuleName)
	EventTypeSet         = fmt.Sprintf("%s:set", ModuleName)
	EventTypeCancel      = fmt.Sprintf("%s:cancel", ModuleName)
	EventTypeAddQuota    = fmt.Sprintf("%s:add_quota", ModuleName)
	EventTypeUpdateQuota = fmt.Sprintf("%s:update_quota", ModuleName)
)
