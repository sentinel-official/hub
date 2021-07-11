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
		Providers: nil,
		Params:    DefaultParams(),
	}, state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		params    Params
		providers Providers
		state     *GenesisState
	)

	state = NewGenesisState(nil, params)
	require.Equal(t, &GenesisState{
		Providers: nil,
		Params:    params,
	}, state)
	require.Len(t, state.Providers, 0)

	state = NewGenesisState(providers, params)
	require.Equal(t, &GenesisState{
		Providers: providers,
		Params:    params,
	}, state)
	require.Len(t, state.Providers, 0)

	providers = append(providers,
		Provider{},
	)
	state = NewGenesisState(providers, params)
	require.Equal(t, &GenesisState{
		Providers: providers,
		Params:    params,
	}, state)
	require.Len(t, state.Providers, 1)

	providers = append(providers,
		Provider{},
	)
	state = NewGenesisState(providers, params)
	require.Equal(t, &GenesisState{
		Providers: providers,
		Params:    params,
	}, state)
	require.Len(t, state.Providers, 2)
}
