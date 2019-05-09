package vpn

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/querier"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	StoreKeyNode    = types.StoreKeyNode
	StoreKeySession = types.StoreKeySession

	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute

	StatusRegistered   = types.StatusRegistered
	StatusActive       = types.StatusActive
	StatusInactive     = types.StatusInactive
	StatusDeregistered = types.StatusDeregistered
	StatusInit         = types.StatusInit
	StatusEnd          = types.StatusEnd

	QueryNode         = querier.QueryNode
	QueryNodesOfOwner = querier.QueryNodesOfOwner
)

type (
	Node = types.Node
	Session = types.Session
	SessionBandwidthInfo = types.SessionBandwidthInfo

	MsgRegisterNode = types.MsgRegisterNode
	MsgUpdateNodeDetails = types.MsgUpdateNodeDetails
	MsgUpdateNodeStatus = types.MsgUpdateNodeStatus
	MsgDeregisterNode = types.MsgDeregisterNode
	MsgInitSession = types.MsgInitSession
	MsgUpdateSessionBandwidthInfo = types.MsgUpdateSessionBandwidthInfo

	Keeper = keeper.Keeper

	QueryNodeParams = querier.QueryNodeParams
	QueryNodesOfOwnerPrams = querier.QueryNodesOfOwnerPrams
)

var (
	NodeKeyPrefix          = types.NodeKeyPrefix
	NodesCountKeyPrefix    = types.NodesCountKeyPrefix
	NodeKey                = types.NodeKey
	NodesCountKey          = types.NodesCountKey
	SessionKeyPrefix       = types.SessionKeyPrefix
	SessionsCountKeyPrefix = types.SessionsCountKeyPrefix
	SessionKey             = types.SessionKey
	SessionCountKey        = types.SessionsCountKey
	RegisterCodec          = types.RegisterCodec

	KeyActiveNodeIDs = types.KeyActiveNodeIDs

	NewMsgRegisterNode               = types.NewMsgRegisterNode
	NewMsgUpdateNodeDetails          = types.NewMsgUpdateNodeDetails
	NewMsgUpdateNodeStatus           = types.NewMsgUpdateNodeStatus
	NewMsgDeregisterNode             = types.NewMsgDeregisterNode
	NewMsgInitSession                = types.NewMsgInitSession
	NewMsgUpdateSessionBandwidthInfo = types.NewMsgUpdateSessionBandwidthInfo

	NewKeeper = keeper.NewKeeper

	NewQuerier                 = querier.NewQuerier
	NewQueryNodeParams         = querier.NewQueryNodeParams
	NewQueryNodesOfOwnerParams = querier.NewQueryNodesOfOwnerParams
)
