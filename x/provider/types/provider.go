package types

import (
	"fmt"
	"net/url"

	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (m *Provider) GetAddress() hubtypes.ProvAddress {
	if m.Address == "" {
		return nil
	}

	address, err := hubtypes.ProvAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Provider) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if len(m.Name) > 64 {
		return fmt.Errorf("name length cannot be greater than %d", 64)
	}
	if len(m.Identity) > 64 {
		return fmt.Errorf("identity length cannot be greater than %d", 64)
	}
	if m.Website != "" {
		if len(m.Website) > 64 {
			return fmt.Errorf("website length cannot be greater than %d", 64)
		}
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return errors.Wrapf(err, "invalid website %s", m.Website)
		}
	}
	if len(m.Description) > 256 {
		return fmt.Errorf("description length cannot be greater than %d", 256)
	}

	return nil
}

type (
	Providers []Provider
)
