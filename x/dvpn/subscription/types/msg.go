package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgAddPlan)(nil)
	_ sdk.Msg = (*MsgSetPlanStatus)(nil)
)

type MsgAddPlan struct {
	From         hub.ProvAddress `json:"from"`
	Price        sdk.Coins       `json:"price"`
	Duration     time.Duration   `json:"duration"`
	MaxBandwidth hub.Bandwidth   `json:"max_bandwidth"`
	MaxDuration  time.Duration   `json:"max_duration"`
}

func (m MsgAddPlan) Route() string {
	return RouterKey
}

func (m MsgAddPlan) Type() string {
	return "add_plan"
}

func (m MsgAddPlan) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if m.Price == nil || m.Price.IsAnyNegative() {
		return ErrorInvalidField("price")
	}
	if m.Duration < 0 {
		return ErrorInvalidField("duration")
	}
	if m.MaxDuration < 0 {
		return ErrorInvalidField("max_duration")
	}

	return nil
}

func (m MsgAddPlan) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddPlan) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

type MsgSetPlanStatus struct {
	From   hub.ProvAddress `json:"from"`
	ID     uint64          `json:"id"`
	Status hub.Status      `json:"status"`
}

func (m MsgSetPlanStatus) Route() string {
	return RouterKey
}

func (m MsgSetPlanStatus) Type() string {
	return "set_plan_status"
}

func (m MsgSetPlanStatus) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if !m.Status.IsValid() {
		return ErrorInvalidField("status")
	}

	return nil
}

func (m MsgSetPlanStatus) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSetPlanStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
