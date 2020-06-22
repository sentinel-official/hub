package types

type GenesisState = Plans

func NewGenesisState(plans Plans) GenesisState {
	return plans
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
