package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/tendermint/tendermint/crypto"
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
	pricePerGb int64, encMethod string, version string) *MsgRegisterNode {
	return &MsgRegisterNode{
		From:  from,
		Coins: coins,
		Details: sdkTypes.VPNDetails{
			ApiPort:    apiPort,
			VpnPort:    vpnPort,
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
		},
	}
}

type MsgUpdateNodeStatus struct {
	From   csdkTypes.AccAddress
	VPNId  string
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

func NewNodeStatusMsg(from csdkTypes.AccAddress, vpnId string, status bool) MsgUpdateNodeStatus {
	return MsgUpdateNodeStatus{
		From:   from,
		VPNId:  vpnId,
		Status: status,
	}
}

type MsgPayVpnService struct {
	Coins   csdkTypes.Coins
	Vpnaddr csdkTypes.AccAddress
	From    csdkTypes.AccAddress
	Pubkey  crypto.PubKey
}

func NewMsgPayVpnService(coins csdkTypes.Coins, vaddr csdkTypes.AccAddress, from csdkTypes.AccAddress, pubkey crypto.PubKey) MsgPayVpnService {
	return MsgPayVpnService{
		Coins:   coins,
		Vpnaddr: vaddr,
		From:    from,
		Pubkey:  pubkey,
	}

}

func (msg MsgPayVpnService) Type() string {
	return "pay-vpn-service"
}

func (msg MsgPayVpnService) GetSignBytes() []byte {
	byte_format, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return byte_format
}

func (msg MsgPayVpnService) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgPayVpnService) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msc MsgPayVpnService) Route() string {
	return "vpn"
}
