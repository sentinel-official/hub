package types

import (
	"encoding/json"

	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

var _ csdk.Msg = (*MsgStartSubscription)(nil)

type MsgStartSubscription struct {
	From    csdk.AccAddress `json:"from"`
	NodeID  sdk.ID          `json:"node_id"`
	Deposit csdk.Coin       `json:"deposit"`
}

func (msg MsgStartSubscription) Type() string {
	return "MsgStartSubscription"
}

func (msg MsgStartSubscription) ValidateBasic() csdk.Error {
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

func (msg MsgStartSubscription) GetSigners() []csdk.AccAddress {
	return []csdk.AccAddress{msg.From}
}

func (msg MsgStartSubscription) Route() string {
	return RouterKey
}

func NewMsgStartSubscription(from csdk.AccAddress,
	nodeID sdk.ID, deposit csdk.Coin) *MsgStartSubscription {

	return &MsgStartSubscription{
		From:    from,
		NodeID:  nodeID,
		Deposit: deposit,
	}
}

var _ csdk.Msg = (*MsgEndSubscription)(nil)

type MsgEndSubscription struct {
	From csdk.AccAddress `json:"from"`
	ID   sdk.ID          `json:"id"`
}

func (msg MsgEndSubscription) Type() string {
	return "MsgEndSubscription"
}

func (msg MsgEndSubscription) ValidateBasic() csdk.Error {
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

func (msg MsgEndSubscription) GetSigners() []csdk.AccAddress {
	return []csdk.AccAddress{msg.From}
}

func (msg MsgEndSubscription) Route() string {
	return RouterKey
}

func NewMsgEndSubscription(from csdk.AccAddress, id sdk.ID) *MsgEndSubscription {
	return &MsgEndSubscription{
		From: from,
		ID:   id,
	}
}
