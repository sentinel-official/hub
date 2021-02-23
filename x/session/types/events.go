package types

import (
	"fmt"
)

const (
	AttributeKeyCount        = "count"
	AttributeKeyID           = "id"
	AttributeKeySubscription = "subscription"
	AttributeKeyAddress      = "address"
)

var (
	EventTypeSetCount  = fmt.Sprintf("%s:set_count", ModuleName)
	EventTypeSetActive = fmt.Sprintf("%s:set_active", ModuleName)
	EventTypeUpdate    = fmt.Sprintf("%s:update", ModuleName)
)
