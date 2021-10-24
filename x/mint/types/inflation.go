package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *Inflation) Validate() error {
	if m.Max.IsNegative() {
		return fmt.Errorf("max cannot be negative")
	}
	if m.Max.GT(sdk.OneDec()) {
		return fmt.Errorf("max cannot be greater than one")
	}
	if m.Min.IsNegative() {
		return fmt.Errorf("min cannot be negative")
	}
	if m.Min.GT(sdk.OneDec()) {
		return fmt.Errorf("min cannot be greater than one")
	}
	if m.Min.GT(m.Max) {
		return fmt.Errorf("min cannot be greater than max")
	}
	if m.RateChange.IsNegative() {
		return fmt.Errorf("rate_change cannot be negative")
	}
	if m.RateChange.GT(sdk.OneDec()) {
		return fmt.Errorf("rate_change cannot be greater than one")
	}
	if m.Timestamp.IsZero() {
		return fmt.Errorf("timestamp cannot be zero")
	}

	return nil
}
