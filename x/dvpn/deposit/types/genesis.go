package types

type GenesisState = Deposits

func NewGenesisState(deposits Deposits) GenesisState {
	return deposits
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
