package types

import (
	"fmt"
)

func NewGenesisState(swaps Swaps, params Params) *GenesisState {
	return &GenesisState{
		Swaps:  swaps,
		Params: params,
	}
}

func (s *GenesisState) Validate() error {
	if err := s.Params.Validate(); err != nil {
		return err
	}

	for _, item := range s.Swaps {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	swaps := make(map[string]bool)
	for _, item := range s.Swaps {
		txHash := item.GetTxHash().String()
		if swaps[txHash] {
			return fmt.Errorf("duplicate swap for tx_hash %s", txHash)
		}

		swaps[txHash] = true
	}

	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Swaps:  nil,
		Params: DefaultParams(),
	}
}
