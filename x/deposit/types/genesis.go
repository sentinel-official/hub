package types

type GenesisState = []Deposit

func NewGenesisState(deposits []Deposit) GenesisState {
	return deposits
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
