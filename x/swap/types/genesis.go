package types

func NewGenesisState(swaps Swaps, params Params) GenesisState {
	return GenesisState{
		Swaps:  swaps,
		Params: params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Swaps:  nil,
		Params: DefaultParams(),
	}
}
