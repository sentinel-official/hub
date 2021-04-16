package v0_5

type (
	GenesisState struct {
		Sessions Sessions `json:"_"`
		Params   Params   `json:"params"`
	}
)
