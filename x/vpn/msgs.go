package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgRegisterVpn struct {
	From    csdkTypes.AccAddress
	Coins   csdkTypes.Coins
	Details sdkTypes.VpnDetails
}

func (msg MsgRegisterVpn) Type() string {
	return "vpn"
}

func (msg MsgRegisterVpn) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgRegisterVpn) GetSignBytes() []byte {
	MsgBytes, err := json.Marshal(msg)

	if err != nil {
		return nil
	}

	return MsgBytes
}

func (msg MsgRegisterVpn) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgRegisterVpn) Route() string {
	return msg.Type()
}

func NewRegisterVpnMsg(from csdkTypes.AccAddress, ip string, port string, coins csdkTypes.Coins, pricePerGb int64, upload int64, download int64,
	latitude int64, longitude int64, city string, country string, encMethod string, version string) *MsgRegisterVpn {

	return &MsgRegisterVpn{
		From:  from,
		Coins: coins,
		Details: sdkTypes.VpnDetails{
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

type MsgNodeStatus struct {
	From   csdkTypes.AccAddress
	VpnId  string
	Status bool
}

func (msg MsgNodeStatus) Type() string {
	return "AliveNode"
}

func (msg MsgNodeStatus) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgNodeStatus) GetSignBytes() []byte {
	MsgBytes, err := json.Marshal(msg)

	if err != nil {
		return nil
	}

	return MsgBytes
}

func (msg MsgNodeStatus) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgNodeStatus) Route() string {
	return msg.Type()
}

func NewNodeStatusMsg(from csdkTypes.AccAddress, vpnId string, status bool) MsgNodeStatus {
	return MsgNodeStatus{
		From:   from,
		VpnId:  vpnId,
		Status: status,
	}
}
