package types

type GenesisState []Provider

func NewGenesisState(providers []Provider) GenesisState {
	return providers
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
