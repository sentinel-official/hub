package vpn

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec){
	cdc.RegisterConcrete(MsgRegisterVpn{},"vpn/register",nil)
}

var msgCdc = codec.New()

func init()  {
	RegisterCodec(msgCdc)
}