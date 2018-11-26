package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgRegisterNode struct {
	From csdkTypes.AccAddress `json:"from"`

	APIPort    int64             `json:"api_port"`
	Location   sdkTypes.Location `json:"location"`
	NetSpeed   sdkTypes.NetSpeed `json:"net_speed"`
	EncMethod  string            `json:"enc_method"`
	PricePerGB int64             `json:"price_per_gb"`
	Version    string            `json:"version"`

	LockerID  string          `json:"locker_id"`
	Coins     csdkTypes.Coins `json:"coins"`
	PubKey    crypto.PubKey   `json:"pub_key"`
	Signature []byte          `json:"signature"`
}

func (msg MsgRegisterNode) Type() string {
	return "msg_register_node"
}

func (msg MsgRegisterNode) ValidateBasic() csdkTypes.Error {
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
	return "vpn"
}

func NewMsgRegisterNode(from csdkTypes.AccAddress, apiPort int64,
	latitude int64, longitude int64, city string, country string,
	upload int64, download int64,
	encMethod string, pricePerGB int64, version string,
	lockerID string, coins csdkTypes.Coins, pubKey crypto.PubKey, signature []byte) *MsgRegisterNode {

	return &MsgRegisterNode{
		From:    from,
		APIPort: apiPort,
		Location: sdkTypes.Location{
			Latitude:  latitude,
			Longitude: longitude,
			City:      city,
			Country:   country,
		},
		NetSpeed: sdkTypes.NetSpeed{
			Upload:   upload,
			Download: download,
		},
		EncMethod:  encMethod,
		PricePerGB: pricePerGB,
		Version:    version,
		LockerID:   lockerID,
		Coins:      coins,
		PubKey:     pubKey,
		Signature:  signature,
	}
}

type MsgUpdateNodeStatus struct {
	From csdkTypes.AccAddress `json:"from"`

	VPNID  string `json:"vpnid"`
	Status bool   `json:"status"`
}

func (msg MsgUpdateNodeStatus) Type() string {
	return "msg_update_node_status"
}

func (msg MsgUpdateNodeStatus) ValidateBasic() csdkTypes.Error {
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
	return "vpn"
}

func NewMsgUpdateNodeStatus(from csdkTypes.AccAddress, vpnID string, status bool) *MsgUpdateNodeStatus {
	return &MsgUpdateNodeStatus{
		From:   from,
		VPNID:  vpnID,
		Status: status,
	}
}

type MsgPayVPNService struct {
	From csdkTypes.AccAddress `json:"from"`

	VPNID string `json:"vpnid"`

	LockerID  string          `json:"locker_id"`
	Coins     csdkTypes.Coins `json:"coins"`
	PubKey    crypto.PubKey   `json:"pub_key"`
	Signature []byte          `json:"signature"`
}

func (msg MsgPayVPNService) Type() string {
	return "msg_pay_vpn_service"
}

func (msg MsgPayVPNService) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgPayVPNService) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgPayVPNService) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgPayVPNService) Route() string {
	return "vpn"
}

func NewMsgPayVPNService(from csdkTypes.AccAddress, vpnID string,
	lockerID string, coins csdkTypes.Coins, pubKey crypto.PubKey, signature []byte) *MsgPayVPNService {

	return &MsgPayVPNService{
		From:      from,
		VPNID:     vpnID,
		LockerID:  lockerID,
		Coins:     coins,
		PubKey:    pubKey,
		Signature: signature,
	}
}

type MsgUpdateSessionStatus struct {
	From csdkTypes.AccAddress `json:"from"`

	SessionID string `json:"session_id"`
	Status    bool   `json:"status"`
}

func (msg MsgUpdateSessionStatus) Type() string {
	return "msg_update_session_status"
}

func (msg MsgUpdateSessionStatus) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgUpdateSessionStatus) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgUpdateSessionStatus) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateSessionStatus) Route() string {
	return "vpn"
}

func NewMsgUpdateSessionStatus(from csdkTypes.AccAddress, sessionID string, status bool) *MsgUpdateSessionStatus {
	return &MsgUpdateSessionStatus{
		From:      from,
		SessionID: sessionID,
		Status:    status,
	}
}

type MsgDeregisterNode struct {
	From csdkTypes.AccAddress `json:"from"`

	VPNID string `json:"vpnid"`

	LockerID  string        `json:"locker_id"`
	PubKey    crypto.PubKey `json:"pub_key"`
	Signature []byte        `json:"signature"`
}

func (msg MsgDeregisterNode) Type() string {
	return "msg_deregister_node"
}

func (msg MsgDeregisterNode) ValidateBasic() csdkTypes.Error {
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
	return "vpn"
}

func NewMsgDeregisterNode(from csdkTypes.AccAddress, vpnID string,
	lockerID string, pubKey crypto.PubKey, signature []byte) *MsgDeregisterNode {

	return &MsgDeregisterNode{
		From:      from,
		VPNID:     vpnID,
		LockerID:  lockerID,
		PubKey:    pubKey,
		Signature: signature,
	}
}
