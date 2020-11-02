package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegister)(nil)
	_ sdk.Msg = (*MsgUpdate)(nil)
	_ sdk.Msg = (*MsgSetStatus)(nil)
)

// MsgRegister is for registering a VPN node.
type MsgRegister struct {
	From      sdk.AccAddress  `json:"from"`
	Provider  hub.ProvAddress `json:"provider,omitempty"`
	Price     sdk.Coins       `json:"price,omitempty"`
	RemoteURL string          `json:"remote_url"`
}

func NewMsgRegister(from sdk.AccAddress, provider hub.ProvAddress, price sdk.Coins, remoteURL string) MsgRegister {
	return MsgRegister{
		From:      from,
		Provider:  provider,
		Price:     price,
		RemoteURL: remoteURL,
	}
}

func (m MsgRegister) Route() string {
	return RouterKey
}

func (m MsgRegister) Type() string {
	return fmt.Sprintf("%s:register", ModuleName)
}

func (m MsgRegister) ValidateBasic() error {
	if m.From == nil || m.From.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Either provider or price should be nil
	if (m.Provider != nil && m.Price != nil) ||
		(m.Provider == nil && m.Price == nil) {
		return errors.Wrapf(ErrorInvalidField, "%s", "provider and price")
	}

	// Provider can be nil. If not, it shouldn't be empty
	if m.Provider != nil && m.Provider.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "provider")
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "price")
	}

	// RemoteURL can't be empty and length should be (0, 64]
	if len(m.RemoteURL) == 0 || len(m.RemoteURL) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "remote_url")
	}

	return nil
}

func (m MsgRegister) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegister) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgUpdate is for updating the information of a VPN node.
type MsgUpdate struct {
	From      hub.NodeAddress `json:"from"`
	Provider  hub.ProvAddress `json:"provider,omitempty"`
	Price     sdk.Coins       `json:"price,omitempty"`
	RemoteURL string          `json:"remote_url,omitempty"`
}

func NewMsgUpdate(from hub.NodeAddress, provider hub.ProvAddress, price sdk.Coins, remoteURL string) MsgUpdate {
	return MsgUpdate{
		From:      from,
		Provider:  provider,
		Price:     price,
		RemoteURL: remoteURL,
	}
}

func (m MsgUpdate) Route() string {
	return RouterKey
}

func (m MsgUpdate) Type() string {
	return fmt.Sprintf("%s:update", ModuleName)
}

func (m MsgUpdate) ValidateBasic() error {
	if m.From == nil || m.From.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Provider and Price both shouldn't nil at the same time
	if m.Provider != nil && m.Price != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "provider and price")
	}

	// Provider can be nil. If not, it shouldn't be empty
	if m.Provider != nil && m.Provider.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "provider")
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "price")
	}

	// RemoteURL length should be [0, 64]
	if len(m.RemoteURL) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "remote_url")
	}

	return nil
}

func (m MsgUpdate) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

// MsgSetStatus is for updating the status of a VPN node.
type MsgSetStatus struct {
	From   hub.NodeAddress `json:"from"`
	Status hub.Status      `json:"status"`
}

func NewMsgSetStatus(from hub.NodeAddress, status hub.Status) MsgSetStatus {
	return MsgSetStatus{
		From:   from,
		Status: status,
	}
}

func (m MsgSetStatus) Route() string {
	return RouterKey
}

func (m MsgSetStatus) Type() string {
	return fmt.Sprintf("%s:set_status", ModuleName)
}

func (m MsgSetStatus) ValidateBasic() error {
	if m.From == nil || m.From.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Status should be valid
	if !m.Status.Equal(hub.StatusActive) && !m.Status.Equal(hub.StatusInactive) {
		return errors.Wrapf(ErrorInvalidField, "%s", "status")
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
