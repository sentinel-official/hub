package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var _ sdk.Msg = (*MsgRegisterNode)(nil)

type MsgRegisterNode struct {
	From          sdk.AccAddress `json:"from"`
	T             string         `json:"type"`
	Version       string         `json:"version"`
	Moniker       string         `json:"moniker"`
	PricesPerGB   sdk.Coins      `json:"prices_per_gb"`
	InternetSpeed hub.Bandwidth  `json:"internet_speed"`
	Encryption    string         `json:"encryption"`
}

func (msg MsgRegisterNode) Type() string {
	return "register_node"
}

func (msg MsgRegisterNode) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.T == "" {
		return ErrorInvalidField("type")
	}
	if msg.Version == "" {
		return ErrorInvalidField("version")
	}
	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB == nil ||
		msg.PricesPerGB.Len() == 0 || !isValidCoins(msg.PricesPerGB) {
		return ErrorInvalidField("prices_per_gb")
	}
	if !msg.InternetSpeed.AllPositive() {
		return ErrorInvalidField("internet_speed")
	}
	if msg.Encryption == "" {
		return ErrorInvalidField("encryption")
	}

	return nil
}

func (msg MsgRegisterNode) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgRegisterNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgRegisterNode) Route() string {
	return RouterKey
}

func NewMsgRegisterNode(from sdk.AccAddress,
	t, version, moniker string, pricesPerGB sdk.Coins,
	internetSpeed hub.Bandwidth, encryption string) *MsgRegisterNode {
	return &MsgRegisterNode{
		From:          from,
		T:             t,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ sdk.Msg = (*MsgUpdateNodeInfo)(nil)

type MsgUpdateNodeInfo struct {
	From          sdk.AccAddress `json:"from"`
	ID            hub.NodeID     `json:"id"`
	T             string         `json:"type"`
	Version       string         `json:"version"`
	Moniker       string         `json:"moniker"`
	PricesPerGB   sdk.Coins      `json:"prices_per_gb"`
	InternetSpeed hub.Bandwidth  `json:"internet_speed"`
	Encryption    string         `json:"encryption"`
}

func (msg MsgUpdateNodeInfo) Type() string {
	return "update_node_info"
}

func (msg MsgUpdateNodeInfo) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB != nil &&
		(msg.PricesPerGB.Len() == 0 || !msg.PricesPerGB.IsValid()) {
		return ErrorInvalidField("prices_per_gb")
	}
	if msg.InternetSpeed.AnyNegative() {
		return ErrorInvalidField("internet_speed")
	}

	return nil
}

func (msg MsgUpdateNodeInfo) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateNodeInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgUpdateNodeInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeInfo(from sdk.AccAddress, id hub.NodeID,
	t, version, moniker string, pricesPerGB sdk.Coins,
	internetSpeed hub.Bandwidth, encryption string) *MsgUpdateNodeInfo {
	return &MsgUpdateNodeInfo{
		From:          from,
		ID:            id,
		T:             t,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ sdk.Msg = (*MsgDeregisterNode)(nil)

type MsgDeregisterNode struct {
	From sdk.AccAddress `json:"from"`
	ID   hub.NodeID     `json:"id"`
}

func (msg MsgDeregisterNode) Type() string {
	return "deregister_node"
}

func (msg MsgDeregisterNode) ValidateBasic() sdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	return nil
}

func (msg MsgDeregisterNode) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgDeregisterNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgDeregisterNode) Route() string {
	return RouterKey
}

func NewMsgDeregisterNode(from sdk.AccAddress, id hub.NodeID) *MsgDeregisterNode {
	return &MsgDeregisterNode{
		From: from,
		ID:   id,
	}
}
