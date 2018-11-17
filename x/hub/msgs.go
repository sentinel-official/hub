package hub

import csdkTypes "github.com/cosmos/cosmos-sdk/types"

type MsgLockerStatus struct {
	LockerID string `json:"locker_id"`
	Status   string `json:"status"`
}

type MsgLockCoins struct {
	LockerID string               `json:"locker_id"`
	Address  csdkTypes.AccAddress `json:"address"`
	Coins    csdkTypes.Coins      `json:"coins"`
}

type MsgReleaseCoins struct {
	LockerID string `json:"locker_id"`
}

type MsgReleaseCoinsToMany struct {
	LockerID  string                 `json:"locker_id"`
	Addresses []csdkTypes.AccAddress `json:"addresses"`
	Shares    []csdkTypes.Coins      `json:"shares"`
}
