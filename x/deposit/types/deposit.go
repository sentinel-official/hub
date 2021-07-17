package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *Deposit) GetAddress() sdk.AccAddress {
	if m.Address == "" {
		return nil
	}

	address, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (m *Deposit) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", m.Address)
	}
	if m.Coins == nil {
		return fmt.Errorf("coins cannot be nil")
	}
	if m.Coins.Len() == 0 {
		return fmt.Errorf("coins cannot be empty")
	}
	if !m.Coins.IsValid() {
		return fmt.Errorf("coins must be valid")
	}

	return nil
}

type (
	Deposits []Deposit
)
