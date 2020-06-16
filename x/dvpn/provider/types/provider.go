package types

import hub "github.com/sentinel-official/hub/types"

type Provider struct {
	ID          hub.ProviderID `json:"id"`
	Name        string         `json:"name"`
	Website     string         `json:"website"`
	Description string         `json:"description"`
}
