package types

import (
	"fmt"
)

func (p *Payout) Validate() error {
	if p.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if p.Hours < 0 {
		return fmt.Errorf("hours cannot be negative")
	}
	if p.Hours == 0 {
		return fmt.Errorf("hours cannot be zero")
	}
	if p.Price.Denom != "" {
		if p.Price.IsNil() {
			return fmt.Errorf("price cannot be nil")
		}
		if p.Price.IsNegative() {
			return fmt.Errorf("price cannot be negative")
		}
		if p.Price.IsZero() {
			return fmt.Errorf("price cannot be zero")
		}
		if !p.Price.IsValid() {
			return fmt.Errorf("price must be valid")
		}
	}
	if p.Timestamp.IsZero() {
		return fmt.Errorf("timestamp cannot be zero")
	}

	return nil
}

type (
	Payouts []Payout
)
