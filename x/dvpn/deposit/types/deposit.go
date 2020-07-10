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
Coins:   %s`), d.Address, d.Coins)
}

type Deposits []Deposit
