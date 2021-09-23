package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (m *Plan) GetProvider() hubtypes.ProvAddress {
	if m.Provider == "" {
		return nil
	}

	address, err := hubtypes.ProvAddressFromBech32(m.Provider)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Plan) PriceForDenom(d string) (sdk.Coin, bool) {
	for _, coin := range m.Price {
		if coin.Denom == d {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (m *Plan) Validate() error {
	if m.Id == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.Provider == "" {
		return fmt.Errorf("provider cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.Provider); err != nil {
		return errors.Wrapf(err, "invalid provider %s", m.Provider)
	}
	if m.Price != nil {
		if m.Price.Len() == 0 {
			return fmt.Errorf("price cannot be empty")
		}
		if !m.Price.IsValid() {
			return fmt.Errorf("price must be valid")
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
	if !m.Status.Equal(hubtypes.StatusActive) && !m.Status.Equal(hubtypes.StatusInactive) {
		return fmt.Errorf("status must be either active or inactive")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

type (
	Plans []Plan
)
