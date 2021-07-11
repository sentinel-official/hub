package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

func (p *Plan) GetProvider() hubtypes.ProvAddress {
	if p.Provider == "" {
		return nil
	}

	address, err := hubtypes.ProvAddressFromBech32(p.Provider)
	if err != nil {
		panic(err)
	}

	return address
}

func (p *Plan) PriceForDenom(d string) (sdk.Coin, bool) {
	for _, coin := range p.Price {
		if coin.Denom == d {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (p *Plan) Validate() error {
	if p.Id == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if p.Provider == "" {
		return fmt.Errorf("provider cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(p.Provider); err != nil {
		return errors.Wrapf(err, "invalid provider %s", p.Provider)
	}
	if p.Price != nil {
		if p.Price.Len() == 0 {
			return fmt.Errorf("price cannot be empty")
		}
		if !p.Price.IsValid() {
			return fmt.Errorf("price must be valid")
		}
	}
	if p.Validity < 0 {
		return fmt.Errorf("validity cannot be negative")
	}
	if p.Validity == 0 {
		return fmt.Errorf("validity cannot be zero")
	}
	if p.Bytes.IsNegative() {
		return fmt.Errorf("bytes cannot be negative")
	}
	if p.Bytes.IsZero() {
		return fmt.Errorf("bytes cannot be zero")
	}
	if !p.Status.Equal(hubtypes.StatusActive) && !p.Status.Equal(hubtypes.StatusInactive) {
		return fmt.Errorf("status must be either active or inactive")
	}
	if p.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

type (
	Plans []Plan
)
