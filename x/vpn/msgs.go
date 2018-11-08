package vpn

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubtypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type MsgRegisterVpn struct {
	From    sdkTypes.AccAddress
	Coins   sdkTypes.Coins
	Details hubtypes.VpnDetails
}

func (msg MsgRegisterVpn) Type() string {
	return "RegisterVpn"
}

func (msg MsgRegisterVpn) ValidateBasic() sdkTypes.Error {
	return nil
}

func (msg MsgRegisterVpn) GetSignBytes() []byte {
	MsgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return MsgBytes
}

func (msg MsgRegisterVpn) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{msg.From}
}

func (msg MsgRegisterVpn) Route() string { return msg.Type() }

func NewRegisterVpnMsg(from sdkTypes.AccAddress, ip string, port string, coins sdkTypes.Coins, pricePerGb int64, upload int64, download int64,
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
		},
	}
}
