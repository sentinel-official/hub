package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/sentinel-official/hub/x/vpn/deposit"
	"github.com/sentinel-official/hub/x/vpn/node"
	"github.com/sentinel-official/hub/x/vpn/plan"
	"github.com/sentinel-official/hub/x/vpn/provider"
	"github.com/sentinel-official/hub/x/vpn/session"
	"github.com/sentinel-official/hub/x/vpn/subscription"
	"github.com/sentinel-official/hub/x/vpn/types"
)

type Keeper struct {
	Deposit      deposit.Keeper
	Provider     provider.Keeper
	Node         node.Keeper
	Plan         plan.Keeper
	Subscription subscription.Keeper
	Session      session.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper, bankKeeper bank.Keeper, supplyKeeper supply.Keeper) Keeper {
	nodeParams := paramsKeeper.Subspace(fmt.Sprintf("%s/%s", types.ModuleName, node.ParamsSubspace))
	sessionParams := paramsKeeper.Subspace(fmt.Sprintf("%s/%s", types.ModuleName, session.ParamsSubspace))

	depositKeeper := deposit.NewKeeper(cdc, key)
	providerKeeper := provider.NewKeeper(cdc, key)
	nodeKeeper := node.NewKeeper(cdc, key, nodeParams)
	planKeeper := plan.NewKeeper(cdc, key)
	subscriptionKeeper := subscription.NewKeeper(cdc, key)
	sessionKeeper := session.NewKeeper(cdc, key, sessionParams)

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
