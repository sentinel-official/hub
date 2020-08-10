package types

const (
	EventTypeSetCount    = "subscription:set_count"
	EventTypeSet         = "subscription:set"
	EventTypeEnd         = "subscription:end"
	EventTypeAddQuota    = "subscription:add_quota"
	EventTypeUpdateQuota = "subscription:update_quota"
)

const (
	AttributeKeyAddress   = "address"
	AttributeKeyID        = "id"
	AttributeKeyNode      = "node"
	AttributeKeyCount     = "count"
	AttributeKeyPlan      = "plan"
	AttributeKeyConsumed  = "consumed"
	AttributeKeyAllocated = "allocated"
)
