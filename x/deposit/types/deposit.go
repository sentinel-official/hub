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
  Account Address: %s
  Coins:           %s`, d.Address, d.Coins)
}
