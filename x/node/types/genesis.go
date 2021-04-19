package types

func NewGenesisState(nodes Nodes, params Params) *GenesisState {
	return &GenesisState{
		Nodes:  nodes,
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}
