package types

import (
	"fmt"
	"strings"

	hub "github.com/sentinel-official/hub/types"
)

type Provider struct {
	Address     hub.ProvAddress `json:"address"`
	Name        string          `json:"name"`
	Identity    string          `json:"identity"`
	Website     string          `json:"website"`
	Description string          `json:"description"`
}

func (p Provider) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
Address:     %s
Name:        %s
Identity:    %s
Website:     %s
Description: %s
`, p.Address, p.Identity, p.Name, p.Website, p.Description))
}

type Providers []Provider
