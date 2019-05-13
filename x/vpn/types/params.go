package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var (
	DefaultFreeNodesCount       uint64 = 5
	DefaultFreeSessionsCount    uint64 = 5
	DefaultDeposit                     = csdkTypes.NewInt64Coin("stake", 100)
	DefaultFreeSessionBandwidth        = sdkTypes.NewBandwidth(sdkTypes.GB, sdkTypes.GB)
	DefaultNodeInactiveInterval int64  = 50
	DefaultSessionEndInterval   int64  = 25
)

var (
	KeyFreeNodesCount       = []byte("FreeNodesCount")
	KeyFreeSessionsCount    = []byte("FreeSessionsCount")
	KeyDeposit              = []byte("Deposit")
	KeyFreeSessionBandwidth = []byte("FreeSessionBandwidth")
	KeyNodeInactiveInterval = []byte("NodeInactiveInterval")
	KeySessionEndInterval   = []byte("SessionEndInterval")
)

var _ params.ParamSet = (*Params)(nil)

type Params struct {
	FreeNodesCount       uint64             `json:"free_nodes_count"`
	FreeSessionsCount    uint64             `json:"free_sessions_count"`
	Deposit              csdkTypes.Coin     `json:"deposit"`
	FreeSessionBandwidth sdkTypes.Bandwidth `json:"free_session_bandwidth"`
	NodeInactiveInterval int64              `json:"node_inactive_interval"`
	SessionEndInterval   int64              `json:"session_end_interval"`
}

func NewParams(freeNodesCount, freeSessionsCount uint64,
	deposit csdkTypes.Coin, freeSessionBandwidth sdkTypes.Bandwidth,
	nodeInactiveInterval, sessionEndInterval int64) Params {

	return Params{
		FreeNodesCount:       freeNodesCount,
		FreeSessionsCount:    freeSessionsCount,
		Deposit:              deposit,
		FreeSessionBandwidth: freeSessionBandwidth,
		NodeInactiveInterval: nodeInactiveInterval,
		SessionEndInterval:   sessionEndInterval,
	}
}

func (p *Params) ParamSetPairs() subspace.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyFreeNodesCount, &p.FreeNodesCount},
		{KeyFreeSessionsCount, &p.FreeSessionsCount},
		{KeyDeposit, &p.Deposit},
		{KeyFreeSessionBandwidth, &p.FreeSessionBandwidth},
		{KeyNodeInactiveInterval, &p.NodeInactiveInterval},
		{KeySessionEndInterval, &p.SessionEndInterval},
	}
}

func DefaultParams() Params {
	return Params{
		FreeNodesCount:       DefaultFreeNodesCount,
		FreeSessionsCount:    DefaultFreeSessionsCount,
		Deposit:              DefaultDeposit,
		FreeSessionBandwidth: DefaultFreeSessionBandwidth,
		NodeInactiveInterval: DefaultNodeInactiveInterval,
		SessionEndInterval:   DefaultSessionEndInterval,
	}
}
