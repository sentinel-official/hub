package types

import (
	"fmt"
	"strings"
)

func (d Deposit) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Address: %s
Coins  : %s
`), d.Address, d.Coins)
}

func (d Deposit) Validate() error {
	if d.Address == "" {
		return fmt.Errorf("address should not be empty")
	}
	if d.Coins == nil || !d.Coins.IsValid() {
		return fmt.Errorf("coins should not be nil or invalid")
	}

	return nil
}

type Deposits []Deposit
