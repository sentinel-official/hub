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

	addr, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Quota) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.AllocatedBytes.IsNil() {
		return fmt.Errorf("allocated_bytes cannot be nil")
	}
	if m.AllocatedBytes.IsNegative() {
		return fmt.Errorf("allocated_bytes cannot be negative")
	}
	if m.ConsumedBytes.IsNil() {
		return fmt.Errorf("consumed_bytes cannot be nil")
	}
	if m.ConsumedBytes.IsNegative() {
		return fmt.Errorf("consumed_bytes cannot be negative")
	}
	if m.ConsumedBytes.GT(m.AllocatedBytes) {
		return fmt.Errorf("consumed_bytes cannot be greater than allocated_bytes")
	}

	return nil
}

type (
	Quotas []Quota
)
