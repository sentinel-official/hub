package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	params "github.com/cosmos/cosmos-sdk/x/params/keeper"

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

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, paramsKeeper params.Keeper,
	bankKeeper bank.Keeper, distributionKeeper distribution.Keeper) Keeper {
	var (
		depositKeeper      = deposit.NewKeeper(cdc, key)
		providerKeeper     = provider.NewKeeper(cdc, key, paramsKeeper.Subspace(provider.ParamsSubspace))
		nodeKeeper         = node.NewKeeper(cdc, key, paramsKeeper.Subspace(node.ParamsSubspace))
		planKeeper         = plan.NewKeeper(cdc, key)
		subscriptionKeeper = subscription.NewKeeper(cdc, key, paramsKeeper.Subspace(subscription.ParamsSubspace))
		sessionKeeper      = session.NewKeeper(cdc, key, paramsKeeper.Subspace(session.ParamsSubspace))
	)

	depositKeeper.WithBankKeeper(bankKeeper)

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
