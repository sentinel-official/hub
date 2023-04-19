package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *Quota) GetAccountAddress() sdk.AccAddress {
	if m.AccountAddress == "" {
		return nil
	}

	addr, err := sdk.AccAddressFromBech32(m.AccountAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Quota) Validate() error {
	if m.AccountAddress == "" {
		return fmt.Errorf("account_address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.AccountAddress); err != nil {
		return errors.Wrapf(err, "invalid account_address %s", m.AccountAddress)
	}
	if m.Allocated.IsNil() {
		return fmt.Errorf("allocated cannot be nil")
	}
	if m.Allocated.IsNegative() {
		return fmt.Errorf("allocated cannot be negative")
	}
	if m.Consumed.IsNil() {
		return fmt.Errorf("consumed cannot be nil")
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
