package v05

import (
	hubtypes "github.com/sentinel-official/hub/types"
)

type (
	GenesisPlan struct {
		Plan  Plan                   `json:"_"`
		Nodes []hubtypes.NodeAddress `json:"nodes"`
	}

	GenesisPlans []GenesisPlan

	GenesisState GenesisPlans
)
