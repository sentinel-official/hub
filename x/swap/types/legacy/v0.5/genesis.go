package v05

type (
	GenesisState struct {
		Swaps  Swaps  `json:"_"`
		Params Params `json:"params"`
	}
)
