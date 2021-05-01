package types

import (
	"fmt"
)

func NewGenesisState(providers Providers, params Params) *GenesisState {
	return &GenesisState{
		Providers: providers,
		Params:    params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	for _, provider := range state.Providers {
		if err := provider.Validate(); err != nil {
			return err
		}
	}

	providers := make(map[string]bool)
	for _, provider := range state.Providers {
		if providers[provider.Address] {
			return fmt.Errorf("found duplicate provider address %s", provider.Address)
		}

		providers[provider.Address] = true
	}

	return nil
}
