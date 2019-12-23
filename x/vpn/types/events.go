package types

var (
	EventTypeMsgRegisterNode   = "msg_register_node"
	EventTypeMsgUpdateNodeInfo = "msg_update_node_info"
	EventTypeMsgDeregisterNode = "msg_deregister_node"
	
	EventTypeMsgAddFreeClient    = "msg_add_free_client"
	EventTypeMsgRemoveFreeClient = "msg_remove_free_client"
	
	EventTypeMsgRegisterVPNOnResolver   = "msg_register_vpn_on_resolver"
	EventTypeMsgDeregisterVPNOnResolver = "msg_deregister_vpn_on_resolver"
	
	EventTypeMsgStartSubscription = "msg_start_subscription"
	EventTypeMsgEndSubscription   = "msg_end_subscription"
	
	EventTypeMsgUpdateSessionInfo = "msg_update_session_info"
	
	EventTypeMsgRegisterResolver   = "msg_register_resolver"
	EventTypeMsgUpdateResolverInfo = "msg_update_resolver_info"
	EventTypeMsgDeregisterResolver = "msg_deregister_resolver"
	
	AttributeKeyClientAddress = "client_address"
	AttributeKeyFromAddress   = "from_address"
	AttributeKeyNodeID        = "node_id"
	AttributeSubscriptionID   = "subscription_id"
	AttributeSessionID        = "session_id"
	AttributeKeyResolverID    = "resolver_id"
	AttributeKeyStatus        = "status"
	AttributeKeyCommission    = "commission"
	AttributeKeyDeposit       = "deposit"
)
