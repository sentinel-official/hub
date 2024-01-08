package types

import (
	sdkerrors "cosmossdk.io/errors"

	deposittypes "github.com/sentinel-official/hub/v12/x/deposit/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types"
	providertypes "github.com/sentinel-official/hub/v12/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types"
)

func NewGenesisState(
	deposits deposittypes.GenesisState,
	providers *providertypes.GenesisState,
	nodes *nodetypes.GenesisState,
	plans plantypes.GenesisState,
	subscriptions *subscriptiontypes.GenesisState,
	sessions *sessiontypes.GenesisState,
) *GenesisState {
	return &GenesisState{
		Deposits:      deposits,
		Providers:     providers,
		Nodes:         nodes,
		Plans:         plans,
		Subscriptions: subscriptions,
		Sessions:      sessions,
	}
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

func (m *GenesisState) Validate() error {
	if err := deposittypes.ValidateGenesisState(m.Deposits); err != nil {
		return sdkerrors.Wrapf(err, "invalid deposit genesis")
	}
	if err := providertypes.ValidateGenesis(m.Providers); err != nil {
		return sdkerrors.Wrapf(err, "invalid provider genesis")
	}
	if err := nodetypes.ValidateGenesis(m.Nodes); err != nil {
		return sdkerrors.Wrapf(err, "invalid node genesis")
	}
	if err := plantypes.ValidateGenesis(m.Plans); err != nil {
		return sdkerrors.Wrapf(err, "invalid plan genesis")
	}
	if err := subscriptiontypes.ValidateGenesis(m.Subscriptions); err != nil {
		return sdkerrors.Wrapf(err, "invalid subscription genesis")
	}
	if err := sessiontypes.ValidateGenesis(m.Sessions); err != nil {
		return sdkerrors.Wrapf(err, "invalid session genesis")
	}

	return nil
}
