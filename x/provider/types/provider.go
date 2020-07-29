package types

import (
	"fmt"
	"strings"

	hub "github.com/sentinel-official/hub/types"
)

type Provider struct {
	Address     hub.ProvAddress `json:"address"`
	Name        string          `json:"name"`
	Identity    string          `json:"identity,omitempty"`
	Website     string          `json:"website,omitempty"`
	Description string          `json:"description,omitempty"`
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
		return fmt.Errorf("address should not be nil and empty")
	}
	if len(p.Name) == 0 || len(p.Name) > 64 {
		return fmt.Errorf("name length should be (0, 64]")
	}
	if len(p.Identity) > 64 {
		return fmt.Errorf("identity length should be (0, 64]")
	}
	if len(p.Website) > 64 {
		return fmt.Errorf("website length should be (0, 64]")
	}
	if len(p.Description) > 256 {
		return fmt.Errorf("description length should be (0, 256]")
	}

	return nil
}

type Providers []Provider
