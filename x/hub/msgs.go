package hub

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type MsgLockCoins struct {
	FromAddress sdkTypes.AccAddress `json:"from_address"`
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

type MsgReleaseCoins struct {
	FromAddress sdkTypes.AccAddress `json:"from_address"`
	LockId      string              `json:"lock_id"`
}

func (msg MsgReleaseCoins) Type() string {
	return "release_coins"
}

func (msg MsgReleaseCoins) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgReleaseCoins) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.FromAddress}
}

func (msg MsgReleaseCoins) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}
