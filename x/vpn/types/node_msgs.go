package types

import (
	"encoding/json"

	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

var _ csdk.Msg = (*MsgRegisterNode)(nil)

type MsgRegisterNode struct {
	From          csdk.AccAddress `json:"from"`
	Type_         string          `json:"type"` // nolint:golint
	Version       string          `json:"version"`
	Moniker       string          `json:"moniker"`
	PricesPerGB   csdk.Coins      `json:"prices_per_gb"`
	InternetSpeed sdk.Bandwidth   `json:"internet_speed"`
	Encryption    string          `json:"encryption"`
}

func (msg MsgRegisterNode) Type() string {
	return "MsgRegisterNode"
}

// nolint: gocyclo
func (msg MsgRegisterNode) ValidateBasic() csdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.Type_ == "" {
		return ErrorInvalidField("type")
	}
	if msg.Version == "" {
		return ErrorInvalidField("version")
	}
	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB == nil ||
		msg.PricesPerGB.Len() == 0 || !msg.PricesPerGB.IsValid() {

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

func (msg MsgRegisterNode) GetSigners() []csdk.AccAddress {
	return []csdk.AccAddress{msg.From}
}

func (msg MsgRegisterNode) Route() string {
	return RouterKey
}

func NewMsgRegisterNode(from csdk.AccAddress,
	_type, version, moniker string, pricesPerGB csdk.Coins,
	internetSpeed sdk.Bandwidth, encryption string) *MsgRegisterNode {

	return &MsgRegisterNode{
		From:          from,
		Type_:         _type,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ csdk.Msg = (*MsgUpdateNodeInfo)(nil)

type MsgUpdateNodeInfo struct {
	From          csdk.AccAddress `json:"from"`
	ID            sdk.ID          `json:"id"`
	Type_         string          `json:"type"` // nolint:golint
	Version       string          `json:"version"`
	Moniker       string          `json:"moniker"`
	PricesPerGB   csdk.Coins      `json:"prices_per_gb"`
	InternetSpeed sdk.Bandwidth   `json:"internet_speed"`
	Encryption    string          `json:"encryption"`
}

func (msg MsgUpdateNodeInfo) Type() string {
	return "MsgUpdateNodeInfo"
}

func (msg MsgUpdateNodeInfo) ValidateBasic() csdk.Error {
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

func (msg MsgUpdateNodeInfo) GetSigners() []csdk.AccAddress {
	return []csdk.AccAddress{msg.From}
}

func (msg MsgUpdateNodeInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeInfo(from csdk.AccAddress, id sdk.ID,
	_type, version, moniker string, pricesPerGB csdk.Coins,
	internetSpeed sdk.Bandwidth, encryption string) *MsgUpdateNodeInfo {

	return &MsgUpdateNodeInfo{
		From:          from,
		ID:            id,
		Type_:         _type,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ csdk.Msg = (*MsgUpdateNodeStatus)(nil)

type MsgUpdateNodeStatus struct {
	From   csdk.AccAddress `json:"from"`
	ID     sdk.ID          `json:"id"`
	Status string          `json:"status"`
}

func (msg MsgUpdateNodeStatus) Type() string {
	return "MsgUpdateNodeStatus"
}

func (msg MsgUpdateNodeStatus) ValidateBasic() csdk.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.Status != StatusActive && msg.Status != StatusInactive {
		return ErrorInvalidField("status")
	}

	return nil
}

func (msg MsgUpdateNodeStatus) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateNodeStatus) GetSigners() []csdk.AccAddress {
	return []csdk.AccAddress{msg.From}
}

func (msg MsgUpdateNodeStatus) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeStatus(from csdk.AccAddress, id sdk.ID, status string) *MsgUpdateNodeStatus {

	return &MsgUpdateNodeStatus{
		From:   from,
		ID:     id,
		Status: status,
	}
}

var _ csdk.Msg = (*MsgDeregisterNode)(nil)

type MsgDeregisterNode struct {
	From csdk.AccAddress `json:"from"`
	ID   sdk.ID          `json:"id"`
}

func (msg MsgDeregisterNode) Type() string {
	return "MsgDeregisterNode"
}

func (msg MsgDeregisterNode) ValidateBasic() csdk.Error {
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

func (msg MsgDeregisterNode) GetSigners() []csdk.AccAddress {
	return []csdk.AccAddress{msg.From}
}

func (msg MsgDeregisterNode) Route() string {
	return RouterKey
}

func NewMsgDeregisterNode(from csdk.AccAddress, id sdk.ID) *MsgDeregisterNode {
	return &MsgDeregisterNode{
		From: from,
		ID:   id,
	}
}
