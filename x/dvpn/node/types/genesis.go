package types

type GenesisState struct {
	Nodes  Nodes  `json:"_"`
	Params Params `json:"params"`
}

func NewGenesisState(nodes Nodes, params Params) GenesisState {
	return GenesisState{
		Nodes:  nodes,
		Params: params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Nodes:  nil,
		Params: DefaultParams(),
	}
}
