package types

const (
	EventTypeSetCount    = "subscription:set_count"
	EventTypeSet         = "subscription:set"
	EventTypeCancel      = "subscription:cancel"
	EventTypeAddQuota    = "subscription:add_quota"
	EventTypeUpdateQuota = "subscription:update_quota"
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
