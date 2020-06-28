package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var (
	ModuleCdc *codec.Codec
)

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgStartSubscription{}, "x/dvpn/subscription/MsgStartSubscription", nil)
	cdc.RegisterConcrete(MsgAddAddressForSubscription{}, "x/dvpn/subscription/MsgAddAddressForSubscription", nil)
	cdc.RegisterConcrete(MsgRemoveAddressForSubscription{}, "x/dvpn/subscription/MsgRemoveAddressForSubscription", nil)
	cdc.RegisterConcrete(MsgEndSubscription{}, "x/dvpn/subscription/MsgEndSubscription", nil)
}
