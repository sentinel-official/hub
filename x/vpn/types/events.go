package types

var (
	EventTypeMsgRegisterNode        = "msg_register_node"
	EventTypeMsgUpdateInfo          = "msg_update_info"
	EventTypeMsgAddFreeClient       = "msg_add_free_client"
	EventTypeMsgRemoveFreeClient    = "msg_remove_free_client"
	EventTypeMsgAddVPNOnResolver    = "msg_add_vpn_on_resolver"
	EventTypeMsgRemoveVPNOnResolver = "msg_remove_vpn_on_resolver"
	EventTypeMsgStartSubscription   = "msg_start_subscription"
	EventTypeMsgEndSubscription     = "msg_end_subscription"
	EventTypeMsgUpdateSessionInfo   = "msg_update_session_info"
	EventTypeMsgRegisterResolver    = "msg_register_resolver"
	EventTypeMsgUpdateResolverInfo  = "msg_update_resolver_info"
	EventTypeMsgDeregisterResolver  = "msg_deregister_resolver"
	
	AttributeKeyAddress    = "address"
	AttributeKeyID         = "id"
	AttributeKeyStatus     = "status"
	AttributeKeyCommission = "commission"
)
