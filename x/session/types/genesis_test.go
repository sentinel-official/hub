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
		Sessions: nil,
		Params:   DefaultParams(),
	}, state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		subscriptions Sessions
		params        Params
		state         *GenesisState
	)

	state = NewGenesisState(nil, params)
	require.Equal(t, &GenesisState{
		Sessions: nil,
		Params:   params,
	}, state)
	require.Len(t, state.Sessions, 0)

	state = NewGenesisState(subscriptions, params)
	require.Equal(t, &GenesisState{
		Sessions: subscriptions,
		Params:   params,
	}, state)
	require.Len(t, state.Sessions, 0)

	subscriptions = append(subscriptions,
		Session{},
	)
	state = NewGenesisState(subscriptions, params)
	require.Equal(t, &GenesisState{
		Sessions: subscriptions,
		Params:   params,
	}, state)
	require.Len(t, state.Sessions, 1)

	subscriptions = append(subscriptions,
		Session{},
	)
	state = NewGenesisState(subscriptions, params)
	require.Equal(t, &GenesisState{
		Sessions: subscriptions,
		Params:   params,
	}, state)
	require.Len(t, state.Sessions, 2)
}
