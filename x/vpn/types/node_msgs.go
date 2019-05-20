package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var _ csdkTypes.Msg = (*MsgRegisterNode)(nil)

type MsgRegisterNode struct {
	From          csdkTypes.AccAddress `json:"from"`
	Type_         string               `json:"type_"`
	Version       string               `json:"version"`
	Moniker       string               `json:"moniker"`
	PricesPerGB   csdkTypes.Coins      `json:"prices_per_gb"`
	InternetSpeed sdkTypes.Bandwidth   `json:"internet_speed"`
	Encryption    string               `json:"encryption"`
}

func (msg MsgRegisterNode) Type() string {
	return "MsgRegisterNode"
}

func (msg MsgRegisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if len(msg.Type_) == 0 {
		return ErrorInvalidField("type_")
	}
	if len(msg.Version) == 0 {
		return ErrorInvalidField("version")
	}
	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB == nil ||
		msg.PricesPerGB.Len() == 0 || !msg.PricesPerGB.IsValid() {

		return ErrorInvalidField("prices_per_gb")
	}
	if !msg.InternetSpeed.IsPositive() {
		return ErrorInvalidField("internet_speed")
	}
	if len(msg.Encryption) == 0 {
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

func (msg MsgRegisterNode) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgRegisterNode) Route() string {
	return RouterKey
}

func NewMsgRegisterNode(from csdkTypes.AccAddress,
	type_, version, moniker string, pricesPerGB csdkTypes.Coins,
	internetSpeed sdkTypes.Bandwidth, encryption string) MsgRegisterNode {

	return MsgRegisterNode{
		From:          from,
		Type_:         type_,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ csdkTypes.Msg = (*MsgUpdateNodeInfo)(nil)

type MsgUpdateNodeInfo struct {
	From          csdkTypes.AccAddress `json:"from"`
	ID            sdkTypes.ID          `json:"id"`
	Type_         string               `json:"type_"`
	Version       string               `json:"version"`
	Moniker       string               `json:"moniker"`
	PricesPerGB   csdkTypes.Coins      `json:"prices_per_gb"`
	InternetSpeed sdkTypes.Bandwidth   `json:"internet_speed"`
	Encryption    string               `json:"encryption"`
}

func (msg MsgUpdateNodeInfo) Type() string {
	return "MsgUpdateNodeInfo"
}

func (msg MsgUpdateNodeInfo) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID == nil || msg.ID.Len() == 0 {
		return ErrorInvalidField("id")
	}
	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB != nil &&
		(msg.PricesPerGB.Len() == 0 || !msg.PricesPerGB.IsValid()) {

		return ErrorInvalidField("prices_per_gb")
	}
	if msg.InternetSpeed.IsNegative() {
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

func (msg MsgUpdateNodeInfo) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateNodeInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeInfo(from csdkTypes.AccAddress, id sdkTypes.ID,
	type_, version, moniker string, pricesPerGB csdkTypes.Coins,
	internetSpeed sdkTypes.Bandwidth, encryption string) MsgUpdateNodeInfo {

	return MsgUpdateNodeInfo{
		From:          from,
		ID:            id,
		Type_:         type_,
		Version:       version,
		Moniker:       moniker,
		PricesPerGB:   pricesPerGB,
		InternetSpeed: internetSpeed,
		Encryption:    encryption,
	}
}

var _ csdkTypes.Msg = (*MsgUpdateNodeStatus)(nil)

type MsgUpdateNodeStatus struct {
	From   csdkTypes.AccAddress `json:"from"`
	ID     sdkTypes.ID          `json:"id"`
	Status string               `json:"status"`
}

func (msg MsgUpdateNodeStatus) Type() string {
	return "MsgUpdateNodeStatus"
}

func (msg MsgUpdateNodeStatus) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID == nil || msg.ID.Len() == 0 {
		return ErrorInvalidField("id")
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

func (msg MsgUpdateNodeStatus) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateNodeStatus) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeStatus(from csdkTypes.AccAddress, id sdkTypes.ID,
	status string) MsgUpdateNodeStatus {

	return MsgUpdateNodeStatus{
		From:   from,
		ID:     id,
		Status: status,
	}
}

var _ csdkTypes.Msg = (*MsgDeregisterNode)(nil)

type MsgDeregisterNode struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   sdkTypes.ID          `json:"id"`
}

func (msg MsgDeregisterNode) Type() string {
	return "MsgDeregisterNode"
}

func (msg MsgDeregisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID == nil || msg.ID.Len() == 0 {
		return ErrorInvalidField("id")
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

func (msg MsgDeregisterNode) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgDeregisterNode) Route() string {
	return RouterKey
}

func NewMsgDeregisterNode(from csdkTypes.AccAddress, id sdkTypes.ID) MsgDeregisterNode {
	return MsgDeregisterNode{
		From: from,
		ID:   id,
	}
}
