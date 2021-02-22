package types

func NewGenesisState(sessions Sessions, params Params) GenesisState {
	return GenesisState{
		Sessions: sessions,
		Params:   params,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Sessions: nil,
		Params:   DefaultParams(),
	}
}
