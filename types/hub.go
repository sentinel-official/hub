package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type CoinLocker struct {
	Address csdkTypes.AccAddress `json:"address"`
	Coins   csdkTypes.Coins      `json:"coins"`
	Status  string               `json:"status"`
}

const (
	KeyCoinLocker = "coin_locker"

	StatusLock    = "LOCKED"
	StatusRelease = "RELEASED"
)
