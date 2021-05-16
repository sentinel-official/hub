package v05

type (
	GenesisState struct {
		Nodes  Nodes  `json:"_"`
		Params Params `json:"params"`
	}
)
