package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgAdd)(nil)
	_ sdk.Msg = (*MsgSetStatus)(nil)
	_ sdk.Msg = (*MsgAddNode)(nil)
	_ sdk.Msg = (*MsgRemoveNode)(nil)
)

// MsgAdd is adding a subscription plan.
type MsgAdd struct {
	From     hub.ProvAddress `json:"from"`
	Price    sdk.Coins       `json:"price"`
	Validity time.Duration   `json:"validity"`
	Bytes    sdk.Int         `json:"bytes"`
}

func NewMsgAdd(from hub.ProvAddress, price sdk.Coins, validity time.Duration, bytes sdk.Int) MsgAdd {
	return MsgAdd{
		From:     from,
		Price:    price,
		Validity: validity,
		Bytes:    bytes,
	}
}

func (m MsgAdd) Route() string {
	return RouterKey
}

func (m MsgAdd) Type() string {
	return fmt.Sprintf("%s:add", ModuleName)
}

func (m MsgAdd) ValidateBasic() sdk.Error {
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

	// Bytes shouldn't be negative and zero
	if !m.Bytes.IsPositive() {
		return ErrorInvalidField("bytes")
	}

	return nil
}

func (m MsgAdd) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAdd) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

// MsgSetStatus is for updating the status of a plan.
type MsgSetStatus struct {
	From   hub.ProvAddress `json:"from"`
	ID     uint64          `json:"id"`
	Status hub.Status      `json:"status"`
}

func NewMsgSetStatus(from hub.ProvAddress, id uint64, status hub.Status) MsgSetStatus {
	return MsgSetStatus{
		From:   from,
		ID:     id,
		Status: status,
	}
}

func (m MsgSetStatus) Route() string {
	return RouterKey
}

func (m MsgSetStatus) Type() string {
	return fmt.Sprintf("%s:set_status", ModuleName)
}

func (m MsgSetStatus) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// ID shouldn't be zero
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}

	// Status should be valid
	if !m.Status.Equal(hub.StatusActive) && !m.Status.Equal(hub.StatusInactive) {
		return ErrorInvalidField("status")
	}

	return nil
}

func (m MsgSetStatus) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSetStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

// MsgAddNode is for adding a node for a plan.
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
	return fmt.Sprintf("%s:add_node", ModuleName)
}

func (m MsgAddNode) ValidateBasic() sdk.Error {
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

// MsgRemoveNode is for removing a node for a plan.
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
	return fmt.Sprintf("%s:remove_node", ModuleName)
}

func (m MsgRemoveNode) ValidateBasic() sdk.Error {
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
