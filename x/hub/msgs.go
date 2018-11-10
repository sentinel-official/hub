package hub

import ccsdkTypes "github.com/cosmos/cosmos-sdk/types"

type MsgCoinLocker struct {
	LockerId string                `json:"locker_id"`
	Address  ccsdkTypes.AccAddress `json:"address"`
	Coins    ccsdkTypes.Coins      `json:"coins"`
	Locked   bool                  `json:"locked"`
}

type MsgLockCoins struct {
	LockerId string                `json:"locker_id"`
	Address  ccsdkTypes.AccAddress `json:"address"`
	Coins    ccsdkTypes.Coins      `json:"coins"`
}

type MsgReleaseCoins struct {
	LockerId string `json:"locker_id"`
}

type MsgReleaseCoinsToMany struct {
	LockerId  string                  `json:"locker_id"`
	Addresses []ccsdkTypes.AccAddress `json:"addresses"`
	Shares    []ccsdkTypes.Coins      `json:"shares"`
}
