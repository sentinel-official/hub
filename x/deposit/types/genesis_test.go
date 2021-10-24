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
	require.Nil(t, state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		deposits Deposits
		state    GenesisState
	)

	state = NewGenesisState(nil)
	require.Nil(t, state)
	require.Len(t, state, 0)

	state = NewGenesisState(deposits)
	require.Equal(t, GenesisState(nil), state)
	require.Len(t, state, 0)

	deposits = append(deposits,
		Deposit{},
	)
	state = NewGenesisState(deposits)
	require.Equal(t, GenesisState(deposits), state)
	require.Len(t, state, 1)

	deposits = append(deposits,
		Deposit{},
	)
	state = NewGenesisState(deposits)
	require.Equal(t, GenesisState(deposits), state)
	require.Len(t, state, 2)
}
