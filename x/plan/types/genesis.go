package types

type GenesisPlans []GenesisPlan

type (
	GenesisState = GenesisPlans
)

func NewGenesisState(plans GenesisPlans) GenesisState {
	return plans
}

func DefaultGenesisState() GenesisState {
	return nil
}
