package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Deposit struct {
	Address sdk.AccAddress `json:"address"`
	Coins   sdk.Coins      `json:"coins"`
}

func (d Deposit) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Address: %s
Coins  : %s
`), d.Address, d.Coins)
}

func (d Deposit) Validate() error {
	if d.Address == nil || d.Address.Empty() {
		return fmt.Errorf("address should not be nil or empty")
	}
	if d.Coins == nil || !d.Coins.IsValid() {
		return fmt.Errorf("coins should not be nil or invalid")
	}

	return nil
}

type Deposits []Deposit
