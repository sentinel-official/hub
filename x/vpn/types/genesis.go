package types

type GenesisState struct {
	Nodes    []Node    `json:"nodes"`
	Sessions []Session `json:"sessions"`
	Params   Params    `json:"params"`
}

func NewGenesisState(nodes []Node, sessions []Session, params Params) GenesisState {
	return GenesisState{
		Nodes:    nodes,
		Sessions: sessions,
		Params:   params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}
