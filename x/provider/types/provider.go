package types

import (
	"fmt"
	"net/url"

	"github.com/cosmos/cosmos-sdk/types/errors"

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

func (p *Provider) Validate() error {
	if p.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(p.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", p.Address)
	}
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if len(p.Name) > 64 {
		return fmt.Errorf("name length cannot be greater than %d", 64)
	}
	if len(p.Identity) > 64 {
		return fmt.Errorf("identity length cannot be greater than %d", 64)
	}
	if p.Website != "" {
		if len(p.Website) > 64 {
			return fmt.Errorf("website length cannot be greater than %d", 64)
		}
		if _, err := url.ParseRequestURI(p.Website); err != nil {
			return errors.Wrapf(err, "invalid website %s", p.Website)
		}
	}
	if len(p.Description) > 256 {
		return fmt.Errorf("description length cannot be greater than %d", 256)
	}

	return nil
}

type (
	Providers []Provider
)
