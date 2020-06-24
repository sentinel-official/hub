package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

type Keeper struct {
	Provider     provider.Keeper
	Node         node.Keeper
	Subscription subscription.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	pk := provider.NewKeeper(cdc, key)
	nk := node.NewKeeper(cdc, key)
	sk := subscription.NewKeeper(cdc, key)

	nk.WithProviderKeeper(&pk)
	nk.WithSubscriptionKeeper(&sk)

	sk.WithProviderKeeper(&pk)
	sk.WithNodeKeeper(&nk)

	return Keeper{
		Provider:     pk,
		Node:         nk,
		Subscription: sk,
	}
}
