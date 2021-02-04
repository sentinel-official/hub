package types

type GenesisState struct {
	Providers Providers `json:"_"`
	Params    Params    `json:"params"`
}

func NewGenesisState(providers Providers, params Params) GenesisState {
	return GenesisState{
		Providers: providers,
		Params:    params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Providers: nil,
		Params:    DefaultParams(),
	}
}
