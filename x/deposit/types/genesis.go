package types

type (
	GenesisState Deposits
)

func NewGenesisState(deposits Deposits) GenesisState {
	return GenesisState(deposits)
}

func DefaultGenesisState() GenesisState {
	return nil
}
