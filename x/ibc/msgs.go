package ibc

import (
	"encoding/json"

	ccsdkTypes "github.com/cosmos/cosmos-sdk/types"
	csdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgIBCTransaction struct {
	Relayer   ccsdkTypes.AccAddress `json:"relayer"`
	Sequence  int64                 `json:"sequence"`
	IBCPacket csdkTypes.IBCPacket   `json:"ibc_packet"`
}

func (msg MsgIBCTransaction) Route() string {
	return "ibc"
}

func (msg MsgIBCTransaction) Type() string {
	return "msg_ibc_transaction"
}

func (msg MsgIBCTransaction) ValidateBasic() ccsdkTypes.Error {
	return nil
}

func (msg MsgIBCTransaction) GetSigners() []ccsdkTypes.AccAddress {
	return []ccsdkTypes.AccAddress{msg.Relayer}
}

func (msg MsgIBCTransaction) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}
