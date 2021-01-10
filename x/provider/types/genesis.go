package types

type (
	GenesisState = Providers
)

func NewGenesisState(providers Providers) GenesisState {
	return providers
}

func DefaultGenesisState() GenesisState {
	return nil
}
