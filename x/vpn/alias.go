package vpn

import (
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/querier"
	"github.com/sentinel-official/hub/x/vpn/types"
)

const (
	Codespace                        = types.Codespace
	ModuleName                       = types.ModuleName
	QuerierRoute                     = types.QuerierRoute
	RouterKey                        = types.RouterKey
	StoreKeySession                  = types.StoreKeySession
	StoreKeyNode                     = types.StoreKeyNode
	StoreKeySubscription             = types.StoreKeySubscription
	StatusRegistered                 = types.StatusRegistered
	StatusActive                     = types.StatusActive
	StatusInactive                   = types.StatusInactive
	StatusDeRegistered               = types.StatusDeRegistered
	QueryNode                        = types.QueryNode
	QueryNodesOfAddress              = types.QueryNodesOfAddress
	QueryAllNodes                    = types.QueryAllNodes
	QuerySubscription                = types.QuerySubscription
	QuerySubscriptionsOfNode         = types.QuerySubscriptionsOfNode
	QuerySubscriptionsOfAddress      = types.QuerySubscriptionsOfAddress
	QueryAllSubscriptions            = types.QueryAllSubscriptions
	QuerySessionsCountOfSubscription = types.QuerySessionsCountOfSubscription
	QuerySession                     = types.QuerySession
	QuerySessionOfSubscription       = types.QuerySessionOfSubscription
	QuerySessionsOfSubscription      = types.QuerySessionsOfSubscription
	QueryAllSessions                 = types.QueryAllSessions
	DefaultParamspace                = keeper.DefaultParamspace
)

var (
	// functions aliases
	RegisterCodec                             = types.RegisterCodec
	ErrorMarshal                              = types.ErrorMarshal
	ErrorUnmarshal                            = types.ErrorUnmarshal
	ErrorUnknownMsgType                       = types.ErrorUnknownMsgType
	ErrorInvalidQueryType                     = types.ErrorInvalidQueryType
	ErrorInvalidField                         = types.ErrorInvalidField
	ErrorUnauthorized                         = types.ErrorUnauthorized
	ErrorNodeDoesNotExist                     = types.ErrorNodeDoesNotExist
	ErrorInvalidNodeStatus                    = types.ErrorInvalidNodeStatus
	ErrorInvalidDeposit                       = types.ErrorInvalidDeposit
	ErrorSubscriptionDoesNotExist             = types.ErrorSubscriptionDoesNotExist
	ErrorSubscriptionAlreadyExists            = types.ErrorSubscriptionAlreadyExists
	ErrorInvalidSubscriptionStatus            = types.ErrorInvalidSubscriptionStatus
	ErrorInvalidBandwidth                     = types.ErrorInvalidBandwidth
	ErrorInvalidBandwidthSignature            = types.ErrorInvalidBandwidthSignature
	ErrorSessionAlreadyExists                 = types.ErrorSessionAlreadyExists
	ErrorInvalidSessionStatus                 = types.ErrorInvalidSessionStatus
	NewGenesisState                           = types.NewGenesisState
	DefaultGenesisState                       = types.DefaultGenesisState
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
	NewParams                                 = types.NewParams
	DefaultParams                             = types.DefaultParams
	NewQueryNodeParams                        = types.NewQueryNodeParams
	NewQueryNodesOfAddressParams              = types.NewQueryNodesOfAddressParams
	NewQuerySubscriptionParams                = types.NewQuerySubscriptionParams
	NewQuerySubscriptionsOfNodePrams          = types.NewQuerySubscriptionsOfNodePrams
	NewQuerySubscriptionsOfAddressParams      = types.NewQuerySubscriptionsOfAddressParams
	NewQuerySessionsCountOfSubscriptionParams = types.NewQuerySessionsCountOfSubscriptionParams
	NewQuerySessionParams                     = types.NewQuerySessionParams
	NewQuerySessionOfSubscriptionPrams        = types.NewQuerySessionOfSubscriptionPrams
	NewQuerySessionsOfSubscriptionPrams       = types.NewQuerySessionsOfSubscriptionPrams
	NewMsgUpdateSessionInfo                   = types.NewMsgUpdateSessionInfo
	NewBandwidthSignatureData                 = types.NewBandwidthSignatureData
	NewMsgStartSubscription                   = types.NewMsgStartSubscription
	NewMsgEndSubscription                     = types.NewMsgEndSubscription
	NewKeeper                                 = keeper.NewKeeper
	ParamKeyTable                             = keeper.ParamKeyTable
	NewQuerier                                = querier.NewQuerier

	// variable aliases
	ModuleCdc                            = types.ModuleCdc
	NodesCountKey                        = types.NodesCountKey
	NodeKeyPrefix                        = types.NodeKeyPrefix
	NodesCountOfAddressKeyPrefix         = types.NodesCountOfAddressKeyPrefix
	NodeIDByAddressKeyPrefix             = types.NodeIDByAddressKeyPrefix
	SubscriptionsCountKey                = types.SubscriptionsCountKey
	SubscriptionKeyPrefix                = types.SubscriptionKeyPrefix
	SubscriptionsCountOfNodeKeyPrefix    = types.SubscriptionsCountOfNodeKeyPrefix
	SubscriptionIDByNodeIDKeyPrefix      = types.SubscriptionIDByNodeIDKeyPrefix
	SubscriptionsCountOfAddressKeyPrefix = types.SubscriptionsCountOfAddressKeyPrefix
	SubscriptionIDByAddressKeyPrefix     = types.SubscriptionIDByAddressKeyPrefix
	SessionsCountKey                     = types.SessionsCountKey
	SessionKeyPrefix                     = types.SessionKeyPrefix
	SessionsCountOfSubscriptionKeyPrefix = types.SessionsCountOfSubscriptionKeyPrefix
	SessionIDBySubscriptionIDKeyPrefix   = types.SessionIDBySubscriptionIDKeyPrefix
	DefaultFreeNodesCount                = types.DefaultFreeNodesCount
	DefaultDeposit                       = types.DefaultDeposit
	DefaultNodeInactiveInterval          = types.DefaultNodeInactiveInterval
	DefaultSessionInactiveInterval       = types.DefaultSessionInactiveInterval
	KeyFreeNodesCount                    = types.KeyFreeNodesCount
	KeyDeposit                           = types.KeyDeposit
	KeyNodeInactiveInterval              = types.KeyNodeInactiveInterval
	KeySessionInactiveInterval           = types.KeySessionInactiveInterval
)

