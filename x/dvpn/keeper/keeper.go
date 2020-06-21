package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
)

type Keeper struct {
	Provider provider.Keeper
	Node     node.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	pk := provider.NewKeeper(cdc, key)
	nk := node.NewKeeper(cdc, key, pk)

	return Keeper{
		Provider: pk,
		Node:     nk,
	}
}
