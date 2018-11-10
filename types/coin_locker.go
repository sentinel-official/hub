package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type CoinLocker struct {
	Address csdkTypes.AccAddress `json:"address"`
	Coins   csdkTypes.Coins      `json:"coins"`
	Locked  bool                 `json:"locked"`
}
