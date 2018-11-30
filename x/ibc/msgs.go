package ibc

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgIBCTransaction struct {
	Relayer   csdkTypes.AccAddress `json:"relayer"`
	Sequence  uint64               `json:"sequence"`
	IBCPacket sdkTypes.IBCPacket   `json:"ibc_packet"`
}

func (msg MsgIBCTransaction) Route() string {
	return sdkTypes.KeyIBC
}

func (msg MsgIBCTransaction) Type() string {
	return "msg_ibc_transaction"
}

func (msg MsgIBCTransaction) ValidateBasic() csdkTypes.Error {
	if msg.Relayer == nil || msg.Relayer.Empty() {
		return errorEmptyRelayer()
	}

	if msg.Sequence < 0 {
		return errorInvalidIBCSequence()
	}

	if len(msg.IBCPacket.SrcChainID) == 0 {
		return errorEmptySrcChainID()
	}

	if len(msg.IBCPacket.DestChainID) == 0 {
		return errorEmptyDestChainID()
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
