package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRegisterVpn{}, "x/vpn/msg_register_vpn", nil)
	cdc.RegisterConcrete(MsgAliveNode{}, "x/vpn/msg_alive_node", nil)
	cdc.RegisterConcrete(sdkTypes.IBCMsgRegisterVpn{}, "x/vpn/ibc_msg_register_vpn", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
