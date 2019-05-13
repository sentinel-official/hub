package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgInitSession struct {
	From csdkTypes.AccAddress `json:"from"`

	NodeID  sdkTypes.ID    `json:"node_id"`
	Deposit csdkTypes.Coin `json:"deposit"`
}

func (msg MsgInitSession) Type() string {
	return "msg_init_session"
}

func (msg MsgInitSession) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	if msg.NodeID.Len() == 0 || !msg.NodeID.Valid() {
		return ErrorInvalidField("node_id")
	}
	if len(msg.Deposit.Denom) == 0 || !msg.Deposit.IsPositive() {
		return ErrorInvalidField("deposit")
	}

	return nil
}

func (msg MsgInitSession) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgInitSession) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgInitSession) Route() string {
	return RouterKey
}

func NewMsgInitSession(from csdkTypes.AccAddress,
	nodeID sdkTypes.ID, deposit csdkTypes.Coin) *MsgInitSession {

	return &MsgInitSession{
		From: from,

		NodeID:  nodeID,
		Deposit: deposit,
	}
}

type MsgUpdateSessionBandwidthInfo struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   sdkTypes.ID          `json:"id"`

	Consumed      sdkTypes.Bandwidth `json:"consumed"`
	NodeOwnerSign []byte             `json:"node_owner_sign"`
	ClientSign    []byte             `json:"client_sign"`
}

func (msg MsgUpdateSessionBandwidthInfo) Type() string {
	return "msg_update_session_bandwidth_info"
}

func (msg MsgUpdateSessionBandwidthInfo) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID.Len() == 0 || !msg.ID.Valid() {
		return ErrorInvalidField("id")
	}

	if !msg.Consumed.IsPositive() {
		return ErrorInvalidField("consumed")
	}
	if len(msg.NodeOwnerSign) == 0 {
		return ErrorInvalidField("node_owner_sign")
	}
	if len(msg.ClientSign) == 0 {
		return ErrorInvalidField("client_sign")
	}

	return nil
}

func (msg MsgUpdateSessionBandwidthInfo) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgUpdateSessionBandwidthInfo) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateSessionBandwidthInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionBandwidthInfo(from csdkTypes.AccAddress,
	id sdkTypes.ID, consumed sdkTypes.Bandwidth,
	nodeOwnerSign, clientSign []byte) *MsgUpdateSessionBandwidthInfo {

	return &MsgUpdateSessionBandwidthInfo{
		From: from,
		ID:   id,

		Consumed:      consumed,
		NodeOwnerSign: nodeOwnerSign,
		ClientSign:    clientSign,
	}
}
