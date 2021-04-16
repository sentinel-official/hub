package v0_5

type (
	GenesisState struct {
		Swaps  Swaps  `json:"_"`
		Params Params `json:"params"`
	}
)
