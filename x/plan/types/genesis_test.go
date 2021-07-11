package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	var (
		state GenesisState
	)

	state = DefaultGenesisState()
	require.Equal(t, GenesisState(nil), state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		plans GenesisPlans
		state GenesisState
	)

	state = NewGenesisState(nil)
	require.Nil(t, state)
	require.Len(t, state, 0)

	state = NewGenesisState(plans)
	require.Equal(t, GenesisState(nil), state)
	require.Len(t, state, 0)

	plans = append(plans,
		GenesisPlan{},
	)
	state = NewGenesisState(plans)
	require.Equal(t, GenesisState(plans), state)
	require.Len(t, state, 1)

	plans = append(plans,
		GenesisPlan{},
	)
	state = NewGenesisState(plans)
	require.Equal(t, GenesisState(plans), state)
	require.Len(t, state, 2)
}
