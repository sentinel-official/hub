package types

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/x/dvpn/deposit"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/session"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
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
	deposit.RegisterCodec(cdc)
	provider.RegisterCodec(cdc)
	node.RegisterCodec(cdc)
	subscription.RegisterCodec(cdc)
	session.RegisterCodec(cdc)
}
