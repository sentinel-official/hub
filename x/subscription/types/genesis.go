package types

type (
	GenesisSubscriptions []GenesisSubscription
)

func NewGenesisState(subscriptions GenesisSubscriptions, params Params) *GenesisState {
	return &GenesisState{
		Subscriptions: subscriptions,
		Params:        params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}
