package types

type GenesisSubscription struct {
	Subscription Subscription `json:"_"`
	Quotas       Quotas       `json:"quotas"`
}

type GenesisSubscriptions []GenesisSubscription

type GenesisState = GenesisSubscriptions

func NewGenesisState(subscriptions GenesisSubscriptions) GenesisState {
	return subscriptions
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
