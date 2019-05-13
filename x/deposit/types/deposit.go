package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Deposit struct {
	Address csdkTypes.AccAddress `json:"address"`
	Coins   csdkTypes.Coins      `json:"coins"`
}
