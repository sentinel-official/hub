package vpn

import (

"encoding/json"

sdkTypes "github.com/cosmos/cosmos-sdk/types"
vpnTypes "github.com/ironman0x7b2/sentinel-hub/types"

)
//var _ sdkTypes.Msg = MsgRegisterVpn{}
type MsgRegisterVpn struct {
	From     sdkTypes.AccAddress
	Coin     sdkTypes.Coins
	Register vpnTypes.Registervpn
}

func (msg MsgRegisterVpn) Type() string {
	return "vpn"
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

func (msg MsgRegisterVpn) Route() string { return "bank" }

func CreateRegisterVpnMsg(from sdkTypes.AccAddress, ip string, port string, coins sdkTypes.Coins, ppgb int64,uploadSpeed int64, downloadSpeed int64, latitude int64,longitude int64,city string,country string, enc_method string,node_type string,version string) *MsgRegisterVpn{
	return &MsgRegisterVpn{
	From:from,
	Coin:coins,
	Register: vpnTypes.Registervpn{
		Ip:ip,
		Port:port,
		PricePerGb:ppgb,
		NetSpeed:vpnTypes.NetSpeed{
			UploadSpeed:uploadSpeed,
			DownloadSpeed:downloadSpeed,
		},
		Location:vpnTypes.Location{
			Latitude:latitude,
			Longitude:longitude,
			City:city,
			Country:country,
		},
		NodeType:node_type,
		Version:version,
		EncMethod:enc_method,
	},
	}
}
