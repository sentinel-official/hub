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

	sessions := make(map[uint64]bool)
	for _, item := range state.Sessions {
		if sessions[item.Id] {
			return fmt.Errorf("found duplicate session for id %d", item.Id)
		}

		sessions[item.Id] = true
	}

	for _, session := range state.Sessions {
		if err := session.Validate(); err != nil {
			return err
		}
	}

	return nil
}
