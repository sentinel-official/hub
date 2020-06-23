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
	_ sdk.Msg = (*MsgAddNode)(nil)
	_ sdk.Msg = (*MsgRemoveNode)(nil)
)

type MsgAddPlan struct {
	From         hub.ProvAddress `json:"from"`
	Price        sdk.Coins       `json:"price"`
	Duration     time.Duration   `json:"duration"`
	MaxBandwidth hub.Bandwidth   `json:"max_bandwidth"`
	MaxDuration  time.Duration   `json:"max_duration"`
}

func NewMsgAddPlan(from hub.ProvAddress, price sdk.Coins, duration time.Duration,
	maxBandwidth hub.Bandwidth, maxDuration time.Duration) MsgAddPlan {
	return MsgAddPlan{
		From:         from,
		Price:        price,
		Duration:     duration,
		MaxBandwidth: maxBandwidth,
		MaxDuration:  maxDuration,
	}
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

func NewMsgSetPlanStatus(from hub.ProvAddress, id uint64, status hub.Status) MsgSetPlanStatus {
	return MsgSetPlanStatus{
		From:   from,
		ID:     id,
		Status: status,
	}
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

type MsgAddNode struct {
	From    hub.ProvAddress `json:"from"`
	ID      uint64          `json:"id"`
	Address hub.NodeAddress `json:"address"`
}

func NewMsgAddNode(from hub.ProvAddress, id uint64, address hub.NodeAddress) MsgAddNode {
	return MsgAddNode{
		From:    from,
		ID:      id,
		Address: address,
	}
}

func (m MsgAddNode) Route() string {
	return RouterKey
}

func (m MsgAddNode) Type() string {
	return "add_node"
}

func (m MsgAddNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	return nil
}

func (m MsgAddNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

type MsgRemoveNode struct {
	From    hub.ProvAddress `json:"from"`
	ID      uint64          `json:"id"`
	Address hub.NodeAddress `json:"address"`
}

func NewMsgRemoveNode(from hub.ProvAddress, id uint64, address hub.NodeAddress) MsgRemoveNode {
	return MsgRemoveNode{
		From:    from,
		ID:      id,
		Address: address,
	}
}

func (m MsgRemoveNode) Route() string {
	return RouterKey
}

func (m MsgRemoveNode) Type() string {
	return "remove_node"
}

func (m MsgRemoveNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	return nil
}

func (m MsgRemoveNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRemoveNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
