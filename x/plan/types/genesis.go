package types

import (
	"fmt"
)

type (
	GenesisPlans []GenesisPlan

	GenesisState GenesisPlans
)

func NewGenesisState(plans GenesisPlans) GenesisState {
	return GenesisState(plans)
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState(nil)
}

func ValidateGenesis(state GenesisState) error {
	for _, item := range state {
		if err := item.Plan.Validate(); err != nil {
			return err
		}
	}

	plans := make(map[uint64]bool)
	for _, item := range state {
		id := item.Plan.Id
		if plans[id] {
			return fmt.Errorf("duplicate plan id %d", id)
		}

		plans[id] = true
	}

	for _, item := range state {
		nodes := make(map[string]bool)
		for _, address := range item.Nodes {
			if nodes[address] {
				return fmt.Errorf("duplicate node for plan %d", item.Plan.Id)
			}

			nodes[address] = true
		}
	}

	return nil
}
