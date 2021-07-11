package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (d *Deposit) GetAddress() sdk.AccAddress {
	if d.Address == "" {
		return nil
	}

	address, err := sdk.AccAddressFromBech32(d.Address)
	if err != nil {
		panic(err)
	}

	return address
}

func (d *Deposit) Validate() error {
	if d.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(d.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", d.Address)
	}
	if d.Coins == nil {
		return fmt.Errorf("coins cannot be nil")
	}
	if d.Coins.Len() == 0 {
		return fmt.Errorf("coins cannot be empty")
	}
	if !d.Coins.IsValid() {
		return fmt.Errorf("coins must be valid")
	}

	return nil
}

type (
	Deposits []Deposit
)
