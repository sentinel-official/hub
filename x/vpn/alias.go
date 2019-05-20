package vpn

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/querier"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	Codespace            = types.Codespace
	StoreKeyNode         = types.StoreKeyNode
	StoreKeySession      = types.StoreKeySession
	StoreKeySubscription = types.StoreKeySubscription
	QuerierRoute         = types.QuerierRoute
	RouterKey            = types.RouterKey
	StatusRegistered     = types.StatusRegistered
	StatusActive         = types.StatusActive
	StatusInactive       = types.StatusInactive
	StatusDeRegistered   = types.StatusDeRegistered
	StatusStarted        = types.StatusStarted
	StatusEnded          = types.StatusEnded
	TagNodeID            = types.TagNodeID
	TagSessionID         = types.TagSessionID
	DefaultParamspace    = keeper.DefaultParamspace
	QueryNode            = querier.QueryNode
	QueryNodesOfAddress  = querier.QueryNodesOfAddress
)

type (
	GenesisState = types.GenesisState
	Node = types.Node
	MsgRegisterNode = types.MsgRegisterNode
	MsgUpdateNodeInfo = types.MsgUpdateNodeInfo
	MsgUpdateNodeStatus = types.MsgUpdateNodeStatus
	MsgDeregisterNode = types.MsgDeregisterNode
	Params = types.Params
	Session = types.Session
	MsgUpdateSessionInfo = types.MsgUpdateSessionInfo
	Subscription = types.Subscription
	MsgStartSubscription = types.MsgStartSubscription
	MsgEndSubscription = types.MsgEndSubscription
	Keeper = keeper.Keeper
	QueryNodeParams = querier.QueryNodeParams
	QueryNodesOfAddressPrams = querier.QueryNodesOfAddressPrams
)

var (
	RegisterCodec                  = types.RegisterCodec
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	NodesCountKeyPrefix            = types.NodesCountKeyPrefix
	NodeKeyPrefix                  = types.NodeKeyPrefix
	SubscriptionKeyPrefix          = types.SubscriptionKeyPrefix
	SessionKeyPrefix               = types.SessionKeyPrefix
	NodesCountKey                  = types.NodesCountKey
	NodeID                         = types.NodeID
	NodeKey                        = types.NodeKey
	SubscriptionID                 = types.SubscriptionID
	SubscriptionKey                = types.SubscriptionKey
	SessionID                      = types.SessionID
	SessionKey                     = types.SessionKey
	ActiveNodeIDsKey               = types.ActiveNodeIDsKey
	ActiveSessionIDsKey            = types.ActiveSessionIDsKey
	NewMsgRegisterNode             = types.NewMsgRegisterNode
	NewMsgUpdateNodeInfo           = types.NewMsgUpdateNodeInfo
	NewMsgUpdateNodeStatus         = types.NewMsgUpdateNodeStatus
	NewMsgDeregisterNode           = types.NewMsgDeregisterNode
	DefaultFreeNodesCount          = types.DefaultFreeNodesCount
	DefaultDeposit                 = types.DefaultDeposit
	DefaultNodeInactiveInterval    = types.DefaultNodeInactiveInterval
	DefaultSessionInactiveInterval = types.DefaultSessionInactiveInterval
	KeyFreeNodesCount              = types.KeyFreeNodesCount
	KeyDeposit                     = types.KeyDeposit
	KeyNodeInactiveInterval        = types.KeyNodeInactiveInterval
	KeySessionInactiveInterval     = types.KeySessionInactiveInterval
	NewParams                      = types.NewParams
	DefaultParams                  = types.DefaultParams
	NewMsgUpdateSessionInfo        = types.NewMsgUpdateSessionInfo
	NewMsgStartSubscription        = types.NewMsgStartSubscription
	NewMsgEndSubscription          = types.NewMsgEndSubscription
	NewKeeper                      = keeper.NewKeeper
	ParamKeyTable                  = keeper.ParamKeyTable
	NewQuerier                     = querier.NewQuerier
	NewQueryNodeParams             = querier.NewQueryNodeParams
	NewQueryNodesOfAddressParams   = querier.NewQueryNodesOfAddressParams
)
