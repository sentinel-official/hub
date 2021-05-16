package v05

type (
	GenesisState struct {
		Providers Providers `json:"_"`
		Params    Params    `json:"params"`
	}
)
