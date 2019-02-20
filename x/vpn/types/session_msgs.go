package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgInitSession struct {
	From         csdkTypes.AccAddress `json:"from"`
	NodeID       sdkTypes.ID          `json:"node_id"`
	AmountToLock csdkTypes.Coin       `json:"amount_to_lock"`
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
	if len(msg.AmountToLock.Denom) == 0 || !msg.AmountToLock.IsPositive() {
		return ErrorInvalidField("amount_to_lock")
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
	nodeID sdkTypes.ID, amountToLock csdkTypes.Coin) *MsgInitSession {

	return &MsgInitSession{
		From:         from,
		NodeID:       nodeID,
		AmountToLock: amountToLock,
	}
}

type MsgUpdateSessionBandwidth struct {
	From          csdkTypes.AccAddress `json:"from"`
	ID            sdkTypes.ID          `json:"id"`
	Bandwidth     sdkTypes.Bandwidth   `json:"bandwidth"`
	ClientSign    []byte               `json:"client_sign"`
	NodeOwnerSign []byte               `json:"node_owner_sign"`
}

func (msg MsgUpdateSessionBandwidth) Type() string {
	return "msg_update_session_bandwidth"
}

func (msg MsgUpdateSessionBandwidth) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID.Len() == 0 || !msg.ID.Valid() {
		return ErrorInvalidField("id")
	}
	if !msg.Bandwidth.IsPositive() {
		return ErrorInvalidField("bandwidth")
	}
	if len(msg.ClientSign) == 0 {
		return ErrorInvalidField("client_sign")
	}
	if len(msg.NodeOwnerSign) == 0 {
		return ErrorInvalidField("node_owner_sign")
	}

	return nil
}

func (msg MsgUpdateSessionBandwidth) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgUpdateSessionBandwidth) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateSessionBandwidth) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionBandwidth(from csdkTypes.AccAddress,
	id sdkTypes.ID, upload, download csdkTypes.Int,
	clientSign []byte, nodeOwnerSign []byte) *MsgUpdateSessionBandwidth {

	return &MsgUpdateSessionBandwidth{
		From: from,
		ID:   id,
		Bandwidth: sdkTypes.Bandwidth{
			Upload:   upload,
			Download: download,
		},
		ClientSign:    clientSign,
		NodeOwnerSign: nodeOwnerSign,
	}
}
