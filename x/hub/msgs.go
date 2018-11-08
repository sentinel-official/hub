package hub

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type MsgLockCoins struct {
	FromAddress sdkTypes.AccAddress `json:"from_address"`
	ChainId     string              `json:"chain_id"`
	LockId      string              `json:"lock_id"`
	Address     sdkTypes.AccAddress `json:"address"`
	Coins       sdkTypes.Coins      `json:"coins"`
}

func (msg MsgLockCoins) Type() string {
	return "lock_coins"
}

func (msg MsgLockCoins) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgLockCoins) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.FromAddress}
}

func (msg MsgLockCoins) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}

type MsgUnlockCoins struct {
	FromAddress sdkTypes.AccAddress `json:"from_address"`
	ChainId     string              `json:"chain_id"`
	LockId      string              `json:"lock_id"`
}

func (msg MsgUnlockCoins) Type() string {
	return "unlock_coins"
}

func (msg MsgUnlockCoins) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgUnlockCoins) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.FromAddress}
}

func (msg MsgUnlockCoins) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}

type MsgSplitUnlockCoins struct {
	FromAddress sdkTypes.AccAddress    `json:"from_address"`
	ChainId     string                 `json:"chain_id"`
	LockId      string                 `json:"lock_id"`
	Splits      []hubTypes.LockedCoins `json:"splits"`
}

func (msg MsgSplitUnlockCoins) Type() string {
	return "split_unlock_coins"
}

func (msg MsgSplitUnlockCoins) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgSplitUnlockCoins) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.FromAddress}
}

func (msg MsgSplitUnlockCoins) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}
