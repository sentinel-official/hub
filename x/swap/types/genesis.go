package types

type GenesisState struct {
	Swaps  Swaps  `json:"_"`
	Params Params `json:"params"`
}

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
