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
	_ sdk.Msg = (*MsgStartPlanSubscription)(nil)
	_ sdk.Msg = (*MsgStartNodeSubscription)(nil)
	_ sdk.Msg = (*MsgEndSubscription)(nil)
)

type MsgAddPlan struct {
	From      hub.ProvAddress `json:"from"`
	Price     sdk.Coins       `json:"price"`
	Validity  time.Duration   `json:"validity"`
	Bandwidth hub.Bandwidth   `json:"bandwidth"`
	Duration  time.Duration   `json:"duration"`
}

func NewMsgAddPlan(from hub.ProvAddress, price sdk.Coins, validity time.Duration,
	bandwidth hub.Bandwidth, duration time.Duration) MsgAddPlan {
	return MsgAddPlan{
		From:      from,
		Price:     price,
		Validity:  validity,
		Bandwidth: bandwidth,
		Duration:  duration,
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

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return ErrorInvalidField("price")
	}

	// Validity can't be negative
	if m.Validity < 0 {
		return ErrorInvalidField("validity")
	}

	// Bandwidth can be zero. If not, it should be positive
	if !m.Bandwidth.IsAllZero() && !m.Bandwidth.IsValid() {
		return ErrorInvalidField("bandwidth")
	}

	// Duration can't be negative
	if m.Duration < 0 {
		return ErrorInvalidField("duration")
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
	if m.ID == 0 {
		return ErrorInvalidField("id")
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
	if m.ID == 0 {
		return ErrorInvalidField("id")
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
	if m.ID == 0 {
		return ErrorInvalidField("id")
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

type MsgStartPlanSubscription struct {
	From  sdk.AccAddress `json:"from"`
	ID    uint64         `json:"id"`
	Denom string         `json:"denom"`
}

func (m MsgStartPlanSubscription) Route() string {
	return RouterKey
}

func (m MsgStartPlanSubscription) Type() string {
	return "stat_plan_subscription"
}

func (m MsgStartPlanSubscription) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}
	if len(m.Denom) == 0 {
		return ErrorInvalidField("denom")
	}

	return nil
}

func (m MsgStartPlanSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgStartPlanSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

type MsgStartNodeSubscription struct {
	From    sdk.AccAddress  `json:"from"`
	Address hub.NodeAddress `json:"address"`
	Deposit sdk.Coin        `json:"deposit"`
}

func (m MsgStartNodeSubscription) Route() string {
	return RouterKey
}

func (m MsgStartNodeSubscription) Type() string {
	return "start_node_subscription"
}

func (m MsgStartNodeSubscription) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Address can't be nil and empty
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	// Deposit can be empty. If not, it should be valid
	if !hub.IsEmptyCoin(m.Deposit) && !m.Deposit.IsValid() {
		return ErrorInvalidField("deposit")
	}

	return nil
}

func (m MsgStartNodeSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgStartNodeSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

type MsgEndSubscription struct {
	From sdk.AccAddress `json:"from"`
	ID   uint64         `json:"id"`
}

func (m MsgEndSubscription) Route() string {
	return RouterKey
}

func (m MsgEndSubscription) Type() string {
	return "end_subscription"
}

func (m MsgEndSubscription) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}

	return nil
}

func (m MsgEndSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgEndSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}
