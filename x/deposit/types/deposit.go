package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	if _, err := sdk.AccAddressFromBech32(d.Address); err != nil {
		return err
	}
	if d.Coins == nil || !d.Coins.IsValid() {
		return fmt.Errorf("invalid coins; expected non-nil and valid value")
	}

	return nil
}

type (
	Deposits []Deposit
)
