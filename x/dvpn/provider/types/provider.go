package types

import (
	"fmt"
	"strings"

	hub "github.com/sentinel-official/hub/types"
)

type Provider struct {
	ID          hub.ProviderID `json:"id"`
	Name        string         `json:"name"`
	Website     string         `json:"website"`
	Description string         `json:"description"`
}

func (p Provider) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
ID: %s
Name: %s
Website: %s
Description: %s
`, p.ID, p.Name, p.Website, p.Description))
}

type Providers []Provider
