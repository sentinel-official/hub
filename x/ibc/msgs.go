package ibc

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
	"github.com/ironman0x7b2/sentinel-hub/x/hub"
)

type IBCMsgCoinLocker struct {
	SrcChainId  string            `json:"src_chain_id"`
	DestChainId string            `json:"dest_chain_id"`
	Message     hub.MsgCoinLocker `json:"message"`
}

type IBCMsgLockCoins struct {
	SrcChainId  string           `json:"src_chain_id"`
	DestChainId string           `json:"dest_chain_id"`
	Message     hub.MsgLockCoins `json:"message"`
}

type IBCMsgReleaseCoins struct {
	SrcChainId  string              `json:"src_chain_id"`
	DestChainId string              `json:"dest_chain_id"`
	Message     hub.MsgReleaseCoins `json:"message"`
}

type IBCMsgReleaseCoinsToMany struct {
	SrcChainId  string                    `json:"src_chain_id"`
	DestChainId string                    `json:"dest_chain_id"`
	Message     hub.MsgReleaseCoinsToMany `json:"message"`
}

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
