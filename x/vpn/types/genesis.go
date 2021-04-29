package types

import (
	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func NewGenesisState(deposits deposittypes.GenesisState, providers *providertypes.GenesisState, nodes *nodetypes.GenesisState,
	plans plantypes.GenesisState, subscriptions *subscriptiontypes.GenesisState, sessions *sessiontypes.GenesisState) *GenesisState {
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
	if err := deposittypes.ValidateGenesis(s.Deposits); err != nil {
		return err
	}
	if err := providertypes.ValidateGenesis(s.Providers); err != nil {
		return err
	}
	if err := nodetypes.ValidateGenesis(s.Nodes); err != nil {
		return err
	}
	if err := plantypes.ValidateGenesis(s.Plans); err != nil {
		return err
	}
	if err := subscriptiontypes.ValidateGenesis(s.Subscriptions); err != nil {
		return err
	}
	if err := sessiontypes.ValidateGenesis(s.Sessions); err != nil {
		return err
	}

	return nil
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		deposittypes.DefaultGenesisState(),
		providertypes.DefaultGenesisState(),
		nodetypes.DefaultGenesisState(),
		plantypes.DefaultGenesisState(),
		subscriptiontypes.DefaultGenesisState(),
		sessiontypes.DefaultGenesisState(),
	)
}
