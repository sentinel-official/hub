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
	_ sdk.Msg = (*MsgAddNodeForPlan)(nil)
	_ sdk.Msg = (*MsgRemoveNodeForPlan)(nil)
	_ sdk.Msg = (*MsgStartSubscription)(nil)
	_ sdk.Msg = (*MsgAddAddressForSubscription)(nil)
	_ sdk.Msg = (*MsgEndSubscription)(nil)
)

// MsgAddPlan is adding a subscription plan.
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

	// Validity shouldn't be negative and zero
	if m.Validity <= 0 {
		return ErrorInvalidField("validity")
	}

	// Bandwidth shouldn't be negative and zero
	if !m.Bandwidth.IsValid() {
		return ErrorInvalidField("bandwidth")
	}

	// Duration shouldn't be negative and zero
	if m.Duration <= 0 {
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

// MsgSetPlanStatus is for updating the status of a plan.
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

	// ID shouldn't be zero
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}

	// Status should be valid
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

// MsgAddNodeForPlan is for adding a node for a plan.
type MsgAddNodeForPlan struct {
	From    hub.ProvAddress `json:"from"`
	ID      uint64          `json:"id"`
	Address hub.NodeAddress `json:"address"`
}

func NewMsgAddNodeForPlan(from hub.ProvAddress, id uint64, address hub.NodeAddress) MsgAddNodeForPlan {
	return MsgAddNodeForPlan{
		From:    from,
		ID:      id,
		Address: address,
	}
}

func (m MsgAddNodeForPlan) Route() string {
	return RouterKey
}

func (m MsgAddNodeForPlan) Type() string {
	return "add_node_for_plan"
}

func (m MsgAddNodeForPlan) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// ID shouldn't be zero
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}

	// Address shouldn't be nil or empty
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	return nil
}

func (m MsgAddNodeForPlan) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddNodeForPlan) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

// MsgRemoveNodeForPlan is for removing a node for a plan.
type MsgRemoveNodeForPlan struct {
	From    hub.ProvAddress `json:"from"`
	ID      uint64          `json:"id"`
	Address hub.NodeAddress `json:"address"`
}

func NewMsgRemoveNodeForPlan(from hub.ProvAddress, id uint64, address hub.NodeAddress) MsgRemoveNodeForPlan {
	return MsgRemoveNodeForPlan{
		From:    from,
		ID:      id,
		Address: address,
	}
}

func (m MsgRemoveNodeForPlan) Route() string {
	return RouterKey
}

func (m MsgRemoveNodeForPlan) Type() string {
	return "remove_node_for_plan"
}

func (m MsgRemoveNodeForPlan) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// ID shouldn't be zero
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}

	// Address shouldn't be nil or empty
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	return nil
}

func (m MsgRemoveNodeForPlan) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRemoveNodeForPlan) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

// MsgStartSubscription is for starting a subscription.
type MsgStartSubscription struct {
	From sdk.AccAddress `json:"from"`

	ID    uint64 `json:"id,omitempty"`
	Denom string `json:"denom,omitempty"`

	Address hub.NodeAddress `json:"address,omitempty"`
	Deposit sdk.Coin        `json:"deposit,omitempty"`
}

func NewMsgStartSubscription(from sdk.AccAddress, id uint64, denom string,
	address hub.NodeAddress, deposit sdk.Coin) MsgStartSubscription {
	return MsgStartSubscription{
		From:    from,
		ID:      id,
		Denom:   denom,
		Address: address,
		Deposit: deposit,
	}
}

func (m MsgStartSubscription) Route() string {
	return RouterKey
}

func (m MsgStartSubscription) Type() string {
	return "start_subscription"
}

func (m MsgStartSubscription) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	if m.ID > 0 {
		// ID shouldn't be zero
		if m.ID == 0 {
			return ErrorInvalidField("id")
		}

		// Denom length should be [3, 16]
		if len(m.Denom) < 3 || len(m.Denom) > 16 {
			return ErrorInvalidField("denom")
		}

		return nil
	}

	// Address shouldn't be nil or empty
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	// Deposit can be empty. If not, it should be valid
	if !hub.IsEmptyCoin(m.Deposit) && !m.Deposit.IsValid() {
		return ErrorInvalidField("deposit")
	}

	return nil
}

func (m MsgStartSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgStartSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgAddAddressForSubscription is for adding an address for subscription.
type MsgAddAddressForSubscription struct {
	From    sdk.AccAddress `json:"from"`
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
}

func (m MsgAddAddressForSubscription) Route() string {
	return RouterKey
}

func (m MsgAddAddressForSubscription) Type() string {
	return "add_address_for_subscription"
}

func (m MsgAddAddressForSubscription) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// ID shouldn't be zero
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}

	// Address shouldn't be nil or empty
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	return nil
}

func (m MsgAddAddressForSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddAddressForSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgEndSubscription is for ending a subscription.
type MsgEndSubscription struct {
	From sdk.AccAddress `json:"from"`
	ID   uint64         `json:"id"`
}

func NewMsgEndSubscription(from sdk.AccAddress, id uint64) MsgEndSubscription {
	return MsgEndSubscription{
		From: from,
		ID:   id,
	}
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

	// ID shouldn't be zero
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
