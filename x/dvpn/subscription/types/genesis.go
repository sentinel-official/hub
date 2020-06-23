package types

import (
	hub "github.com/sentinel-official/hub/types"
)

type GenesisPlan struct {
	Plan  Plan              `json:"plan"`
	Nodes []hub.NodeAddress `json:"nodes"`
}

type GenesisPlans []GenesisPlan

type GenesisState struct {
	Plans GenesisPlans `json:"plans"`
}

func NewGenesisState(plans GenesisPlans) GenesisState {
	return GenesisState{
		Plans: plans,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Plans: GenesisPlans{},
	}
}
