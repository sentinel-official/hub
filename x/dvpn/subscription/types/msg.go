package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgStartSubscription)(nil)
	_ sdk.Msg = (*MsgAddAddressForSubscription)(nil)
	_ sdk.Msg = (*MsgRemoveAddressForSubscription)(nil)
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

func NewMsgAddAddressForSubscription(from sdk.AccAddress, id uint64, address sdk.AccAddress) MsgAddAddressForSubscription {
	return MsgAddAddressForSubscription{
		From:    from,
		ID:      id,
		Address: address,
	}
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

// MsgRemoveAddressForSubscription is for removing an address for subscription.
type MsgRemoveAddressForSubscription struct {
	From    sdk.AccAddress `json:"from"`
	ID      uint64         `json:"id"`
	Address sdk.AccAddress `json:"address"`
}

func NewMsgRemoveAddressForSubscription(from sdk.AccAddress, id uint64, address sdk.AccAddress) MsgRemoveAddressForSubscription {
	return MsgRemoveAddressForSubscription{
		From:    from,
		ID:      id,
		Address: address,
	}
}

func (m MsgRemoveAddressForSubscription) Route() string {
	return RouterKey
}

func (m MsgRemoveAddressForSubscription) Type() string {
	return "remove_address_for_subscription"
}

func (m MsgRemoveAddressForSubscription) ValidateBasic() sdk.Error {
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

func (m MsgRemoveAddressForSubscription) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRemoveAddressForSubscription) GetSigners() []sdk.AccAddress {
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
