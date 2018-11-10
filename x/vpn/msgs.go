package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubtypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgRegisterVpn struct {
	From    csdkTypes.AccAddress
	Coins   csdkTypes.Coins
	Details hubtypes.VpnDetails
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
	latitude int64, longitude int64, city string, country string, enc_method string, version string) *MsgRegisterVpn {

	return &MsgRegisterVpn{
		From:  from,
		Coins: coins,
		Details: hubtypes.VpnDetails{
			Ip:         ip,
			Port:       port,
			PricePerGb: pricePerGb,
			NetSpeed: hubtypes.NetSpeed{
				Upload:   upload,
				Download: download,
			},
			Location: hubtypes.Location{
				Latitude:  latitude,
				Longitude: longitude,
				City:      city,
				Country:   country,
			},
			Version:   version,
			EncMethod: enc_method,
			Info: hubtypes.Info{
				Status:      false,
				BlockHeight: 0,
			},
		},
	}
}

type MsgAliveNode struct {
	From csdkTypes.AccAddress
}

func (msg MsgAliveNode) Type() string {
	return "AliveNode"
}

func (msg MsgAliveNode) ValidateBasic() csdkTypes.Error {
	return nil
}

func (msg MsgAliveNode) GetSignBytes() []byte {
	MsgBytes, err := json.Marshal(msg)

	if err != nil {
		return nil
	}

	return MsgBytes
}

func (msg MsgAliveNode) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgAliveNode) Route() string {
	return msg.Type()
}
