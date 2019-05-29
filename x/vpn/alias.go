package vpn

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/querier"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	Codespace                        = types.Codespace
	StoreKeyNode                     = types.StoreKeyNode
	StoreKeySession                  = types.StoreKeySession
	StoreKeySubscription             = types.StoreKeySubscription
	QuerierRoute                     = types.QuerierRoute
	RouterKey                        = types.RouterKey
	StatusRegistered                 = types.StatusRegistered
	StatusActive                     = types.StatusActive
	StatusInactive                   = types.StatusInactive
	StatusDeRegistered               = types.StatusDeRegistered
	TagNodeID                        = types.TagNodeID
	TagSessionID                     = types.TagSessionID
	DefaultParamspace                = keeper.DefaultParamspace
	QueryNode                        = querier.QueryNode
	QueryNodesOfAddress              = querier.QueryNodesOfAddress
	QueryAllNodes                    = querier.QueryAllNodes
	QuerySubscription                = querier.QuerySubscription
	QuerySubscriptionsOfNode         = querier.QuerySubscriptionsOfNode
	QuerySubscriptionsOfAddress      = querier.QuerySubscriptionsOfAddress
	QueryAllSubscriptions            = querier.QueryAllSubscriptions
	QuerySessionsCountOfSubscription = querier.QuerySessionsCountOfSubscription
	QuerySession                     = querier.QuerySession
	QuerySessionsOfSubscription      = querier.QuerySessionsOfSubscription
	QueryAllSessions                 = querier.QueryAllSessions
)

type (
	GenesisState                           = types.GenesisState
	Node                                   = types.Node
	MsgRegisterNode                        = types.MsgRegisterNode
	MsgUpdateNodeInfo                      = types.MsgUpdateNodeInfo
	MsgUpdateNodeStatus                    = types.MsgUpdateNodeStatus
	MsgDeregisterNode                      = types.MsgDeregisterNode
	Params                                 = types.Params
	Session                                = types.Session
	MsgUpdateSessionInfo                   = types.MsgUpdateSessionInfo
	Subscription                           = types.Subscription
	BandwidthSignatureData                 = types.BandwidthSignatureData
	MsgStartSubscription                   = types.MsgStartSubscription
	MsgEndSubscription                     = types.MsgEndSubscription
	Keeper                                 = keeper.Keeper
	QueryNodeParams                        = querier.QueryNodeParams
	QueryNodesOfAddressPrams               = querier.QueryNodesOfAddressPrams
	QuerySubscriptionParams                = querier.QuerySubscriptionParams
	QuerySubscriptionsOfNodePrams          = querier.QuerySubscriptionsOfNodePrams
	QuerySubscriptionsOfAddressParams      = querier.QuerySubscriptionsOfAddressParams
	QuerySessionsCountOfSubscriptionParams = querier.QuerySessionsCountOfSubscriptionParams
	QuerySessionParams                     = querier.QuerySessionParams
	QuerySessionsOfSubscriptionPrams       = querier.QuerySessionsOfSubscriptionPrams
)

