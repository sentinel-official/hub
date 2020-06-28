package types

type GenesisState = Sessions

func NewGenesisState(sessions Sessions) GenesisState {
	return sessions
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
