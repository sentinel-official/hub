package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ csdkTypes.Msg = (*MsgStartSubscription)(nil)

type MsgStartSubscription struct {
	From    csdkTypes.AccAddress `json:"from"`
	NodeID  uint64               `json:"node_id"`
	Deposit csdkTypes.Coin       `json:"deposit"`
}

func (msg MsgStartSubscription) Type() string {
	return "MsgStartSubscription"
}

func (msg MsgStartSubscription) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if len(msg.Deposit.Denom) == 0 || !msg.Deposit.IsPositive() {
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

func (msg MsgStartSubscription) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgStartSubscription) Route() string {
	return RouterKey
}

func NewMsgStartSubscription(from csdkTypes.AccAddress,
	nodeID uint64, deposit csdkTypes.Coin) *MsgStartSubscription {

	return &MsgStartSubscription{
		From:    from,
		NodeID:  nodeID,
		Deposit: deposit,
	}
}

var _ csdkTypes.Msg = (*MsgEndSubscription)(nil)

type MsgEndSubscription struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   uint64               `json:"id"`
}

func (msg MsgEndSubscription) Type() string {
	return "MsgEndSubscription"
}

func (msg MsgEndSubscription) ValidateBasic() csdkTypes.Error {
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

func (msg MsgEndSubscription) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgEndSubscription) Route() string {
	return RouterKey
}

func NewMsgEndSubscription(from csdkTypes.AccAddress, id uint64) *MsgEndSubscription {
	return &MsgEndSubscription{
		From: from,
		ID:   id,
	}
}
