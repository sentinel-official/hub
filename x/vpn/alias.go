package vpn

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/querier"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	Codespace           = types.Codespace
	StoreKeyNode        = types.StoreKeyNode
	StoreKeySession     = types.StoreKeySession
	QuerierRoute        = types.QuerierRoute
	RouterKey           = types.RouterKey
	StatusRegister      = types.StatusRegister
	StatusActive        = types.StatusActive
	StatusInactive      = types.StatusInactive
	StatusDeregister    = types.StatusDeregister
	StatusInit          = types.StatusInit
	StatusEnd           = types.StatusEnd
	DefaultParamspace   = keeper.DefaultParamspace
	QueryNode           = querier.QueryNode
	QueryNodesOfAddress = querier.QueryNodesOfAddress
)

type (
	GenesisState = types.GenesisState
	Node = types.Node
	MsgRegisterNode = types.MsgRegisterNode
	MsgUpdateNodeDetails = types.MsgUpdateNodeDetails
	MsgUpdateNodeStatus = types.MsgUpdateNodeStatus
	MsgDeregisterNode = types.MsgDeregisterNode
	Params = types.Params
	SessionBandwidthInfo = types.SessionBandwidthInfo
	Session = types.Session
	MsgInitSession = types.MsgInitSession
	MsgUpdateSessionBandwidthInfo = types.MsgUpdateSessionBandwidthInfo
	Keeper = keeper.Keeper
	QueryNodeParams = querier.QueryNodeParams
	QueryNodesOfAddressPrams = querier.QueryNodesOfAddressPrams
)

var (
	RegisterCodec                    = types.RegisterCodec
	NewGenesisState                  = types.NewGenesisState
	DefaultGenesisState              = types.DefaultGenesisState
	NodeKeyPrefix                    = types.NodeKeyPrefix
	NodesCountKeyPrefix              = types.NodesCountKeyPrefix
	ActiveNodeIDsAtHeightPrefix      = types.ActiveNodeIDsAtHeightPrefix
	SessionKeyPrefix                 = types.SessionKeyPrefix
	SessionsCountKeyPrefix           = types.SessionsCountKeyPrefix
	ActiveSessionIDsAtHeightPrefix   = types.ActiveSessionIDsAtHeightPrefix
	NodeKey                          = types.NodeKey
	NodesCountKey                    = types.NodesCountKey
	ActiveNodeIDsAtHeightKey         = types.ActiveNodeIDsAtHeightKey
	SessionKey                       = types.SessionKey
	SessionCountKey                  = types.SessionsCountKey
	ActiveSessionIDsAtHeightKey      = types.ActiveSessionIDsAtHeightKey
	NewMsgRegisterNode               = types.NewMsgRegisterNode
	NewMsgUpdateNodeDetails          = types.NewMsgUpdateNodeDetails
	NewMsgUpdateNodeStatus           = types.NewMsgUpdateNodeStatus
	NewMsgDeregisterNode             = types.NewMsgDeregisterNode
	DefaultFreeNodesCount            = types.DefaultFreeNodesCount
	DefaultFreeSessionsCount         = types.DefaultFreeSessionsCount
	DefaultDeposit                   = types.DefaultDeposit
	DefaultFreeSessionBandwidth      = types.DefaultFreeSessionBandwidth
	DefaultNodeInactiveInterval      = types.DefaultNodeInactiveInterval
	DefaultSessionEndInterval        = types.DefaultSessionEndInterval
	KeyFreeNodesCount                = types.KeyFreeNodesCount
	KeyFreeSessionsCount             = types.KeyFreeSessionsCount
	KeyDeposit                       = types.KeyDeposit
	KeyFreeSessionBandwidth          = types.KeyFreeSessionBandwidth
	KeyNodeInactiveInterval          = types.KeyNodeInactiveInterval
	KeySessionEndInterval            = types.KeySessionEndInterval
	NewParams                        = types.NewParams
	DefaultParams                    = types.DefaultParams
	NewMsgInitSession                = types.NewMsgInitSession
	NewMsgUpdateSessionBandwidthInfo = types.NewMsgUpdateSessionBandwidthInfo
	NewKeeper                        = keeper.NewKeeper
	ParamKeyTable                    = keeper.ParamKeyTable
	NewQuerier                       = querier.NewQuerier
	NewQueryNodeParams               = querier.NewQueryNodeParams
	NewQueryNodesOfAddressParams     = querier.NewQueryNodesOfAddressParams
)
