package ibc

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type MsgIBCTransaction struct {
	Relayer   sdkTypes.AccAddress `json:"relayer"`
	Sequence  int64               `json:"sequence"`
	IBCPacket hubTypes.IBCPacket  `json:"ibc_packet"`
}

func (msg MsgIBCTransaction) Route() string {
	return "ibc"
}

func (msg MsgIBCTransaction) Type() string {
	return "msg_ibc_transaction"
}

func (msg MsgIBCTransaction) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgIBCTransaction) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.Relayer}
}

func (msg MsgIBCTransaction) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}