// nolint: gochecknoglobals
var (
	RegisterCodec                             = types.RegisterCodec
	NewGenesisState                           = types.NewGenesisState
	DefaultGenesisState                       = types.DefaultGenesisState
	NodesCountKey                             = types.NodesCountKey
	NodeKeyPrefix                             = types.NodeKeyPrefix
	NodesCountOfAddressKeyPrefix              = types.NodesCountOfAddressKeyPrefix
	NodeIDByAddressKeyPrefix                  = types.NodeIDByAddressKeyPrefix
	SubscriptionsCountKey                     = types.SubscriptionsCountKey
	SubscriptionKeyPrefix                     = types.SubscriptionKeyPrefix
	SubscriptionsCountOfNodeKeyPrefix         = types.SubscriptionsCountOfNodeKeyPrefix
	SubscriptionIDByNodeIDKeyPrefix           = types.SubscriptionIDByNodeIDKeyPrefix
	SubscriptionsCountOfAddressKeyPrefix      = types.SubscriptionsCountOfAddressKeyPrefix
	SubscriptionIDByAddressKeyPrefix          = types.SubscriptionIDByAddressKeyPrefix
	SessionsCountKey                          = types.SessionsCountKey
	SessionsCountOfSubscriptionKeyPrefix      = types.SessionsCountOfSubscriptionKeyPrefix
	SessionKeyPrefix                          = types.SessionKeyPrefix
	SessionIDBySubscriptionIDKeyPrefix        = types.SessionIDBySubscriptionIDKeyPrefix
	NodeKey                                   = types.NodeKey
	NodesCountOfAddressKey                    = types.NodesCountOfAddressKey
	NodeIDByAddressKey                        = types.NodeIDByAddressKey
	SubscriptionKey                           = types.SubscriptionKey
	SubscriptionsCountOfNodeKey               = types.SubscriptionsCountOfNodeKey
	SubscriptionIDByNodeIDKey                 = types.SubscriptionIDByNodeIDKey
	SubscriptionsCountOfAddressKey            = types.SubscriptionsCountOfAddressKey
	SubscriptionIDByAddressKey                = types.SubscriptionIDByAddressKey
	SessionKey                                = types.SessionKey
	SessionsCountOfSubscriptionKey            = types.SessionsCountOfSubscriptionKey
	SessionIDBySubscriptionIDKey              = types.SessionIDBySubscriptionIDKey
	ActiveNodeIDsKey                          = types.ActiveNodeIDsKey
	ActiveSessionIDsKey                       = types.ActiveSessionIDsKey
	NewMsgRegisterNode                        = types.NewMsgRegisterNode
	NewMsgUpdateNodeInfo                      = types.NewMsgUpdateNodeInfo
	NewMsgUpdateNodeStatus                    = types.NewMsgUpdateNodeStatus
	NewMsgDeregisterNode                      = types.NewMsgDeregisterNode
	DefaultFreeNodesCount                     = types.DefaultFreeNodesCount
	DefaultDeposit                            = types.DefaultDeposit
	DefaultNodeInactiveInterval               = types.DefaultNodeInactiveInterval
	DefaultSessionInactiveInterval            = types.DefaultSessionInactiveInterval
	KeyFreeNodesCount                         = types.KeyFreeNodesCount
	KeyDeposit                                = types.KeyDeposit
	KeyNodeInactiveInterval                   = types.KeyNodeInactiveInterval
	KeySessionInactiveInterval                = types.KeySessionInactiveInterval
	NewParams                                 = types.NewParams
	DefaultParams                             = types.DefaultParams
	NewMsgUpdateSessionInfo                   = types.NewMsgUpdateSessionInfo
	NewBandwidthSignatureData                 = types.NewBandwidthSignatureData
	NewMsgStartSubscription                   = types.NewMsgStartSubscription
	NewMsgEndSubscription                     = types.NewMsgEndSubscription
	NewKeeper                                 = keeper.NewKeeper
	ParamKeyTable                             = keeper.ParamKeyTable
	NewQuerier                                = querier.NewQuerier
	NewQueryNodeParams                        = querier.NewQueryNodeParams
	NewQueryNodesOfAddressParams              = querier.NewQueryNodesOfAddressParams
	NewQuerySubscriptionParams                = querier.NewQuerySubscriptionParams
	NewQuerySubscriptionsOfNodePrams          = querier.NewQuerySubscriptionsOfNodePrams
	NewQuerySubscriptionsOfAddressParams      = querier.NewQuerySubscriptionsOfAddressParams
	NewQuerySessionsCountOfSubscriptionParams = querier.NewQuerySessionsCountOfSubscriptionParams
	NewQuerySessionParams                     = querier.NewQuerySessionParams
	NewQuerySessionsOfSubscriptionPrams       = querier.NewQuerySessionsOfSubscriptionPrams
)
