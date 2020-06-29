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
	cdc.RegisterConcrete(MsgStartSubscription{}, "sentinel/MsgStartSubscription", nil)
	cdc.RegisterConcrete(MsgAddAddressForSubscription{}, "sentinel/MsgAddAddressForSubscription", nil)
	cdc.RegisterConcrete(MsgRemoveAddressForSubscription{}, "sentinel/MsgRemoveAddressForSubscription", nil)
	cdc.RegisterConcrete(MsgEndSubscription{}, "sentinel/MsgEndSubscription", nil)
}
