package types

import (
	"fmt"
	"strings"

	hub "github.com/sentinel-official/hub/types"
)

type Provider struct {
	Address     hub.ProvAddress `json:"address"`
	Name        string          `json:"name"`
	Website     string          `json:"website"`
	Description string          `json:"description"`
}

func (p Provider) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
Address: %s
Name: %s
Website: %s
Description: %s
`, p.Address, p.Name, p.Website, p.Description))
}

type Providers []Provider
