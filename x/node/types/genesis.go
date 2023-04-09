package types

import (
	"fmt"
)

func NewGenesisState(nodes Nodes, params Params) *GenesisState {
	return &GenesisState{
		Nodes:  nodes,
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	m := make(map[string]bool)
	for _, node := range state.Nodes {
		if m[node.Address] {
			return fmt.Errorf("found duplicate node for address %s", node.Address)
		}

		m[node.Address] = true
	}

	for _, item := range state.Nodes {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
