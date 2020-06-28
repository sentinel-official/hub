package types

const (
	EventTypeSetPlan                  = "set_plan"
	EventTypeSetPlansCount            = "set_plans_count"
	EventTypeSetPlanStatus            = "set_plan_status"
	EventTypeSetNodeAddressForPlan    = "set_node_address_for_plan"
	EventTypeRemoveNodeAddressForPlan = "remove_node_address_for_plan"
)

const (
	AttributeKeyAddress = "address"
	AttributeKeyID      = "id"
	AttributeKeyCount   = "count"
	AttributeKeyStatus  = "status"
)
