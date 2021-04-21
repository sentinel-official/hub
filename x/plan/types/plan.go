package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

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
		return fmt.Errorf("id should not be zero")
	}
	if _, err := hubtypes.ProvAddressFromBech32(p.Provider); err != nil {
		return fmt.Errorf("provider should not be nil or empty")
	}
	if p.Price != nil && !p.Price.IsValid() {
		return fmt.Errorf("price should be either nil or valid")
	}
	if p.Validity <= 0 {
		return fmt.Errorf("validity should be positive")
	}
	if !p.Bytes.IsPositive() {
		return fmt.Errorf("bytes should be positive")
	}
	if !p.Status.Equal(hubtypes.StatusActive) && !p.Status.Equal(hubtypes.StatusInactive) {
		return fmt.Errorf("status should be either active or inactive")
	}
	if p.StatusAt.IsZero() {
		return fmt.Errorf("status_at should not be zero")
	}

	return nil
}

type Plans []Plan
