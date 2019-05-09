package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgRegisterNode struct {
	From csdkTypes.AccAddress `json:"from"`

	Moniker          string             `json:"moniker"`
	DepositAmount    csdkTypes.Coin     `json:"deposit_amount"`
	PricesPerGB      csdkTypes.Coins    `json:"prices_per_gb"`
	InternetSpeed    sdkTypes.Bandwidth `json:"internet_speed"`
	EncryptionMethod string             `json:"encryption_method"`
	Type_            string             `json:"type_"`
	Version          string             `json:"version"`
}

func (msg MsgRegisterNode) Type() string {
	return "msg_register_node"
}

func (msg MsgRegisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}

	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.DepositAmount.Denom != "sent" || !msg.DepositAmount.IsPositive() {
		return ErrorInvalidField("deposit_amount")
	}
	if msg.PricesPerGB == nil || msg.PricesPerGB.Len() == 0 ||
		!msg.PricesPerGB.IsValid() || !msg.PricesPerGB.IsAllPositive() {
		return ErrorInvalidField("prices_per_gb")
	}
	if !msg.InternetSpeed.IsPositive() {
		return ErrorInvalidField("internet_speed")
	}
	if len(msg.EncryptionMethod) == 0 {
		return ErrorInvalidField("encryption_method")
	}
	if len(msg.Type_) == 0 {
		return ErrorInvalidField("type_")
	}
	if len(msg.Version) == 0 {
		return ErrorInvalidField("version")
	}

	return nil
}

func (msg MsgRegisterNode) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgRegisterNode) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgRegisterNode) Route() string {
	return RouterKey
}

func NewMsgRegisterNode(from csdkTypes.AccAddress,
	moniker string, depositAmount csdkTypes.Coin, pricesPerGB csdkTypes.Coins,
	internetSpeed sdkTypes.Bandwidth, encryptionMethod, type_, version string) *MsgRegisterNode {

	return &MsgRegisterNode{
		From: from,

		Moniker:          moniker,
		DepositAmount:    depositAmount,
		PricesPerGB:      pricesPerGB,
		InternetSpeed:    internetSpeed,
		EncryptionMethod: encryptionMethod,
		Type_:            type_,
		Version:          version,
	}
}

type MsgUpdateNodeDetails struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   sdkTypes.ID          `json:"id"`

	Moniker          string             `json:"moniker"`
	PricesPerGB      csdkTypes.Coins    `json:"prices_per_gb"`
	InternetSpeed    sdkTypes.Bandwidth `json:"internet_speed"`
	EncryptionMethod string             `json:"encryption_method"`
	Type_            string             `json:"type_"`
	Version          string             `json:"version"`
}

func (msg MsgUpdateNodeDetails) Type() string {
	return "msg_update_node_details"
}

func (msg MsgUpdateNodeDetails) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID.Len() == 0 || !msg.ID.Valid() {
		return ErrorInvalidField("id")
	}

	if len(msg.Moniker) > 128 {
		return ErrorInvalidField("moniker")
	}
	if msg.PricesPerGB != nil && (msg.PricesPerGB.Len() == 0 ||
		!msg.PricesPerGB.IsValid() || !msg.PricesPerGB.IsAllPositive()) {
		return ErrorInvalidField("prices_per_gb")
	}
	if msg.InternetSpeed.IsNegative() {
		return ErrorInvalidField("internet_speed")
	}

	return nil
}

func (msg MsgUpdateNodeDetails) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgUpdateNodeDetails) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateNodeDetails) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeDetails(from csdkTypes.AccAddress, id sdkTypes.ID,
	moniker string, pricesPerGB csdkTypes.Coins, internetSpeed sdkTypes.Bandwidth,
	encryptionMethod, type_, version string) *MsgUpdateNodeDetails {

	return &MsgUpdateNodeDetails{
		From: from,
		ID:   id,

		Moniker:          moniker,
		PricesPerGB:      pricesPerGB,
		InternetSpeed:    internetSpeed,
		EncryptionMethod: encryptionMethod,
		Type_:            type_,
		Version:          version,
	}
}

type MsgUpdateNodeStatus struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   sdkTypes.ID          `json:"id"`

	Status string `json:"status"`
}

func (msg MsgUpdateNodeStatus) Type() string {
	return "msg_update_node_status"
}

func (msg MsgUpdateNodeStatus) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID.Len() == 0 || !msg.ID.Valid() {
		return ErrorInvalidField("id")
	}

	if msg.Status != StatusActive && msg.Status != StatusInactive {
		return ErrorInvalidField("status")
	}

	return nil
}

func (msg MsgUpdateNodeStatus) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgUpdateNodeStatus) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateNodeStatus) Route() string {
	return RouterKey
}

func NewMsgUpdateNodeStatus(from csdkTypes.AccAddress, id sdkTypes.ID,
	status string) *MsgUpdateNodeStatus {

	return &MsgUpdateNodeStatus{
		From: from,
		ID:   id,

		Status: status,
	}
}

type MsgDeregisterNode struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   sdkTypes.ID          `json:"id"`
}

func (msg MsgDeregisterNode) Type() string {
	return "msg_deregister_node"
}

func (msg MsgDeregisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.ID.Len() == 0 || !msg.ID.Valid() {
		return ErrorInvalidField("id")
	}

	return nil
}

func (msg MsgDeregisterNode) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgDeregisterNode) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgDeregisterNode) Route() string {
	return RouterKey
}

func NewMsgDeregisterNode(from csdkTypes.AccAddress, id sdkTypes.ID) *MsgDeregisterNode {
	return &MsgDeregisterNode{
		From: from,
		ID:   id,
	}
}
