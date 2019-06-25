package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var _ sdk.Msg = (*MsgStartSubscription)(nil)

type MsgStartSubscription struct {
	From    sdk.AccAddress `json:"from"`
	NodeID  hub.ID         `json:"node_id"`
	Deposit sdk.Coin       `json:"deposit"`
}

func (msg MsgStartSubscription) Type() string {
	return "MsgStartSubscription"
}

func (msg MsgStartSubscription) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.Deposit.Denom == "" || !msg.Deposit.IsPositive() {
		return ErrorInvalidField("deposit")
	}

	return nil
}

func (msg MsgStartSubscription) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgStartSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgStartSubscription) Route() string {
	return RouterKey
}

func NewMsgStartSubscription(from sdk.AccAddress,
	nodeID hub.ID, deposit sdk.Coin) *MsgStartSubscription {

	return &MsgStartSubscription{
		From:    from,
		NodeID:  nodeID,
		Deposit: deposit,
	}
}

var _ sdk.Msg = (*MsgEndSubscription)(nil)

type MsgEndSubscription struct {
	From sdk.AccAddress `json:"from"`
	ID   hub.ID         `json:"id"`
}

func (msg MsgEndSubscription) Type() string {
	return "MsgEndSubscription"
}

func (msg MsgEndSubscription) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	return nil
}

func (msg MsgEndSubscription) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgEndSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgEndSubscription) Route() string {
	return RouterKey
}

func NewMsgEndSubscription(from sdk.AccAddress, id hub.ID) *MsgEndSubscription {
	return &MsgEndSubscription{
		From: from,
		ID:   id,
	}
}
