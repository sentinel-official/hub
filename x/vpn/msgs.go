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
	if msg.From == nil || msg.From.Empty() {
		return errorEmptyAddress()
	}

	if msg.APIPort <= 1024 || msg.APIPort > 65535 {
		return errorInvalidAPIPort()
	}

	if len(msg.Location.City) == 0 || len(msg.Location.Country) == 0 ||
		msg.Location.Latitude < sdkTypes.MinLatitude || msg.Location.Latitude > sdkTypes.MaxLatitude ||
		msg.Location.Longitude < sdkTypes.MinLongitude || msg.Location.Longitude > sdkTypes.MaxLongitude {
		return errorInvalidLocation()
	}

	if msg.NetSpeed.Download <= 0 || msg.NetSpeed.Upload <= 0 {
		return errorInvalidNetSpeed()
	}

	if len(msg.EncMethod) == 0 {
		return errorEmptyEncMethod()
	}

	if msg.PricePerGB < 0 {
		return errorInvalidPricePerGB()
	}

	if len(msg.Version) == 0 {
		return errorEmptyVersion()
	}

	if len(msg.LockerID) == 0 {
		return errorEmptyLockerID()
	}

	if msg.Coins == nil || msg.Coins.Len() == 0 || msg.Coins.IsValid() == false || msg.Coins.IsPositive() == false {
		return errorInvalidCoins()
	}

	if msg.PubKey == nil || len(msg.PubKey.Bytes()) == 0 {
		return errorEmptyPubKey()
	}

	if len(msg.Signature) == 0 {
		return errorEmptySignature()
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
	if msg.From == nil || msg.From.Empty() {
		return errorEmptyAddress()
	}

	if len(msg.VPNID) == 0 {
		return errorEmptyVPNID()
	}

	if len(msg.LockerID) == 0 {
		return errorEmptyLockerID()
	}

	if msg.Coins == nil || msg.Coins.Len() == 0 || msg.Coins.IsValid() == false || msg.Coins.IsPositive() == false {
		return errorInvalidCoins()
	}

	if msg.PubKey == nil || len(msg.PubKey.Bytes()) == 0 {
		return errorEmptyPubKey()
	}

	if len(msg.Signature) == 0 {
		return errorEmptySignature()
	}

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
	Status    string `json:"status"`
}

func (msg MsgUpdateSessionStatus) Type() string {
	return "msg_update_session_status"
}

func (msg MsgUpdateSessionStatus) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorEmptyAddress()
	}

	if len(msg.SessionID) == 0 {
		return errorEmptySessionID()
	}

	if msg.Status != "ACTIVE" && msg.Status != "INACTIVE" {
		return errorInvalidSessionStatus()
	}

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
	return "session"
}

func NewMsgUpdateSessionStatus(from csdkTypes.AccAddress, sessionID string, status string) *MsgUpdateSessionStatus {
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
	if msg.From == nil || msg.From.Empty() {
		return errorEmptyAddress()
	}

	if len(msg.VPNID) == 0 {
		return errorEmptyVPNID()
	}

	if len(msg.LockerID) == 0 {
		return errorEmptyLockerID()
	}

	if msg.PubKey == nil || len(msg.PubKey.Bytes()) == 0 {
		return errorEmptyPubKey()
	}

	if len(msg.Signature) == 0 {
		return errorEmptySignature()
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
