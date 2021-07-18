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
	plans := make(map[uint64]bool)
	for _, item := range state {
		if plans[item.Plan.Id] {
			return fmt.Errorf("found duplicate plan for id %d", item.Plan.Id)
		}

		plans[item.Plan.Id] = true
	}

	for _, item := range state {
		nodes := make(map[string]bool)
		for _, address := range item.Nodes {
			if nodes[address] {
				return fmt.Errorf("found duplicate node %s for plan %d", address, item.Plan.Id)
			}

			nodes[address] = true
		}
	}

	for _, item := range state {
		if err := item.Plan.Validate(); err != nil {
			return err
		}
	}

	return nil
}
