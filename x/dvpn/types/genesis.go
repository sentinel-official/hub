package types

import (
	"github.com/sentinel-official/hub/x/dvpn/deposit"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/session"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

type GenesisState struct {
	Deposits     deposit.GenesisState      `json:"deposits"`
	Providers    provider.GenesisState     `json:"providers"`
	Nodes        node.GenesisState         `json:"nodes"`
	Subscription subscription.GenesisState `json:"subscription"`
	Sessions     session.Sessions          `json:"sessions"`
}

func NewGenesisState(deposits deposit.GenesisState, providers provider.GenesisState,
	nodes node.GenesisState, subscription subscription.GenesisState, sessions session.Sessions) GenesisState {
	return GenesisState{
		Deposits:     deposits,
		Providers:    providers,
		Nodes:        nodes,
		Subscription: subscription,
		Sessions:     sessions,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Deposits:     deposit.DefaultGenesisState(),
		Providers:    provider.DefaultGenesisState(),
		Nodes:        node.DefaultGenesisState(),
		Subscription: subscription.DefaultGenesisState(),
		Sessions:     session.DefaultGenesisState(),
	}
}
