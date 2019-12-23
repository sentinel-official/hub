package types

type GenesisState struct {
	Nodes         []Node         `json:"nodes"`
	Subscriptions []Subscription `json:"subscriptions"`
	Sessions      []Session      `json:"sessions"`
	Resolvers     []Resolver     `json:"resolvers"`
	Params        Params         `json:"params"`
}

func NewGenesisState(nodes []Node, subscriptions []Subscription, sessions []Session, params Params) GenesisState {
	return GenesisState{
		Nodes:         nodes,
		Subscriptions: subscriptions,
		Sessions:      sessions,
		Params:        params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}
