package types

import (
	ccsdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type CoinLocker struct {
	Address ccsdkTypes.AccAddress `json:"address"`
	Coins   ccsdkTypes.Coins      `json:"coins"`
	Locked  bool                  `json:"locked"`
}
