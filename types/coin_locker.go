package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type LockedCoins struct {
	Address sdkTypes.AccAddress `json:"address"`
	Coins   sdkTypes.Coins      `json:"coins"`
}

type IBCMsgLockCoins struct {
	LockId  string              `json:"lock_id"`
	Address sdkTypes.AccAddress `json:"address"`
	Coins   sdkTypes.Coins      `json:"coins"`
}

type IBCMsgUnLockCoins struct {
	LockId  string              `json:"lock_id"`
	Address sdkTypes.AccAddress `json:"address"`
	Coins   sdkTypes.Coins      `json:"coins"`
}
