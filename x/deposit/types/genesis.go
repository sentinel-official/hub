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
	for _, deposit := range state {
		if err := deposit.Validate(); err != nil {
			return err
		}
	}

	deposits := make(map[string]bool)
	for _, deposit := range state {
		if deposits[deposit.Address] {
			return fmt.Errorf("found duplicate deposit address %s", deposit.Address)
		}

		deposits[deposit.Address] = true
	}

	return nil
}
