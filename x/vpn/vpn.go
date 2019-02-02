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

	KeyActiveNodeIDs    = types.KeyActiveNodeIDs
	KeyActiveSessionIDs = types.KeyActiveSessionIDs

	StatusRegistered   = types.StatusRegistered
	StatusActive       = types.StatusActive
	StatusInactive     = types.StatusInactive
	StatusDeregistered = types.StatusDeregistered
	StatusStart        = types.StatusDeregistered
	StatusEnd          = types.StatusEnd

	QueryNode         = querier.QueryNode
	QueryNodes        = querier.QueryNodes
	QueryNodesOfOwner = querier.QueryNodesOfOwner
)

type (
	NodeDetails = types.NodeDetails
	SessionDetails = types.SessionDetails

	MsgRegisterNode = types.MsgRegisterNode
	MsgUpdateNodeDetails = types.MsgUpdateNodeDetails
	MsgUpdateNodeStatus = types.MsgUpdateNodeStatus
	MsgDeregisterNode = types.MsgDeregisterNode
	MsgAddSession = types.MsgAddSession
	MsgUpdateSessionBandwidth = types.MsgUpdateSessionBandwidth

	Keeper = keeper.Keeper

	QueryNodeParams = querier.QueryNodeParams
	QueryNodesOfOwnerPrams = querier.QueryNodesOfOwnerPrams
)

var (
	RegisterCodec = types.RegisterCodec

	NewMsgRegisterNode      = types.NewMsgRegisterNode
	NewMsgUpdateNodeDetails = types.NewMsgUpdateNodeDetails
	NewMsgUpdateNodeStatus  = types.NewMsgUpdateNodeStatus
	NewMsgDeregisterNode    = types.NewMsgDeregisterNode

	NewKeeper = keeper.NewKeeper

	NewQuerier                 = querier.NewQuerier
	NewQueryNodeParams         = querier.NewQueryNodeParams
	NewQueryNodesOfOwnerParams = querier.NewQueryNodesOfOwnerParams
)
