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
		Nodes:  nil,
		Params: DefaultParams(),
	}, state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		nodes  Nodes
		params Params
		state  *GenesisState
	)

	state = NewGenesisState(nil, params)
	require.Equal(t, &GenesisState{
		Nodes:  nil,
		Params: params,
	}, state)
	require.Len(t, state.Nodes, 0)

	state = NewGenesisState(nodes, params)
	require.Equal(t, &GenesisState{
		Nodes:  nodes,
		Params: params,
	}, state)
	require.Len(t, state.Nodes, 0)

	nodes = append(nodes,
		Node{},
	)
	state = NewGenesisState(nodes, params)
	require.Equal(t, &GenesisState{
		Nodes:  nodes,
		Params: params,
	}, state)
	require.Len(t, state.Nodes, 1)

	nodes = append(nodes,
		Node{},
	)
	state = NewGenesisState(nodes, params)
	require.Equal(t, &GenesisState{
		Nodes:  nodes,
		Params: params,
	}, state)
	require.Len(t, state.Nodes, 2)
}
