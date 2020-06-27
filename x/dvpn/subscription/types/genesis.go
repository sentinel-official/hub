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
	Plans GenesisPlans  `json:"plans"`
	List  Subscriptions `json:"list"`
}

func NewGenesisState(plans GenesisPlans, list Subscriptions) GenesisState {
	return GenesisState{
		Plans: plans,
		List:  list,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Plans: GenesisPlans{},
		List:  Subscriptions{},
	}
}
