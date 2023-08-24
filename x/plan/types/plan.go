package types

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (m *Plan) GetProviderAddress() hubtypes.ProvAddress {
	if m.ProviderAddress == "" {
		return nil
	}

	addr, err := hubtypes.ProvAddressFromBech32(m.ProviderAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Plan) Price(denom string) (sdk.Coin, bool) {
	for _, v := range m.Prices {
		if v.Denom == denom {
			return v, true
		}
	}

	return sdk.Coin{}, false
}

func (m *Plan) Validate() error {
	if m.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.ProviderAddress == "" {
		return fmt.Errorf("provider_address cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.ProviderAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid provider_address %s", m.ProviderAddress)
	}
	if m.Duration < 0 {
		return fmt.Errorf("duration cannot be negative")
	}
	if m.Duration == 0 {
		return fmt.Errorf("duration cannot be zero")
	}
	if m.Gigabytes < 0 {
		return fmt.Errorf("gigabytes cannot be negative")
	}
	if m.Gigabytes == 0 {
		return fmt.Errorf("gigabytes cannot be zero")
	}
	if m.Prices == nil {
		return fmt.Errorf("prices cannot be nil")
	}
	if m.Prices.Len() == 0 {
		return fmt.Errorf("prices cannot be empty")
	}
	if m.Prices.IsAnyNil() {
		return fmt.Errorf("prices cannot contain nil")
	}
	if !m.Prices.IsValid() {
		return fmt.Errorf("prices must be valid")
	}
	if !m.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactive) {
		return fmt.Errorf("status must be one of [active, inactive]")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

type (
	Plans []Plan
)
