package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type GenesisPlan struct {
	Plan  Plan              `json:"_"`
	Nodes []hub.NodeAddress `json:"nodes"`
}

type GenesisPlans []GenesisPlan

type GenesisSubscription struct {
	Subscription Subscription     `json:"_"`
	Members      []sdk.AccAddress `json:"members"`
}

type GenesisSubscriptions []GenesisSubscription

type GenesisState struct {
	Plans GenesisPlans         `json:"plans"`
	List  GenesisSubscriptions `json:"list"`
}

func NewGenesisState(plans GenesisPlans, list GenesisSubscriptions) GenesisState {
	return GenesisState{
		Plans: plans,
		List:  list,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Plans: GenesisPlans{},
		List:  GenesisSubscriptions{},
	}
}
