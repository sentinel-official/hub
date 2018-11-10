package hub

import csdkTypes "github.com/cosmos/cosmos-sdk/types"

type MsgCoinLocker struct {
	LockerId string               `json:"locker_id"`
	Address  csdkTypes.AccAddress `json:"address"`
	Coins    csdkTypes.Coins      `json:"coins"`
	Locked   bool                 `json:"locked"`
}

type MsgLockCoins struct {
	LockerId string               `json:"locker_id"`
	Address  csdkTypes.AccAddress `json:"address"`
	Coins    csdkTypes.Coins      `json:"coins"`
}

type MsgReleaseCoins struct {
	LockerId string `json:"locker_id"`
}

type MsgReleaseCoinsToMany struct {
	LockerId  string                 `json:"locker_id"`
	Addresses []csdkTypes.AccAddress `json:"addresses"`
	Shares    []csdkTypes.Coins      `json:"shares"`
}
