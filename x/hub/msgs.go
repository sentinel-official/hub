package hub

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type MsgCoinLocker struct {
	LockerId string              `json:"locker_id"`
	Address  sdkTypes.AccAddress `json:"address"`
	Coins    sdkTypes.Coins      `json:"coins"`
	Locked   bool                `json:"locked"`
}

type MsgLockCoins struct {
	LockerId string              `json:"locker_id"`
	Address  sdkTypes.AccAddress `json:"address"`
	Coins    sdkTypes.Coins      `json:"coins"`
}

type MsgReleaseCoins struct {
	LockerId string `json:"locker_id"`
}

type MsgReleaseCoinsToMany struct {
	LockerId  string                `json:"locker_id"`
	Addresses []sdkTypes.AccAddress `json:"addresses"`
	Shares    []sdkTypes.Coins      `json:"shares"`
}
