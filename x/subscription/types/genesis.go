package types

type GenesisSubscription struct {
	Subscription Subscription `json:"_"`
	Quotas       Quotas       `json:"quotas"`
}

type GenesisSubscriptions []GenesisSubscription

type GenesisState struct {
	Subscriptions GenesisSubscriptions `json:"_"`
	Params        Params               `json:"params"`
}

func NewGenesisState(subscriptions GenesisSubscriptions, params Params) GenesisState {
	return GenesisState{
		Subscriptions: subscriptions,
		Params:        params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Subscriptions: nil,
		Params:        DefaultParams(),
	}
}
