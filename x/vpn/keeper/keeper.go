package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	depositkeeper "github.com/sentinel-official/hub/x/deposit/keeper"
	nodekeeper "github.com/sentinel-official/hub/x/node/keeper"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plankeeper "github.com/sentinel-official/hub/x/plan/keeper"
	providerkeeper "github.com/sentinel-official/hub/x/provider/keeper"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessionkeeper "github.com/sentinel-official/hub/x/session/keeper"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptionkeeper "github.com/sentinel-official/hub/x/subscription/keeper"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

type Keeper struct {
	Deposit      depositkeeper.Keeper
	Provider     providerkeeper.Keeper
	Node         nodekeeper.Keeper
	Plan         plankeeper.Keeper
	Subscription subscriptionkeeper.Keeper
	Session      sessionkeeper.Keeper
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, paramsKeeper paramskeeper.Keeper, accountKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper, distributionKeeper distributionkeeper.Keeper) Keeper {
	var (
		depositKeeper      = depositkeeper.NewKeeper(cdc, key)
		providerKeeper     = providerkeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(providertypes.ParamsSubspace))
		nodeKeeper         = nodekeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(nodetypes.ParamsSubspace))
		planKeeper         = plankeeper.NewKeeper(cdc, key)
		subscriptionKeeper = subscriptionkeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(subscriptiontypes.ParamsSubspace))
		sessionKeeper      = sessionkeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(sessiontypes.ParamsSubspace))
	)

	depositKeeper.WithBankKeeper(bankKeeper)

	providerKeeper.WithDistributionKeeper(distributionKeeper)
	providerKeeper.WithAccountKeeper(accountKeeper)

	nodeKeeper.WithDistributionKeeper(distributionKeeper)
	nodeKeeper.WithProviderKeeper(&providerKeeper)
	nodeKeeper.WithPlanKeeper(&planKeeper)
	nodeKeeper.WithAccountKeeper(&accountKeeper)

	planKeeper.WithProviderKeeper(&providerKeeper)
	planKeeper.WithNodeKeeper(&nodeKeeper)

	subscriptionKeeper.WithDepositKeeper(&depositKeeper)
	subscriptionKeeper.WithBankKeeper(bankKeeper)
	subscriptionKeeper.WithNodeKeeper(&nodeKeeper)
	subscriptionKeeper.WithPlanKeeper(&planKeeper)

	sessionKeeper.WithAccountKeeper(accountKeeper)
	sessionKeeper.WithDepositKeeper(&depositKeeper)
	sessionKeeper.WithNodeKeeper(&nodeKeeper)
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
