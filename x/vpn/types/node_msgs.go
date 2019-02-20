package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgRegisterNode struct {
	From         csdkTypes.AccAddress `json:"from"`
	AmountToLock csdkTypes.Coin       `json:"amount_to_lock"`
	PricesPerGB  csdkTypes.Coins      `json:"prices_per_gb"`
	NetSpeed     sdkTypes.Bandwidth   `json:"net_speed"`
	APIPort      APIPort              `json:"api_port"`
	EncMethod    string               `json:"enc_method"`
	NodeType     string               `json:"node_type"`
	Version      string               `json:"version"`
}

func (msg MsgRegisterNode) Type() string {
	return "msg_register_node"
}

func (msg MsgRegisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.AmountToLock.Denom != "sent" || !msg.AmountToLock.IsPositive() {
		return ErrorInvalidField("amount_to_lock")
	}
	if msg.PricesPerGB == nil || msg.PricesPerGB.Len() == 0 ||
		!msg.PricesPerGB.IsValid() || !msg.PricesPerGB.IsAllPositive() {
		return ErrorInvalidField("prices_per_gb")
	}
	if !msg.NetSpeed.IsPositive() {
		return ErrorInvalidField("net_speed")
	}
	if !msg.APIPort.Valid() {
		return ErrorInvalidField("api_port")
	}
	if len(msg.EncMethod) == 0 {
		return ErrorInvalidField("enc_method")
	}
	if len(msg.NodeType) == 0 {
		return ErrorInvalidField("node_type")
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
	amountToLock csdkTypes.Coin, pricesPerGB csdkTypes.Coins,
	upload, download csdkTypes.Int, apiPort APIPort,
	encMethod, nodeType, version string) *MsgRegisterNode {

	return &MsgRegisterNode{
		From:         from,
		AmountToLock: amountToLock,
		PricesPerGB:  pricesPerGB,
		NetSpeed: sdkTypes.Bandwidth{
			Upload:   upload,
			Download: download,
		},
		APIPort:   apiPort,
		EncMethod: encMethod,
		NodeType:  nodeType,
		Version:   version,
	}
}

type MsgUpdateNodeDetails struct {
	From        csdkTypes.AccAddress `json:"from"`
	ID          NodeID               `json:"id"`
	PricesPerGB csdkTypes.Coins      `json:"prices_per_gb"`
	NetSpeed    sdkTypes.Bandwidth   `json:"net_speed"`
	APIPort     APIPort              `json:"api_port"`
	EncMethod   string               `json:"enc_method"`
	Version     string               `json:"version"`
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
	if msg.PricesPerGB != nil && (msg.PricesPerGB.Len() == 0 ||
		!msg.PricesPerGB.IsValid() || !msg.PricesPerGB.IsAllPositive()) {
		return ErrorInvalidField("prices_per_gb")
	}
	if msg.NetSpeed.IsNegative() {
		return ErrorInvalidField("net_speed")
	}
	if !msg.APIPort.Valid() {
		return ErrorInvalidField("api_port")
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

func NewMsgUpdateNodeDetails(from csdkTypes.AccAddress,
	id NodeID, pricesPerGB csdkTypes.Coins, upload, download csdkTypes.Int,
	apiPort APIPort, encMethod string, version string) *MsgUpdateNodeDetails {

	return &MsgUpdateNodeDetails{
		From:        from,
		ID:          id,
		PricesPerGB: pricesPerGB,
		NetSpeed: sdkTypes.Bandwidth{
			Upload:   upload,
			Download: download,
		},
		APIPort:   apiPort,
		EncMethod: encMethod,
		Version:   version,
	}
}

type MsgUpdateNodeStatus struct {
	From   csdkTypes.AccAddress `json:"from"`
	ID     NodeID               `json:"id"`
	Status string               `json:"status"`
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

func NewMsgUpdateNodeStatus(from csdkTypes.AccAddress,
	id NodeID, status string) *MsgUpdateNodeStatus {

	return &MsgUpdateNodeStatus{
		From:   from,
		ID:     id,
		Status: status,
	}
}

type MsgDeregisterNode struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   NodeID               `json:"id"`
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

func NewMsgDeregisterNode(from csdkTypes.AccAddress,
	id NodeID) *MsgDeregisterNode {

	return &MsgDeregisterNode{
		From: from,
		ID:   id,
	}
}
