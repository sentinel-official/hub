package types

type GenesisState = Nodes

func NewGenesisState(nodes Nodes) GenesisState {
	return nodes
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
