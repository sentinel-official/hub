package v05

import (
	hubtypes "github.com/sentinel-official/hub/types"
)

type (
	Provider struct {
		Address     hubtypes.ProvAddress `json:"address"`
		Name        string               `json:"name"`
		Identity    string               `json:"identity"`
		Website     string               `json:"website"`
		Description string               `json:"description"`
	}

	Providers []Provider
)
