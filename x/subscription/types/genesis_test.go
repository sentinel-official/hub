package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	var (
		state *GenesisState
	)

	state = DefaultGenesisState()
	require.Equal(t, &GenesisState{
		Subscriptions: nil,
		Params:        DefaultParams(),
	}, state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		subscriptions GenesisSubscriptions
		params        Params
		state         *GenesisState
	)

	state = NewGenesisState(nil, params)
	require.Equal(t, &GenesisState{
		Subscriptions: nil,
		Params:        params,
	}, state)
	require.Len(t, state.Subscriptions, 0)

	state = NewGenesisState(subscriptions, params)
	require.Equal(t, &GenesisState{
		Subscriptions: subscriptions,
		Params:        params,
	}, state)
	require.Len(t, state.Subscriptions, 0)

	subscriptions = append(subscriptions,
		GenesisSubscription{},
	)
	state = NewGenesisState(subscriptions, params)
	require.Equal(t, &GenesisState{
		Subscriptions: subscriptions,
		Params:        params,
	}, state)
	require.Len(t, state.Subscriptions, 1)

	subscriptions = append(subscriptions,
		GenesisSubscription{},
	)
	state = NewGenesisState(subscriptions, params)
	require.Equal(t, &GenesisState{
		Subscriptions: subscriptions,
		Params:        params,
	}, state)
	require.Len(t, state.Subscriptions, 2)
}
