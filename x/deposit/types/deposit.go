package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Deposit struct {
	Address sdk.AccAddress `json:"address"`
	Coins   sdk.Coins      `json:"coins"`
}

func (d Deposit) String() string {
	return fmt.Sprintf(`Deposit
  Address: %s
  Coins:   %s`, d.Address, d.Coins)
}

func (d Deposit) IsValid() error {
	if d.Address == nil || d.Address.Empty() {
		return fmt.Errorf("invalid address")
	}
	if !d.Coins.IsValid() {
		return fmt.Errorf("invalid coins")
	}

	return nil
}
