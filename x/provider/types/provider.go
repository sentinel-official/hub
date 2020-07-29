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

func (p Provider) Validate() error {
	if p.Address == nil || p.Address.Empty() {
		return fmt.Errorf("invalid address; found nil or empty")
	}
	if len(p.Name) == 0 || len(p.Name) > 64 {
		return fmt.Errorf("invalid name; found length zero or greater than 64")
	}
	if len(p.Identity) > 64 {
		return fmt.Errorf("invalid identity; found length greater than 64")
	}
	if len(p.Website) > 64 {
		return fmt.Errorf("invalid website; found length greater than 64")
	}
	if len(p.Description) > 256 {
		return fmt.Errorf("invalid description; found length greater than 256")
	}

	return nil
}

type Providers []Provider
