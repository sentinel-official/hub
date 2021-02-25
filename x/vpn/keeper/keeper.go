package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
)

type Keeper struct {
	Deposit      deposit.Keeper
	Provider     provider.Keeper
	Node         node.Keeper
	Plan         plan.Keeper
	Subscription subscription.Keeper
	Session      session.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper, accountKeeper auth.AccountKeeper,
	bankKeeper bank.Keeper, distributionKeeper distribution.Keeper, supplyKeeper supply.Keeper) Keeper {
	var (
		depositKeeper      = deposit.NewKeeper(cdc, key)
		providerKeeper     = provider.NewKeeper(cdc, key, paramsKeeper.Subspace(provider.ParamsSubspace))
		nodeKeeper         = node.NewKeeper(cdc, key, paramsKeeper.Subspace(node.ParamsSubspace))
		planKeeper         = plan.NewKeeper(cdc, key)
		subscriptionKeeper = subscription.NewKeeper(cdc, key, paramsKeeper.Subspace(subscription.ParamsSubspace))
		sessionKeeper      = session.NewKeeper(cdc, key, paramsKeeper.Subspace(session.ParamsSubspace))
	)

	depositKeeper.WithSupplyKeeper(supplyKeeper)

	providerKeeper.WithDistributionKeeper(distributionKeeper)

	nodeKeeper.WithDistributionKeeper(distributionKeeper)
	nodeKeeper.WithProviderKeeper(&providerKeeper)
	nodeKeeper.WithPlanKeeper(&planKeeper)

	planKeeper.WithProviderKeeper(&providerKeeper)
	planKeeper.WithNodeKeeper(&nodeKeeper)

	subscriptionKeeper.WithDepositKeeper(&depositKeeper)
	subscriptionKeeper.WithBankKeeper(bankKeeper)
	subscriptionKeeper.WithNodeKeeper(&nodeKeeper)
	subscriptionKeeper.WithPlanKeeper(&planKeeper)

	sessionKeeper.WithAccountKeeper(accountKeeper)
	sessionKeeper.WithDepositKeeper(depositKeeper)
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
