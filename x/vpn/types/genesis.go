package types

type GenesisState struct {
	Nodes         []Node         `json:"nodes"`
	Subscriptions []Subscription `json:"subscriptions"`
	Sessions      []Session      `json:"sessions"`
	Resolvers     []Resolver     `json:"resolvers"`
	FreeClients   []FreeClient   `json:"free_clients"`
	Params        Params         `json:"params"`
}

func NewGenesisState(nodes []Node, subscriptions []Subscription, sessions []Session, resolvers []Resolver,
	freeClients []FreeClient, params Params) GenesisState {
	return GenesisState{
		Nodes:         nodes,
		Subscriptions: subscriptions,
		Sessions:      sessions,
		Resolvers:     resolvers,
		FreeClients:   freeClients,
		Params:        params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}
