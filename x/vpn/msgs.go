package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"strconv"
)

type MsgRegisterNode struct {
	From    csdkTypes.AccAddress
	Coins   csdkTypes.Coins
	Details sdkTypes.VPNDetails
}

func (msg MsgRegisterNode) Type() string {
	return "msg_register_node"
}

func (msg MsgRegisterNode) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgRegisterNode) GetSignBytes() []byte {
	MsgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	return MsgBytes
}

func (msg MsgRegisterNode) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgRegisterNode) Route() string {
	return "vpn"
}

func NewRegisterVPNMsg(from csdkTypes.AccAddress, coins csdkTypes.Coins,
	apiPort string, vpnPort string, pubkey crypto.PubKey,
	upload int64, download int64,
	latitude int64, longitude int64, city string, country string,
	pricePerGb int64, encMethod string, version string, sequence int64, signature []byte) *MsgRegisterNode {

	vpnID := from.String() + "/" + strconv.Itoa(int(sequence))
	storeKey := "vpn"
	//TODO: Replace vpnID with keeper.storeKey type

	return &MsgRegisterNode{
		From:  from,
		Coins: coins,
		Details: sdkTypes.VPNDetails{
			ApiPort:    apiPort,
			VPNPort:    vpnPort,
			Pubkey:     pubkey,
			Address:    from,
			PricePerGb: pricePerGb,
			NetSpeed: sdkTypes.NetSpeed{
				Upload:   upload,
				Download: download,
			},
			Location: sdkTypes.Location{
				Latitude:  latitude,
				Longitude: longitude,
				City:      city,
				Country:   country,
			},
			Version:   version,
			EncMethod: encMethod,
			Info: sdkTypes.Info{
				Status:      false,
				BlockHeight: 0,
			},
			LockerId:  storeKey + "/" + vpnID,
			Signature: signature,
		},
	}
}

type MsgUpdateNodeStatus struct {
	From   csdkTypes.AccAddress
	VPNID  string
	Status bool
}

func (msg MsgUpdateNodeStatus) Type() string {
	return "msg_update_node_status"
}

func (msg MsgUpdateNodeStatus) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgUpdateNodeStatus) GetSignBytes() []byte {
	MsgBytes, err := json.Marshal(msg)

	if err != nil {
		return nil
	}

	return MsgBytes
}

func (msg MsgUpdateNodeStatus) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateNodeStatus) Route() string {
	return "vpn"
}

func NewNodeStatusMsg(from csdkTypes.AccAddress, vpnID string, status bool) MsgUpdateNodeStatus {
	return MsgUpdateNodeStatus{
		From:   from,
		VPNID:  vpnID,
		Status: status,
	}
}

type MsgPayVPNService struct {
	Coins     csdkTypes.Coins
	VPNID     string
	From      csdkTypes.AccAddress
	Pubkey    crypto.PubKey
	LockerId  string
	Signature []byte
}

func NewMsgPayVPNService(coins csdkTypes.Coins, vpnID string, from csdkTypes.AccAddress, sequence int64, pubkey crypto.PubKey, signature []byte) MsgPayVPNService {

	sessionID := from.String() + "/" + strconv.Itoa(int(sequence))
	storeKey := "session"

	return MsgPayVPNService{
		Coins:     coins,
		VPNID:     vpnID,
		From:      from,
		Pubkey:    pubkey,
		LockerId:  storeKey + "/" + sessionID,
		Signature: signature,
	}

}

func (msg MsgPayVPNService) Type() string {
	return "pay-vpn-service"
}

func (msg MsgPayVPNService) GetSignBytes() []byte {
	byte_format, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return byte_format
}

func (msg MsgPayVPNService) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgPayVPNService) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msc MsgPayVPNService) Route() string {
	return "vpn"
}

type MsgUpdateSessionStatus struct {
	From      csdkTypes.AccAddress
	SessionID string
	Status    bool
}

func (msg MsgUpdateSessionStatus) Type() string {
	return "msg_update_node_status"
}

func (msg MsgUpdateSessionStatus) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgUpdateSessionStatus) GetSignBytes() []byte {
	MsgBytes, err := json.Marshal(msg)

	if err != nil {
		return nil
	}

	return MsgBytes
}

func (msg MsgUpdateSessionStatus) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateSessionStatus) Route() string {
	return "vpn"
}

func NewSessionStatusMsg(from csdkTypes.AccAddress, sessionID string, status bool) MsgUpdateSessionStatus {
	return MsgUpdateSessionStatus{
		From:      from,
		SessionID: sessionID,
		Status:    status,
	}
}
