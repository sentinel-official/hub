package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type LockedCoins struct {
	Address sdkTypes.AccAddress `json:"address"`
	Coins   sdkTypes.Coins      `json:"coins"`
}
