package types

func NewGenesisState(providers Providers, params Params) *GenesisState {
	return &GenesisState{
		Providers: providers,
		Params:    params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}
