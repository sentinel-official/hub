package types

const (
	EventTypeSetPlan                  = "set_plan"
	EventTypeSetPlansCount            = "set_plans_count"
	EventTypeSetPlanStatus            = "set_plan_status"
	EventTypeSetNodeAddressForPlan    = "set_node_address_for_plan"
	EventTypeDeleteNodeAddressForPlan = "delete_node_address_for_plan"
	EventTypeSetSubscription          = "set_subscription"
	EventTypeSetSubscriptionsCount    = "set_subscriptions_count"
)

const (
	AttributeKeyAddress = "address"
	AttributeKeyID      = "id"
	AttributeKeyCount   = "count"
	AttributeKeyStatus  = "status"
	AttributeKeyPlan    = "plan"
)
