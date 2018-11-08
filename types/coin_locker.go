package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type CoinLocker struct {
	Address sdkTypes.AccAddress `json:"address"`
	Coins   sdkTypes.Coins      `json:"coins"`
	Locked  bool                `json:"locked"`
}

type IBCMsgCoinLocker struct {
	LockerId string `json:"locker_id"`

	Address sdkTypes.AccAddress `json:"address"`
	Coins   sdkTypes.Coins      `json:"coins"`
	Locked  bool                `json:"locked"`
}
