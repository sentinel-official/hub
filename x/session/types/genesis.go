package types

import (
	"fmt"
)

func NewGenesisState(sessions Sessions, params Params) *GenesisState {
	return &GenesisState{
		Sessions: sessions,
		Params:   params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	m := make(map[uint64]bool)
	for _, item := range state.Sessions {
		if m[item.ID] {
			return fmt.Errorf("found duplicate session %d", item.ID)
		}

		m[item.ID] = true
	}

	for _, item := range state.Sessions {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
