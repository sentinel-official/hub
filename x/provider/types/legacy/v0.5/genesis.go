package v0_5

type (
	GenesisState struct {
		Providers Providers `json:"_"`
		Params    Params    `json:"params"`
	}
)
