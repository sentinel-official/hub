package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
)

var (
	DefaultFreeNodesCount          uint64 = 5
	DefaultDeposit                        = csdkTypes.NewInt64Coin("stake", 100)
	DefaultNodeInactiveInterval    int64  = 50
	DefaultSessionInactiveInterval int64  = 25
)

var (
	KeyFreeNodesCount          = []byte("FreeNodesCount")
	KeyDeposit                 = []byte("Deposit")
	KeyNodeInactiveInterval    = []byte("NodeInactiveInterval")
	KeySessionInactiveInterval = []byte("SessionInactiveInterval")
)

var _ params.ParamSet = (*Params)(nil)

type Params struct {
	FreeNodesCount          uint64         `json:"free_nodes_count"`
	Deposit                 csdkTypes.Coin `json:"deposit"`
	NodeInactiveInterval    int64          `json:"node_inactive_interval"`
	SessionInactiveInterval int64          `json:"session_inactive_interval"`
}

func NewParams(freeNodesCount uint64, deposit csdkTypes.Coin,
	nodeInactiveInterval, sessionInactiveInterval int64) Params {

	return Params{
		FreeNodesCount:          freeNodesCount,
		Deposit:                 deposit,
		NodeInactiveInterval:    nodeInactiveInterval,
		SessionInactiveInterval: sessionInactiveInterval,
	}
}

func (p *Params) ParamSetPairs() subspace.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyFreeNodesCount, &p.FreeNodesCount},
		{KeyDeposit, &p.Deposit},
		{KeyNodeInactiveInterval, &p.NodeInactiveInterval},
		{KeySessionInactiveInterval, &p.SessionInactiveInterval},
	}
}

func DefaultParams() Params {
	return Params{
		FreeNodesCount:          DefaultFreeNodesCount,
		Deposit:                 DefaultDeposit,
		NodeInactiveInterval:    DefaultNodeInactiveInterval,
		SessionInactiveInterval: DefaultSessionInactiveInterval,
	}
}
