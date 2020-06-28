package types

const (
	EventTypeSetSubscription              = "set_subscription"
	EventTypeSetSubscriptionsCount        = "set_subscriptions_count"
	EventTypeSetAddressForSubscription    = "set_address_for_subscription"
	EventTypeRemoveAddressForSubscription = "remove_address_for_subscription"
	EventTypeEndSubscription              = "end_subscription"
)

const (
	AttributeKeyAddress = "address"
	AttributeKeyID      = "id"
	AttributeKeyNode    = "node"
	AttributeKeyCount   = "count"
	AttributeKeyPlan    = "plan"
)
