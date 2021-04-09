package types

import (
	"fmt"
	"strings"

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

func (d *Deposit) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Address: %s
Coins  : %s
`), d.Address, d.Coins)
}

func (d Deposit) Validate() error {
	if _, err := sdk.AccAddressFromBech32(d.Address); err != nil {
		return err
	}
	if d.Coins == nil || !d.Coins.IsValid() {
		return fmt.Errorf("coins should not be nil or invalid")
	}

	return nil
}

type Deposits []Deposit
