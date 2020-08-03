package types

const (
	EventTypeSetSubscription            = "set_subscription"
	EventTypeSetSubscriptionsCount      = "set_subscriptions_count"
	EventTypeAddQuotaForSubscription    = "add_quota_for_subscription"
	EventTypeUpdateQuotaForSubscription = "update_quota_for_subscription"
	EventTypeRemoveQuotaForSubscription = "remove_quota_for_subscription"
	EventTypeEndSubscription            = "end_subscription"
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
