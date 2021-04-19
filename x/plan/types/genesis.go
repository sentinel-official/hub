package types

type (
	GenesisPlans []GenesisPlan

	GenesisState GenesisPlans
)

func NewGenesisState(plans GenesisPlans) GenesisState {
	return GenesisState(plans)
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState(nil)
}
