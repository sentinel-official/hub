package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type CoinLocker struct {
	Address sdkTypes.AccAddress `json:"address"`
	Coins   sdkTypes.Coins      `json:"coins"`
	Locked  bool                `json:"locked"`
}
