package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgSubscribeToPlan)(nil)
	_ sdk.Msg = (*MsgSubscribeToNode)(nil)
	_ sdk.Msg = (*MsgCancel)(nil)

	_ sdk.Msg = (*MsgAddQuota)(nil)
	_ sdk.Msg = (*MsgUpdateQuota)(nil)
)

// MsgSubscribeToPlan is for starting a plan subscription.
type MsgSubscribeToPlan struct {
	From  sdk.AccAddress `json:"from"`
	ID    uint64         `json:"id"`
	Denom string         `json:"denom"`
}

func NewMsgSubscribeToPlan(from sdk.AccAddress, id uint64, denom string) MsgSubscribeToPlan {
	return MsgSubscribeToPlan{
		From:  from,
		ID:    id,
		Denom: denom,
	}
}

func (m MsgSubscribeToPlan) Route() string {
	return RouterKey
}

func (m MsgSubscribeToPlan) Type() string {
	return fmt.Sprintf("%s:subscribe_to_plan", ModuleName)
}

func (m MsgSubscribeToPlan) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

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

func (m MsgSubscribeToPlan) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSubscribeToPlan) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgSubscribeToNode is for starting a node subscription.
type MsgSubscribeToNode struct {
	From    sdk.AccAddress  `json:"from"`
	Address hub.NodeAddress `json:"address"`
	Deposit sdk.Coin        `json:"deposit"`
}

func NewMsgSubscribeToNode(from sdk.AccAddress, address hub.NodeAddress, deposit sdk.Coin) MsgSubscribeToNode {
	return MsgSubscribeToNode{
		From:    from,
		Address: address,
		Deposit: deposit,
	}
}

func (m MsgSubscribeToNode) Route() string {
	return RouterKey
}

func (m MsgSubscribeToNode) Type() string {
	return fmt.Sprintf("%s:subscribe_to_node", ModuleName)
}

func (m MsgSubscribeToNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Address shouldn't be nil or empty
	if m.Address == nil || m.Address.Empty() {
		return ErrorInvalidField("address")
	}

	// Deposit should be valid
	if !m.Deposit.IsValid() {
		return ErrorInvalidField("deposit")
	}

	return nil
}

func (m MsgSubscribeToNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSubscribeToNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgCancel is for cancelling a subscription.
type MsgCancel struct {
	From sdk.AccAddress `json:"from"`
	ID   uint64         `json:"id"`
}

func NewMsgCancel(from sdk.AccAddress, id uint64) MsgCancel {
	return MsgCancel{
		From: from,
		ID:   id,
	}
}

func (m MsgCancel) Route() string {
	return RouterKey
}

func (m MsgCancel) Type() string {
	return fmt.Sprintf("%s:cancel", ModuleName)
}

func (m MsgCancel) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// ID shouldn't be zero
	if m.ID == 0 {
		return ErrorInvalidField("id")
	}

	return nil
}

func (m MsgCancel) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgCancel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgAddQuota is for adding the bandwidth quota for an address.
type MsgAddQuota struct {
	From    sdk.AccAddress `json:"from"`
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
	Bytes   sdk.Int        `json:"bytes"`
}

func NewMsgAddQuota(from sdk.AccAddress, id uint64, address sdk.AccAddress, bytes sdk.Int) MsgAddQuota {
	return MsgAddQuota{
		From:    from,
		ID:      id,
		Address: address,
		Bytes:   bytes,
	}
}

func (m MsgAddQuota) Route() string {
	return RouterKey
}

func (m MsgAddQuota) Type() string {
	return fmt.Sprintf("%s:add_quota", ModuleName)
}

func (m MsgAddQuota) ValidateBasic() sdk.Error {
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

	// Bytes should be positive
	if !m.Bytes.IsPositive() {
		return ErrorInvalidField("bytes")
	}

	return nil
}

func (m MsgAddQuota) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddQuota) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgUpdateQuota is for updating the bandwidth quota for an address.
type MsgUpdateQuota struct {
	From    sdk.AccAddress `json:"from"`
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
	Bytes   sdk.Int        `json:"bytes"`
}

func NewMsgUpdateQuota(from sdk.AccAddress, id uint64, address sdk.AccAddress, bytes sdk.Int) MsgUpdateQuota {
	return MsgUpdateQuota{
		From:    from,
		ID:      id,
		Address: address,
		Bytes:   bytes,
	}
}

func (m MsgUpdateQuota) Route() string {
	return RouterKey
}

func (m MsgUpdateQuota) Type() string {
	return fmt.Sprintf("%s:update_quota", ModuleName)
}

func (m MsgUpdateQuota) ValidateBasic() sdk.Error {
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

	// Bytes should be positive
	if !m.Bytes.IsPositive() {
		return ErrorInvalidField("bytes")
	}

	return nil
}

func (m MsgUpdateQuota) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateQuota) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}
