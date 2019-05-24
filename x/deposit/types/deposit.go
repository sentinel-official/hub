package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Deposit struct {
	Address csdkTypes.AccAddress `json:"address"`
	Coins   csdkTypes.Coins      `json:"coins"`
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
