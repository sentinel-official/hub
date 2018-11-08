package ibc

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type MsgIBCReceive struct {
	hubTypes.IBCPacket
	Relayer  sdkTypes.AccAddress
	Sequence int64
}

func (msg MsgIBCReceive) Route() string {
	return "ibc"
}

func (msg MsgIBCReceive) Type() string {
	return "ibc_receive"
}

func (msg MsgIBCReceive) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgIBCReceive) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.Relayer}
}

func (msg MsgIBCReceive) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}
