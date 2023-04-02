package types

import (
	"fmt"
)

type (
	GenesisState Deposits
)

func NewGenesisState(deposits Deposits) GenesisState {
	return GenesisState(deposits)
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState(nil)
}

func ValidateGenesis(state GenesisState) error {
	m := make(map[string]bool)
	for _, item := range state {
		if m[item.Address] {
			return fmt.Errorf("found a duplicate deposit for address %s", item.Address)
		}

		m[item.Address] = true
	}

	for _, item := range state {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
