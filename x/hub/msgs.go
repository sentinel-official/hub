package hub

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

type MsgLockerStatus struct {
	LockerID string `json:"locker_id"`
	Status   string `json:"status"`
}

type MsgLockCoins struct {
	LockerID  string          `json:"locker_id"`
	Coins     csdkTypes.Coins `json:"coins"`
	PubKey    crypto.PubKey   `json:"pub_key"`
	Signature []byte          `json:"signature"`
}

func (msg MsgLockCoins) GetUnSignBytes() []byte {
	bytes, err := json.Marshal(MsgLockCoins{
		LockerID: msg.LockerID,
		Coins:    msg.Coins,
		PubKey:   msg.PubKey,
	})

	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg MsgLockCoins) Verify() bool {
	return msg.PubKey.VerifyBytes(msg.GetUnSignBytes(), msg.Signature)
}

type MsgReleaseCoins struct {
	LockerID  string        `json:"locker_id"`
	PubKey    crypto.PubKey `json:"pub_key"`
	Signature []byte        `json:"signature"`
}

func (msg MsgReleaseCoins) GetUnSignBytes() []byte {
	bytes, err := json.Marshal(MsgReleaseCoins{
		LockerID: msg.LockerID,
		PubKey:   msg.PubKey,
	})

	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg MsgReleaseCoins) Verify() bool {
	return msg.PubKey.VerifyBytes(msg.GetUnSignBytes(), msg.Signature)
}

type MsgReleaseCoinsToMany struct {
	LockerID  string                 `json:"locker_id"`
	Addresses []csdkTypes.AccAddress `json:"addresses"`
	Shares    []csdkTypes.Coins      `json:"shares"`
	PubKey    crypto.PubKey          `json:"pub_key"`
	Signature []byte                 `json:"signature"`
}

func (msg MsgReleaseCoinsToMany) GetUnSignBytes() []byte {
	bytes, err := json.Marshal(MsgReleaseCoinsToMany{
		LockerID:  msg.LockerID,
		Addresses: msg.Addresses,
		Shares:    msg.Shares,
		PubKey:    msg.PubKey,
	})

	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg MsgReleaseCoinsToMany) Verify() bool {
	return msg.PubKey.VerifyBytes(msg.GetUnSignBytes(), msg.Signature)
}
