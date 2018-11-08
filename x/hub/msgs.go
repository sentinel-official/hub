package hub

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type MsgLockCoins struct {
	Signer      sdkTypes.AccAddress `json:"signer"`
	FromChainId string              `json:"from_chain_id"`

	LockerId string              `json:"locker_id"`
	Address  sdkTypes.AccAddress `json:"address"`
	Coins    sdkTypes.Coins      `json:"coins"`
}

func (msg MsgLockCoins) Route() string {
	return msg.Type()
}

func (msg MsgLockCoins) Type() string {
	return "lock_coins"
}

func (msg MsgLockCoins) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgLockCoins) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.Signer}
}

func (msg MsgLockCoins) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}

type MsgUnlockCoins struct {
	Signer      sdkTypes.AccAddress `json:"signer"`
	FromChainId string              `json:"from_chain_id"`

	LockerId string `json:"locker_id"`
}

func (msg MsgUnlockCoins) Route() string {
	return msg.Type()
}

func (msg MsgUnlockCoins) Type() string {
	return "unlock_coins"
}

func (msg MsgUnlockCoins) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgUnlockCoins) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.Signer}
}

func (msg MsgUnlockCoins) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}

type MsgUnlockAndShareCoins struct {
	Signer      sdkTypes.AccAddress `json:"signer"`
	FromChainId string              `json:"from_chain_id"`

	LockerId string                `json:"locker_id"`
	Addrs    []sdkTypes.AccAddress `json:"addrs"`
	Shares   []sdkTypes.Coins      `json:"shares"`
}

func (msg MsgUnlockAndShareCoins) Route() string {
	return msg.Type()
}

func (msg MsgUnlockAndShareCoins) Type() string {
	return "split_unlock_coins"
}

func (msg MsgUnlockAndShareCoins) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgUnlockAndShareCoins) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.Signer}
}

func (msg MsgUnlockAndShareCoins) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}
