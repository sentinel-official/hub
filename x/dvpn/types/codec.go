package types

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/x/dvpn/provider"
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
	provider.RegisterCodec(cdc)
}
