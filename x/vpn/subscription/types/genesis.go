package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisSubscription struct {
	Subscription Subscription     `json:"_"`
	Members      []sdk.AccAddress `json:"members"`
}

type GenesisSubscriptions []GenesisSubscription

type GenesisState = GenesisSubscriptions

func NewGenesisState(subscriptions GenesisSubscriptions) GenesisState {
	return subscriptions
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
