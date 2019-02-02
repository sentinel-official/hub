package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgRegisterNode struct {
	From         csdkTypes.AccAddress `json:"from"`
	AmountToLock csdkTypes.Coin       `json:"amount_to_lock"`
	APIPort      uint16               `json:"api_port"`
	NetSpeed     sdkTypes.Bandwidth   `json:"net_speed"`
	EncMethod    string               `json:"enc_method"`
	PricesPerGB  csdkTypes.Coins      `json:"prices_per_gb"`
	Version      string               `json:"version"`
	NodeType     string               `json:"node_type"`
}

func (msg MsgRegisterNode) Type() string {
	return "msg_register_node"
}

func (msg MsgRegisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorInvalidField("from")
	}
	if msg.AmountToLock.IsPositive() == false || msg.AmountToLock.Denom != "sent" {
		return errorInvalidField("amount_to_lock")
	}
	if msg.APIPort <= 0 || msg.APIPort > 65535 {
		return errorInvalidField("api_port")
	}
	if msg.NetSpeed.Download <= 0 || msg.NetSpeed.Upload <= 0 {
		return errorInvalidField("net_speed")
	}
	if len(msg.EncMethod) == 0 {
		return errorInvalidField("enc_method")
	}
	if msg.PricesPerGB == nil || msg.PricesPerGB.Len() == 0 ||
		msg.PricesPerGB.IsValid() == false || msg.PricesPerGB.IsPositive() == false {
		return errorInvalidField("prices_per_gb")
	}
	if len(msg.Version) == 0 {
		return errorInvalidField("version")
	}
	if len(msg.NodeType) == 0 {
		return errorInvalidField("node_type")
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
	apiPort uint16, upload, download uint64,
	encMethod string, pricesPerGB csdkTypes.Coins,
	version, nodeType string,
	amountToLock csdkTypes.Coin) *MsgRegisterNode {

	return &MsgRegisterNode{
		From:         from,
		AmountToLock: amountToLock,
		APIPort:      apiPort,
		NetSpeed: sdkTypes.Bandwidth{
			Upload:   upload,
			Download: download,
		},
		EncMethod:   encMethod,
		PricesPerGB: pricesPerGB,
		Version:     version,
		NodeType:    nodeType,
	}
}

type MsgUpdateNodeDetails struct {
	From        csdkTypes.AccAddress `json:"from"`
	ID          string               `json:"id"`
	APIPort     uint16               `json:"api_port"`
	NetSpeed    sdkTypes.Bandwidth   `json:"net_speed"`
	EncMethod   string               `json:"enc_method"`
	PricesPerGB csdkTypes.Coins      `json:"prices_per_gb"`
	Version     string               `json:"version"`
}

func (msg MsgUpdateNodeDetails) Type() string {
	return "msg_update_node_details"
}

func (msg MsgUpdateNodeDetails) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorInvalidField("from")
	}
	if len(msg.ID) == 0 {
		return errorInvalidField("id")
	}
	if msg.APIPort < 0 || msg.APIPort > 65535 {
		return errorInvalidField("api_port")
	}
	if msg.NetSpeed.Download < 0 || msg.NetSpeed.Upload < 0 {
		return errorInvalidField("net_speed")
	}
	if (msg.PricesPerGB != nil && msg.PricesPerGB.Len() != 0) &&
		(msg.PricesPerGB.IsValid() == false || msg.PricesPerGB.IsPositive() == false) {
		return errorInvalidField("prices_per_gb")
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
	id string, apiPort uint16, upload, download uint64,
	encMethod string, pricesPerGB csdkTypes.Coins, version string) *MsgUpdateNodeDetails {

	return &MsgUpdateNodeDetails{
		From:    from,
		ID:      id,
		APIPort: apiPort,
		NetSpeed: sdkTypes.Bandwidth{
			Upload:   upload,
			Download: download,
		},
		EncMethod:   encMethod,
		PricesPerGB: pricesPerGB,
		Version:     version,
	}
}

