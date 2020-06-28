package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/sentinel-official/hub/x/dvpn/deposit"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/plan"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/session"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

type Keeper struct {
	Deposit      deposit.Keeper
	Provider     provider.Keeper
	Node         node.Keeper
	Plan         plan.Keeper
	Subscription subscription.Keeper
	Session      session.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey,
	bankKeeper bank.Keeper, supplyKeeper supply.Keeper) Keeper {
	depositKeeper := deposit.NewKeeper(cdc, key)
	providerKeeper := provider.NewKeeper(cdc, key)
	nodeKeeper := node.NewKeeper(cdc, key)
	planKeeper := plan.NewKeeper(cdc, key)
	subscriptionKeeper := subscription.NewKeeper(cdc, key)
	sessionKeeper := session.NewKeeper(cdc, key)

	depositKeeper.WithSupplyKeeper(supplyKeeper)

	nodeKeeper.WithProviderKeeper(&providerKeeper)
	nodeKeeper.WithPlanKeeper(&planKeeper)

	planKeeper.WithProviderKeeper(&providerKeeper)
	planKeeper.WithNodeKeeper(&nodeKeeper)

	subscriptionKeeper.WithDepositKeeper(&depositKeeper)
	subscriptionKeeper.WithBankKeeper(bankKeeper)
	subscriptionKeeper.WithNodeKeeper(&nodeKeeper)
	subscriptionKeeper.WithPlanKeeper(&planKeeper)

	sessionKeeper.WithPlanKeeper(&planKeeper)
	sessionKeeper.WithSubscriptionKeeper(&subscriptionKeeper)

	return Keeper{
		Deposit:      depositKeeper,
		Provider:     providerKeeper,
		Node:         nodeKeeper,
		Plan:         planKeeper,
		Subscription: subscriptionKeeper,
		Session:      sessionKeeper,
	}
}
