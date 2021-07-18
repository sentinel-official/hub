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

	nodes := make(map[string]bool)
	for _, node := range state.Nodes {
		if nodes[node.Address] {
			return fmt.Errorf("found duplicate node for address %s", node.Address)
		}

		nodes[node.Address] = true
	}

	for _, node := range state.Nodes {
		if err := node.Validate(); err != nil {
			return err
		}
	}

	return nil
}
