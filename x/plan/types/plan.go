package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (m *Plan) GetAddress() hubtypes.ProvAddress {
	if m.Address == "" {
		return nil
	}

	addr, err := hubtypes.ProvAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Plan) Price(denom string) (sdk.Coin, bool) {
	if m.Prices == nil {
		return sdk.Coin{}, true
	}

	for _, coin := range m.Prices {
		if coin.Denom == denom {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (m *Plan) Validate() error {
	if m.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Prices != nil {
		if m.Prices.Len() == 0 {
			return fmt.Errorf("prices cannot be empty")
		}
		if m.Prices.IsAnyNil() {
			return fmt.Errorf("prices cannot contain nil")
		}
		if !m.Prices.IsValid() {
			return fmt.Errorf("prices must be valid")
		}
	}
	if m.Validity < 0 {
		return fmt.Errorf("validity cannot be negative")
	}
	if m.Validity == 0 {
		return fmt.Errorf("validity cannot be zero")
	}
	if m.Bytes.IsNegative() {
		return fmt.Errorf("bytes cannot be negative")
	}
	if m.Bytes.IsZero() {
		return fmt.Errorf("bytes cannot be zero")
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
