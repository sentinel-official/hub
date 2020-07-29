package types

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/x/vpn/deposit"
	"github.com/sentinel-official/hub/x/vpn/node"
	"github.com/sentinel-official/hub/x/vpn/plan"
	"github.com/sentinel-official/hub/x/vpn/provider"
	"github.com/sentinel-official/hub/x/vpn/session"
	"github.com/sentinel-official/hub/x/vpn/subscription"
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
	plan.RegisterCodec(cdc)
	subscription.RegisterCodec(cdc)
	session.RegisterCodec(cdc)
}
