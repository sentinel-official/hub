package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegisterNode)(nil)
	_ sdk.Msg = (*MsgUpdateNode)(nil)
	_ sdk.Msg = (*MsgSetNodeStatus)(nil)
)

// MsgRegisterNode is for registering a dVPN node
type MsgRegisterNode struct {
	From          sdk.AccAddress  `json:"from"`
	Provider      hub.ProvAddress `json:"provider"`
	Price         sdk.Coins       `json:"price"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed"`
	RemoteURL     string          `json:"remote_url"`
	Version       string          `json:"version"`
	Category      NodeCategory    `json:"category"`
}

func NewMsgRegisterNode(from sdk.AccAddress, provider hub.ProvAddress, price sdk.Coins,
	speed hub.Bandwidth, remoteURL, version string, category NodeCategory) MsgRegisterNode {
	return MsgRegisterNode{
		From:          from,
		Provider:      provider,
		Price:         price,
		InternetSpeed: speed,
		RemoteURL:     remoteURL,
		Version:       version,
		Category:      category,
	}
}

func (m MsgRegisterNode) Route() string {
	return RouterKey
}

func (m MsgRegisterNode) Type() string {
	return "register_node"
}

func (m MsgRegisterNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Provider can be nil. If not, it shouldn't be empty
	if m.Provider != nil && m.Provider.Empty() {
		return ErrorInvalidField("provider")
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return ErrorInvalidField("price")
	}

	// InternetSpeed can't be zero and should be positive
	if m.InternetSpeed.IsAnyZero() {
		return ErrorInvalidField("internet_speed")
	}

	// RemoteURL can't be empty and length should be (0, 32]
	if len(m.RemoteURL) == 0 || len(m.RemoteURL) > 32 {
		return ErrorInvalidField("remote_url")
	}

	// Version can't be empty and length should be (0, 32]
	if len(m.Version) == 0 || len(m.Version) > 32 {
		return ErrorInvalidField("version")
	}

	// Category can't be invalid
	if !m.Category.IsValid() {
		return ErrorInvalidField("category")
	}

	return nil
}

func (m MsgRegisterNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegisterNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgUpdateNode is for updating the information of a dVPN node
type MsgUpdateNode struct {
	From          hub.NodeAddress `json:"from"`
	Provider      hub.ProvAddress `json:"provider"`
	Price         sdk.Coins       `json:"price"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed"`
	RemoteURL     string          `json:"remote_url"`
	Version       string          `json:"version"`
	Category      NodeCategory    `json:"category"`
}

func NewMsgUpdateNode(from hub.NodeAddress, provider hub.ProvAddress, price sdk.Coins,
	speed hub.Bandwidth, remoteURL, version string, category NodeCategory) MsgUpdateNode {
	return MsgUpdateNode{
		From:          from,
		Provider:      provider,
		Price:         price,
		InternetSpeed: speed,
		RemoteURL:     remoteURL,
		Version:       version,
		Category:      category,
	}
}

func (m MsgUpdateNode) Route() string {
	return RouterKey
}

func (m MsgUpdateNode) Type() string {
	return "update_node"
}

func (m MsgUpdateNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Provider can be nil. If not, it shouldn't be empty
	if m.Provider != nil && m.Provider.Empty() {
		return ErrorInvalidField("provider")
	}

	// Price can be nil. If not, it should be empty or valid
	if m.Price != nil && !hub.AreEmptyCoins(m.Price) && !m.Price.IsValid() {
		return ErrorInvalidField("price")
	}

	// InternetSpeed can be zero. If not, it should be positive
	if !m.InternetSpeed.IsAllZero() && m.InternetSpeed.IsAnyZero() {
		return ErrorInvalidField("internet_speed")
	}

	// RemoteURL can be empty. If not, length should be (0, 32]
	if len(m.RemoteURL) != 0 && len(m.RemoteURL) > 32 {
		return ErrorInvalidField("remote_url")
	}

	// Version can be empty. If not, length should be (0, 32]
	if len(m.Version) != 0 && len(m.Version) > 32 {
		return ErrorInvalidField("version")
	}

	// Category can be Unknown. If not, should be valid
	if m.Category != CategoryUnknown && !m.Category.IsValid() {
		return ErrorInvalidField("category")
	}

	return nil
}

func (m MsgUpdateNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

// MsgSetNodeStatus is for updating the status of a dVPN node
type MsgSetNodeStatus struct {
	From   hub.NodeAddress `json:"from"`
	Status hub.Status      `json:"status"`
}

func NewMsgSetNodeStatus(from hub.NodeAddress, status hub.Status) MsgSetNodeStatus {
	return MsgSetNodeStatus{
		From:   from,
		Status: status,
	}
}

func (m MsgSetNodeStatus) Route() string {
	return RouterKey
}

func (m MsgSetNodeStatus) Type() string {
	return "set_node_status"
}

func (m MsgSetNodeStatus) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if !m.Status.IsValid() {
		return ErrorInvalidField("status")
	}

	return nil
}

func (m MsgSetNodeStatus) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSetNodeStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
