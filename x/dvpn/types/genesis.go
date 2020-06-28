package types

import (
	"github.com/sentinel-official/hub/x/dvpn/deposit"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/plan"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/session"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

type GenesisState struct {
	Deposits      deposit.GenesisState      `json:"deposits"`
	Providers     provider.GenesisState     `json:"providers"`
	Nodes         node.GenesisState         `json:"nodes"`
	Plans         plan.GenesisState         `json:"plans"`
	Subscriptions subscription.GenesisState `json:"subscriptions"`
	Sessions      session.Sessions          `json:"sessions"`
}

func NewGenesisState(deposits deposit.GenesisState, providers provider.GenesisState, nodes node.GenesisState,
	plans plan.GenesisState, subscriptions subscription.GenesisState, sessions session.Sessions) GenesisState {
	return GenesisState{
		Deposits:      deposits,
		Providers:     providers,
		Nodes:         nodes,
		Plans:         plans,
		Subscriptions: subscriptions,
		Sessions:      sessions,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Deposits:      deposit.DefaultGenesisState(),
		Providers:     provider.DefaultGenesisState(),
		Nodes:         node.DefaultGenesisState(),
		Plans:         plan.DefaultGenesisState(),
		Subscriptions: subscription.DefaultGenesisState(),
		Sessions:      session.DefaultGenesisState(),
	}
}
