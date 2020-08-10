package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgUpdate)(nil)
)

// MsgUpdate is for updating the session of a plan.
type MsgUpdate struct {
	From         hub.NodeAddress `json:"from"`
	Subscription uint64          `json:"subscription"`
	Address      sdk.AccAddress  `json:"address"`
	Duration     time.Duration   `json:"duration"`
	Bandwidth    hub.Bandwidth   `json:"bandwidth"`
}

func NewMsgUpdate(from hub.NodeAddress, subscription uint64,
	address sdk.AccAddress, duration time.Duration, bandwidth hub.Bandwidth) MsgUpdate {
	return MsgUpdate{
		From:         from,
		Subscription: subscription,
		Address:      address,
		Duration:     duration,
		Bandwidth:    bandwidth,
	}
}

func (m MsgUpdate) Route() string {
	return RouterKey
}

func (m MsgUpdate) Type() string {
	return fmt.Sprintf("%s:update", ModuleName)
}

func (m MsgUpdate) ValidateBasic() sdk.Error {
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

	// Duration shouldn't be zero
	if m.Duration == 0 {
		return ErrorInvalidField("duration")
	}

	// Bandwidth shouldn't be zero
	if m.Bandwidth.IsAllZero() {
		return ErrorInvalidField("bandwidth")
	}

	return nil
}

func (m MsgUpdate) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
