package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *Quota) GetAddress() sdk.AccAddress {
	if m.Address == "" {
		return nil
	}

	address, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Quota) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Allocated.IsNegative() {
		return fmt.Errorf("allocated cannot be negative")
	}
	if m.Consumed.IsNegative() {
		return fmt.Errorf("consumed cannot be negative")
	}
	if m.Consumed.GT(m.Allocated) {
		return fmt.Errorf("consumed cannot be greater than allocated")
	}

	return nil
}

type (
	Quotas []Quota
)
