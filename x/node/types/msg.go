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
	From          sdk.AccAddress  `json:"from"`
	Moniker       string          `json:"moniker"`
	Provider      hub.ProvAddress `json:"provider,omitempty"`
	Price         sdk.Coins       `json:"price,omitempty"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed"`
	RemoteURL     string          `json:"remote_url"`
	Version       string          `json:"version"`
	Category      Category        `json:"category"`
}

func NewMsgRegister(from sdk.AccAddress, moniker string, provider hub.ProvAddress, price sdk.Coins,
	speed hub.Bandwidth, remoteURL, version string, category Category) MsgRegister {
	return MsgRegister{
		From:          from,
		Moniker:       moniker,
		Provider:      provider,
		Price:         price,
		InternetSpeed: speed,
		RemoteURL:     remoteURL,
		Version:       version,
		Category:      category,
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

	// Moniker can't be empty and length should be (0, 64]
	if len(m.Moniker) == 0 || len(m.Moniker) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "moniker")
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

	// InternetSpeed shouldn't be negative and zero
	if !m.InternetSpeed.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "internet_speed")
	}

	// RemoteURL can't be empty and length should be (0, 64]
	if len(m.RemoteURL) == 0 || len(m.RemoteURL) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "remote_url")
	}

	// Version can't be empty and length should be (0, 64]
	if len(m.Version) == 0 || len(m.Version) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "version")
	}

	// Category should be valid
	if !m.Category.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "category")
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
	From          hub.NodeAddress `json:"from"`
	Moniker       string          `json:"moniker"`
	Provider      hub.ProvAddress `json:"provider,omitempty"`
	Price         sdk.Coins       `json:"price,omitempty"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed,omitempty"`
	RemoteURL     string          `json:"remote_url,omitempty"`
	Version       string          `json:"version,omitempty"`
	Category      Category        `json:"category,omitempty"`
}

func NewMsgUpdate(from hub.NodeAddress, moniker string, provider hub.ProvAddress, price sdk.Coins,
	speed hub.Bandwidth, remoteURL, version string, category Category) MsgUpdate {
	return MsgUpdate{
		From:          from,
		Moniker:       moniker,
		Provider:      provider,
		Price:         price,
		InternetSpeed: speed,
		RemoteURL:     remoteURL,
		Version:       version,
		Category:      category,
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

	// Moniker length should be [0, 64]
	if len(m.Moniker) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "moniker")
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

	// InternetSpeed can be zero. If not, it shouldn't be negative and zero
	if !m.InternetSpeed.IsAllZero() && !m.InternetSpeed.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "internet_speed")
	}

	// RemoteURL length should be [0, 64]
	if len(m.RemoteURL) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "remote_url")
	}

	// Version length should be [0, 64]
	if len(m.Version) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "version")
	}

	// Category can be Unknown. If not, should be valid
	if !m.Category.Equal(CategoryUnknown) && !m.Category.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "category")
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
