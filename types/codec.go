package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Interface)(nil), nil)
	cdc.RegisterConcrete(&AppAccount{}, "types/app_account", nil)
	cdc.RegisterConcrete(IBCPacket{}, "types/ibc_packet", nil)
	cdc.RegisterConcrete(CoinLocker{}, "types/coin_locker", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