type MsgDeregisterNode struct {
	From csdkTypes.AccAddress `json:"from"`
	ID   string               `json:"id"`
}

func (msg MsgDeregisterNode) Type() string {
	return "msg_deregister_node"
}

func (msg MsgDeregisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorInvalidField("from")
	}
	if len(msg.ID) == 0 {
		return errorInvalidField("id")
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

func NewMsgDeregisterNode(from csdkTypes.AccAddress, id string) *MsgDeregisterNode {
	return &MsgDeregisterNode{
		From: from,
		ID:   id,
	}
}

type MsgUpdateNodeStatus struct {
	From   csdkTypes.AccAddress `json:"from"`
	ID     string               `json:"id"`
	Status string               `json:"status"`
}

func (msg MsgUpdateNodeStatus) Type() string {
	return "msg_update_node_status"
}

func (msg MsgUpdateNodeStatus) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorInvalidField("from")
	}
	if len(msg.ID) == 0 {
		return errorInvalidField("id")
	}
	if msg.Status != StatusActive && msg.Status != StatusInactive {
		return errorInvalidField("status")
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

func NewMsgUpdateNodeStatus(from csdkTypes.AccAddress, id, status string) *MsgUpdateNodeStatus {
	return &MsgUpdateNodeStatus{
		From:   from,
		ID:     id,
		Status: status,
	}
}

type MsgAddSession struct {
	From         csdkTypes.AccAddress `json:"from"`
	NodeID       string               `json:"node_id"`
	AmountToLock csdkTypes.Coin       `json:"amount_to_lock"`
}

func (msg MsgAddSession) Type() string {
	return "msg_add_session"
}

func (msg MsgAddSession) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorInvalidField("from")
	}
	if len(msg.NodeID) == 0 {
		return errorInvalidField("node_id")
	}
	if msg.AmountToLock.IsPositive() == false {
		return errorInvalidField("amount_to_lock")
	}

	return nil
}

func (msg MsgAddSession) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgAddSession) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgAddSession) Route() string {
	return RouterKey
}

func NewMsgAddSession(from csdkTypes.AccAddress,
	nodeID string, amountToLock csdkTypes.Coin) *MsgAddSession {
	return &MsgAddSession{
		From:         from,
		NodeID:       nodeID,
		AmountToLock: amountToLock,
	}
}

type MsgUpdateSessionBandwidth struct {
	From              csdkTypes.AccAddress `json:"from"`
	SessionID         string               `json:"session_id"`
	ConsumedBandwidth sdkTypes.Bandwidth   `json:"consumed_bandwidth"`
	ClientSign        []byte               `json:"client_sign"`
	NodeOwnerSign     []byte               `json:"node_owner_sign"`
}

func (msg MsgUpdateSessionBandwidth) Type() string {
	return "msg_update_session_bandwidth"
}

func (msg MsgUpdateSessionBandwidth) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorInvalidField("from")
	}
	if len(msg.SessionID) == 0 {
		return errorInvalidField("session_id")
	}
	if msg.ConsumedBandwidth.Upload == 0 || msg.ConsumedBandwidth.Download == 0 {
		return errorInvalidField("consumed_bandwidth")
	}
	if len(msg.ClientSign) == 0 {
		return errorInvalidField("client_sign")
	}
	if len(msg.NodeOwnerSign) == 0 {
		return errorInvalidField("node_owner_sign")
	}

	return nil
}

func (msg MsgUpdateSessionBandwidth) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgUpdateSessionBandwidth) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateSessionBandwidth) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionBandwidth(from csdkTypes.AccAddress,
	sessionID string, consumedUpload, consumedDownload uint64,
	clientSign []byte, nodeOwnerSign []byte) *MsgUpdateSessionBandwidth {
	return &MsgUpdateSessionBandwidth{
		From:      from,
		SessionID: sessionID,
		ConsumedBandwidth: sdkTypes.Bandwidth{
			Upload:   consumedUpload,
			Download: consumedDownload,
		},
		ClientSign:    clientSign,
		NodeOwnerSign: nodeOwnerSign,
	}
}
