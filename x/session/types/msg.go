package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgUpsert)(nil)
)

// MsgUpsert is for updating or inserting a session of a plan.
type MsgUpsert struct {
	From      hub.NodeAddress `json:"from"`
	ID        uint64          `json:"id"`
	Address   sdk.AccAddress  `json:"address"`
	Duration  time.Duration   `json:"duration"`
	Bandwidth hub.Bandwidth   `json:"bandwidth"`
}

func NewMsgUpsert(from hub.NodeAddress, id uint64, address sdk.AccAddress,
	duration time.Duration, bandwidth hub.Bandwidth) MsgUpsert {
	return MsgUpsert{
		From:      from,
		ID:        id,
		Address:   address,
		Duration:  duration,
		Bandwidth: bandwidth,
	}
}

func (m MsgUpsert) Route() string {
	return RouterKey
}

func (m MsgUpsert) Type() string {
	return fmt.Sprintf("%s:upsert", ModuleName)
}

func (m MsgUpsert) ValidateBasic() error {
	if m.From == nil || m.From.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// ID shouldn't be zero
	if m.ID == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address shouldn't be nil or empty
	if m.Address == nil || m.Address.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	// Duration shouldn't be zero
	if m.Duration == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "duration")
	}

	// Bandwidth shouldn't be zero
	if m.Bandwidth.IsAllZero() {
		return errors.Wrapf(ErrorInvalidField, "%s", "bandwidth")
	}

	return nil
}

func (m MsgUpsert) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpsert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
