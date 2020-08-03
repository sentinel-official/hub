package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgStartSubscription)(nil)
	_ sdk.Msg = (*MsgAddQuotaForSubscription)(nil)
	_ sdk.Msg = (*MsgUpdateQuotaForSubscription)(nil)
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

// MsgAddQuotaForSubscription is for adding the bandwidth quota for an address.
type MsgAddQuotaForSubscription struct {
	From      sdk.AccAddress `json:"from"`
	ID        uint64         `json:"id"`
	Address   sdk.AccAddress `json:"address"`
	Bandwidth hub.Bandwidth  `json:"bandwidth"`
}

func NewMsgAddQuotaForSubscription(from sdk.AccAddress, id uint64,
	address sdk.AccAddress, bandwidth hub.Bandwidth) MsgAddQuotaForSubscription {
	return MsgAddQuotaForSubscription{
		From:      from,
		ID:        id,
		Address:   address,
		Bandwidth: bandwidth,
	}
}

func (m MsgAddQuotaForSubscription) Route() string {
	return RouterKey
}

func (m MsgAddQuotaForSubscription) Type() string {
	return "add_quota_for_subscription"
}

func (m MsgAddQuotaForSubscription) ValidateBasic() sdk.Error {
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

	// Bandwidth should be valid
	if !m.Bandwidth.IsValid() {
		return ErrorInvalidField("bandwidth")
	}

	return nil
}

func (m MsgAddQuotaForSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddQuotaForSubscription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgUpdateQuotaForSubscription is for updating the bandwidth quota for an address.
type MsgUpdateQuotaForSubscription struct {
	From      sdk.AccAddress `json:"from"`
	ID        uint64         `json:"id"`
	Address   sdk.AccAddress `json:"address"`
	Bandwidth hub.Bandwidth  `json:"bandwidth"`
}

func NewMsgUpdateQuotaForSubscription(from sdk.AccAddress, id uint64,
	address sdk.AccAddress, bandwidth hub.Bandwidth) MsgUpdateQuotaForSubscription {
	return MsgUpdateQuotaForSubscription{
		From:      from,
		ID:        id,
		Address:   address,
		Bandwidth: bandwidth,
	}
}

func (m MsgUpdateQuotaForSubscription) Route() string {
	return RouterKey
}

func (m MsgUpdateQuotaForSubscription) Type() string {
	return "update_quota_for_subscription"
}

func (m MsgUpdateQuotaForSubscription) ValidateBasic() sdk.Error {
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

	// Bandwidth should be valid
	if !m.Bandwidth.IsValid() {
		return ErrorInvalidField("bandwidth")
	}

	return nil
}

func (m MsgUpdateQuotaForSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateQuotaForSubscription) GetSigners() []sdk.AccAddress {
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
