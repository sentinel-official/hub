package v05

type (
	GenesisState struct {
		Sessions Sessions `json:"_"`
		Params   Params   `json:"params"`
	}
)
