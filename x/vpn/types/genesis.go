package types

import (
	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
)

func NewGenesisState(deposits deposit.GenesisState, providers provider.GenesisState, nodes node.GenesisState,
	plans plan.GenesisState, subscriptions subscription.GenesisState, sessions session.GenesisState) *GenesisState {
	return &GenesisState{
		Deposits:      deposits,
		Providers:     providers,
		Nodes:         nodes,
		Plans:         plans,
		Subscriptions: subscriptions,
		Sessions:      sessions,
	}
}

func (s *GenesisState) Validate() error {
	if err := deposit.ValidateGenesis(s.Deposits); err != nil {
		return err
	}
	if err := provider.ValidateGenesis(s.Providers); err != nil {
		return err
	}
	if err := node.ValidateGenesis(s.Nodes); err != nil {
		return err
	}
	if err := plan.ValidateGenesis(s.Plans); err != nil {
		return err
	}
	if err := subscription.ValidateGenesis(s.Subscriptions); err != nil {
		return err
	}
	if err := session.ValidateGenesis(s.Sessions); err != nil {
		return err
	}

	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Deposits:      deposit.DefaultGenesisState(),
		Providers:     provider.DefaultGenesisState(),
		Nodes:         node.DefaultGenesisState(),
		Plans:         plan.DefaultGenesisState(),
		Subscriptions: subscription.DefaultGenesisState(),
		Sessions:      session.DefaultGenesisState(),
	}
}