type (
	GenesisState                           = types.GenesisState
	Node                                   = types.Node
	MsgRegisterNode                        = types.MsgRegisterNode
	MsgUpdateNodeInfo                      = types.MsgUpdateNodeInfo
	MsgUpdateNodeStatus                    = types.MsgUpdateNodeStatus
	MsgDeregisterNode                      = types.MsgDeregisterNode
	Params                                 = types.Params
	QueryNodeParams                        = types.QueryNodeParams
	QueryNodesOfAddressPrams               = types.QueryNodesOfAddressPrams
	QuerySubscriptionParams                = types.QuerySubscriptionParams
	QuerySubscriptionsOfNodePrams          = types.QuerySubscriptionsOfNodePrams
	QuerySubscriptionsOfAddressParams      = types.QuerySubscriptionsOfAddressParams
	QuerySessionsCountOfSubscriptionParams = types.QuerySessionsCountOfSubscriptionParams
	QuerySessionParams                     = types.QuerySessionParams
	QuerySessionOfSubscriptionPrams        = types.QuerySessionOfSubscriptionPrams
	QuerySessionsOfSubscriptionPrams       = types.QuerySessionsOfSubscriptionPrams
	Session                                = types.Session
	MsgUpdateSessionInfo                   = types.MsgUpdateSessionInfo
	Subscription                           = types.Subscription
	BandwidthSignatureData                 = types.BandwidthSignatureData
	MsgStartSubscription                   = types.MsgStartSubscription
	MsgEndSubscription                     = types.MsgEndSubscription
	Keeper                                 = keeper.Keeper
)
