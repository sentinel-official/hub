package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgAddSession struct {
	From         csdkTypes.AccAddress `json:"from"`
	NodeID       string               `json:"node_id"`
	AmountToLock csdkTypes.Coin       `json:"amount_to_lock"`
}

func (msg MsgAddSession) Type() string {
	return "msg_add_session"
}

func (msg MsgAddSession) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if len(msg.NodeID) == 0 {
		return ErrorInvalidField("node_id")
	}
	if msg.AmountToLock.IsPositive() == false {
		return ErrorInvalidField("amount_to_lock")
	}

	return nil
}

func (msg MsgAddSession) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgAddSession) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgAddSession) Route() string {
	return RouterKey
}

func NewMsgAddSession(from csdkTypes.AccAddress,
	nodeID string, amountToLock csdkTypes.Coin) *MsgAddSession {

	return &MsgAddSession{
		From:         from,
		NodeID:       nodeID,
		AmountToLock: amountToLock,
	}
}

type MsgUpdateSessionBandwidth struct {
	From              csdkTypes.AccAddress `json:"from"`
	SessionID         string               `json:"session_id"`
	ConsumedBandwidth sdkTypes.Bandwidth   `json:"consumed_bandwidth"`
	ClientSign        []byte               `json:"client_sign"`
	NodeOwnerSign     []byte               `json:"node_owner_sign"`
}

func (msg MsgUpdateSessionBandwidth) Type() string {
	return "msg_update_session_bandwidth"
}

func (msg MsgUpdateSessionBandwidth) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if len(msg.SessionID) == 0 {
		return ErrorInvalidField("session_id")
	}
	if msg.ConsumedBandwidth.Upload == 0 || msg.ConsumedBandwidth.Download == 0 {
		return ErrorInvalidField("consumed_bandwidth")
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
	sessionID string, consumedUpload, consumedDownload uint64,
	clientSign []byte, nodeOwnerSign []byte) *MsgUpdateSessionBandwidth {

	return &MsgUpdateSessionBandwidth{
		From:      from,
		SessionID: sessionID,
		ConsumedBandwidth: sdkTypes.Bandwidth{
			Upload:   consumedUpload,
			Download: consumedDownload,
		},
		ClientSign:    clientSign,
		NodeOwnerSign: nodeOwnerSign,
	}
}
