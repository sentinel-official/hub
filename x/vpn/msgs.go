package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
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
	ip string, port string,
	upload int64, download int64,
	latitude int64, longitude int64, city string, country string,
	pricePerGb int64, encMethod string, version string) *MsgRegisterNode {
	return &MsgRegisterNode{
		From:  from,
		Coins: coins,
		Details: sdkTypes.VPNDetails{
			Ip:         ip,
			Port:       port,
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
