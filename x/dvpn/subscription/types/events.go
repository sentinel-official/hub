package types

const (
	EventTypeSetPlan                   = "set_plan"
	EventTypeSetPlansCount             = "set_plans_count"
	EventTypeSetPlanStatus             = "set_plan_status"
	EventTypeSetNodeAddressForPlan     = "set_node_address_for_plan"
	EventTypeDeleteNodeAddressForPlan  = "delete_node_address_for_plan"
	EventTypeSetSubscription           = "set_subscription"
	EventTypeSetSubscriptionsCount     = "set_subscriptions_count"
	EventTypeSetAddressForSubscription = "set_address_for_subscription"
)

const (
	AttributeKeyAddress = "address"
	AttributeKeyID      = "id"
	AttributeKeyNode    = "node"
	AttributeKeyCount   = "count"
	AttributeKeyStatus  = "status"
	AttributeKeyPlan    = "plan"
)
