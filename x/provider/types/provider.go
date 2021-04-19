package types

import (
	"fmt"
	"strings"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (p *Provider) GetAddress() hubtypes.ProvAddress {
	if p.Address == "" {
		return nil
	}

	address, err := hubtypes.ProvAddressFromBech32(p.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (p *Provider) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
Address:     %s
Name:        %s
Identity:    %s
Website:     %s
Description: %s
`, p.Address, p.Identity, p.Name, p.Website, p.Description))
}

func (p *Provider) Validate() error {
	if _, err := hubtypes.ProvAddressFromBech32(p.Address); err != nil {
		return fmt.Errorf("address should not be nil or empty")
	}
	if len(p.Name) == 0 || len(p.Name) > 64 {
		return fmt.Errorf("name length should be between 1 and 64")
	}
	if len(p.Identity) > 64 {
		return fmt.Errorf("identity length should be between 0 and 64")
	}
	if len(p.Website) > 64 {
		return fmt.Errorf("website length should be between 0 and 64")
	}
	if len(p.Description) > 256 {
		return fmt.Errorf("description length should be between 0 and 256")
	}

	return nil
}

type Providers []Provider
