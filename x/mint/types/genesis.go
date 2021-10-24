package types

import (
	"fmt"
	"time"
)

func NewGenesisState(inflations []Inflation) *GenesisState {
	return &GenesisState{
		Inflations: inflations,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil)
}

func (m *GenesisState) Validate() error {
	inflations := make(map[time.Time]bool)
	for _, item := range m.Inflations {
		if inflations[item.Timestamp] {
			return fmt.Errorf("found duplicate inflation for timestamp %s", item.Timestamp)
		}

		inflations[item.Timestamp] = true
	}

	for _, item := range m.Inflations {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
