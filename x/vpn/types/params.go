package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
)

var (
	DefaultFreeNodesCount          uint64 = 5
	DefaultDeposit                        = sdk.NewInt64Coin("stake", 100)
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
	FreeNodesCount          uint64   `json:"free_nodes_count"`
	Deposit                 sdk.Coin `json:"deposit"`
	NodeInactiveInterval    int64    `json:"node_inactive_interval"`
	SessionInactiveInterval int64    `json:"session_inactive_interval"`
}

func NewParams(freeNodesCount uint64, deposit sdk.Coin,
	nodeInactiveInterval, sessionInactiveInterval int64) Params {
	return Params{
		FreeNodesCount:          freeNodesCount,
		Deposit:                 deposit,
		NodeInactiveInterval:    nodeInactiveInterval,
		SessionInactiveInterval: sessionInactiveInterval,
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`Params
  Free Nodes Count:          %d
  Deposit:                   %s
  Node Inactive Interval:    %d
  Session Inactive Interval: %d`, p.FreeNodesCount, p.Deposit,
		p.NodeInactiveInterval, p.SessionInactiveInterval)
}

func (p *Params) ParamSetPairs() subspace.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyFreeNodesCount, Value: &p.FreeNodesCount},
		{Key: KeyDeposit, Value: &p.Deposit},
		{Key: KeyNodeInactiveInterval, Value: &p.NodeInactiveInterval},
		{Key: KeySessionInactiveInterval, Value: &p.SessionInactiveInterval},
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

func (p Params) Validate() error {
	if !p.Deposit.IsValid() {
		return fmt.Errorf("deposit is invalid: %s ", p.Deposit.String())
	}

	if p.NodeInactiveInterval < 0 {
		return fmt.Errorf("NodeInactiveInterval: %d should be positive interger", p.NodeInactiveInterval)
	}

	if p.SessionInactiveInterval < 0 {
		return fmt.Errorf("SessionInactiveInterval: %d should be positive interger", p.SessionInactiveInterval)
	}

	return nil
}
