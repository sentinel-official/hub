package ibc

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgIBCTransaction struct {
	Relayer   csdkTypes.AccAddress `json:"relayer"`
	Sequence  int64                `json:"sequence"`
	IBCPacket sdkTypes.IBCPacket   `json:"ibc_packet"`
}

func (msg MsgIBCTransaction) Route() string {
	return "ibc"
}

func (msg MsgIBCTransaction) Type() string {
	return "msg_ibc_transaction"
}

func (msg MsgIBCTransaction) ValidateBasic() csdkTypes.Error {
	if msg.Relayer.Empty() {
		return csdkTypes.ErrInvalidAddress("relayer address length should not be zero")
	}

	success := false

	// switch m := msg.IBCPacket.Message.(type) {
	// case hub.MsgLockCoins:
	// 	success = m.Verify()
	// case hub.MsgReleaseCoins:
	// 	success = m.Verify()
	// case hub.MsgReleaseCoinsToMany:
	// 	success = m.Verify()
	// default:
	// 	// TODO: Replace with ErrInvalidIBCPacketMsgType
	// 	return nil
	// }

	if !success {
		// TODO: Replace with ErrMsgVerificationFailed
		return nil
	}

	return nil
}

func (msg MsgIBCTransaction) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.Relayer}
}

func (msg MsgIBCTransaction) GetSignBytes() []byte {
	signBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return signBytes
}
