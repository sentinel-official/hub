package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/sentinel-official/hub/x/dvpn/deposit"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

type Keeper struct {
	Deposit      deposit.Keeper
	Provider     provider.Keeper
	Node         node.Keeper
	Subscription subscription.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey,
	bankKeeper bank.Keeper, supplyKeeper supply.Keeper) Keeper {
	dk := deposit.NewKeeper(cdc, key)
	pk := provider.NewKeeper(cdc, key)
	nk := node.NewKeeper(cdc, key)
	sk := subscription.NewKeeper(cdc, key)

	dk.WithSupplyKeeper(supplyKeeper)

	nk.WithProviderKeeper(&pk)
	nk.WithSubscriptionKeeper(&sk)

	sk.WithDepositKeeper(&dk)
	sk.WithBankKeeper(bankKeeper)
	sk.WithProviderKeeper(&pk)
	sk.WithNodeKeeper(&nk)

	return Keeper{
		Deposit:      dk,
		Provider:     pk,
		Node:         nk,
		Subscription: sk,
	}
}
