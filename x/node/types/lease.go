package types

import (
	"fmt"
)

func (l *Lease) Validate() error {
	if l.Bytes.IsNil() && l.Hours == 0 {
		return fmt.Errorf("[bytes, hours] cannot be empty")
	}
	if !l.Bytes.IsNil() && l.Hours != 0 {
		return fmt.Errorf("[bytes, hours] cannot be non-empty")
	}
	if !l.Bytes.IsNil() {
		if l.Bytes.IsNegative() {
			return fmt.Errorf("bytes cannot be negative")
		}
		if l.Bytes.IsZero() {
			return fmt.Errorf("bytes cannot be zero")
		}
	}
	if l.Hours != 0 {
		if l.Hours < 0 {
			return fmt.Errorf("hours cannot be negative")
		}
	}
	if !l.Price.IsNil() {
		if l.Price.IsNegative() {
			return fmt.Errorf("price cannot be negative")
		}
		if l.Price.IsZero() {
			return fmt.Errorf("price cannot be zero")
		}
		if !l.Price.IsValid() {
			return fmt.Errorf("price must be valid")
		}
	}

	return nil
}

type (
	Leases []Lease
)
