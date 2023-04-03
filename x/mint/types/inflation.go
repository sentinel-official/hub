package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (i *Inflation) Validate() error {
	if i.Max.IsNegative() {
		return fmt.Errorf("max cannot be negative")
	}
	if i.Max.GT(sdk.OneDec()) {
		return fmt.Errorf("max cannot be greater than one")
	}
	if i.Min.IsNegative() {
		return fmt.Errorf("min cannot be negative")
	}
	if i.Min.GT(sdk.OneDec()) {
		return fmt.Errorf("min cannot be greater than one")
	}
	if i.Min.GT(i.Max) {
		return fmt.Errorf("min cannot be greater than max")
	}
	if i.RateChange.IsNegative() {
		return fmt.Errorf("rate_change cannot be negative")
	}
	if i.RateChange.GT(sdk.OneDec()) {
		return fmt.Errorf("rate_change cannot be greater than one")
	}
	if i.Timestamp.IsZero() {
		return fmt.Errorf("timestamp cannot be zero")
	}

	return nil
}
