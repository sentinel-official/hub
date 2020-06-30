package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgStartSubscription)(nil)
	_ sdk.Msg = (*MsgAddMemberForSubscription)(nil)
	_ sdk.Msg = (*MsgRemoveMemberForSubscription)(nil)
	_ sdk.Msg = (*MsgEndSubscription)(nil)
)

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

	if m.ID == 0 {
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

	// Denom length should be [3, 16]
	if len(m.Denom) < 3 || len(m.Denom) > 16 {
		return ErrorInvalidField("denom")
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

// MsgAddMemberForSubscription is for adding a member for a subscription.
type MsgAddMemberForSubscription struct {
	From    sdk.AccAddress `json:"from"`
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
}

func NewMsgAddMemberForSubscription(from sdk.AccAddress, id uint64, address sdk.AccAddress) MsgAddMemberForSubscription {
	return MsgAddMemberForSubscription{
		From:    from,
		ID:      id,
		Address: address,
	}
}

func (m MsgAddMemberForSubscription) Route() string {
	return RouterKey
}

func (m MsgAddMemberForSubscription) Type() string {
	return "add_member_for_subscription"
}

func (m MsgAddMemberForSubscription) ValidateBasic() sdk.Error {
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

	// From and Address both shouldn't be same
	if m.From.Equals(m.Address) {
		return ErrorInvalidField("from and address")
	}

	return nil
}

func (m MsgAddMemberForSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddMemberForSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgRemoveMemberForSubscription is for removing a member for a subscription.
type MsgRemoveMemberForSubscription struct {
	From    sdk.AccAddress `json:"from"`
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
}

func NewMsgRemoveMemberForSubscription(from sdk.AccAddress, id uint64, address sdk.AccAddress) MsgRemoveMemberForSubscription {
	return MsgRemoveMemberForSubscription{
		From:    from,
		ID:      id,
		Address: address,
	}
}

func (m MsgRemoveMemberForSubscription) Route() string {
	return RouterKey
}

func (m MsgRemoveMemberForSubscription) Type() string {
	return "remove_address_for_subscription"
}

func (m MsgRemoveMemberForSubscription) ValidateBasic() sdk.Error {
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

	// From and Address both shouldn't be same
	if m.From.Equals(m.Address) {
		return ErrorInvalidField("from and address")
	}

	return nil
}

func (m MsgRemoveMemberForSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRemoveMemberForSubscription) GetSigners() []sdk.AccAddress {
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
