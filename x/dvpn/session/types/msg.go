package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgUpdateSession)(nil)
)

// MsgUpdateSession is for updating the session of a plan
type MsgUpdateSession struct {
	From         hub.NodeAddress `json:"from"`
	Subscription uint64          `json:"subscription"`
	Address      sdk.AccAddress  `json:"address"`
	Bandwidth    hub.Bandwidth   `json:"bandwidth"`
}

func NewMsgUpdateSession(from hub.NodeAddress, subscription uint64, address sdk.AccAddress, bandwidth hub.Bandwidth) MsgUpdateSession {
	return MsgUpdateSession{
		From:         from,
		Subscription: subscription,
		Address:      address,
		Bandwidth:    bandwidth,
	}
}

func (m MsgUpdateSession) Route() string {
	return RouterKey
}

func (m MsgUpdateSession) Type() string {
	return "update_session"
}

func (m MsgUpdateSession) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Subscription shouldn't be zero
	if m.Subscription == 0 {
		return ErrorInvalidField("subscription")
	}

	// Address shouldn't be nil or empty
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	// Bandwidth shouldn't be zero
	if m.Bandwidth.IsAllZero() {
		return ErrorInvalidField("bandwidth")
	}

	return nil
}

func (m MsgUpdateSession) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateSession) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
