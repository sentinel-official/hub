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
		return fmt.Errorf("invalid id; expected positive value")
	}
	if _, err := hubtypes.ProvAddressFromBech32(p.Provider); err != nil {
		return err
	}
	if p.Price != nil && !p.Price.IsValid() {
		return fmt.Errorf("invalid price; expected non-nil and valid value")
	}
	if p.Validity <= 0 {
		return fmt.Errorf("invalid validity; expected positive value")
	}
	if !p.Bytes.IsPositive() {
		return fmt.Errorf("invalid bytes; expected positive value")
	}
	if !p.Status.Equal(hubtypes.StatusActive) && !p.Status.Equal(hubtypes.StatusInactive) {
		return fmt.Errorf("invalid status; exptected active or inactive")
	}
	if p.StatusAt.IsZero() {
		return fmt.Errorf("invalid status at; expected non-zero value")
	}

	return nil
}

type Plans []Plan
