package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

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
	"github.com/sentinel-official/hub/x/vpn/expected"
)

type Keeper struct {
	Deposit      depositkeeper.Keeper
	Provider     providerkeeper.Keeper
	Node         nodekeeper.Keeper
	Plan         plankeeper.Keeper
	Subscription subscriptionkeeper.Keeper
	Session      sessionkeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramsKeeper expected.ParamsKeeper,
	accountKeeper expected.AccountKeeper,
	bankKeeper expected.BankKeeper,
	distributionKeeper expected.DistributionKeeper,
	feeCollectorName string,
) Keeper {
	var (
		depositKeeper      = depositkeeper.NewKeeper(cdc, key)
		providerKeeper     = providerkeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(providertypes.ParamsSubspace))
		nodeKeeper         = nodekeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(nodetypes.ParamsSubspace))
		planKeeper         = plankeeper.NewKeeper(cdc, key)
		subscriptionKeeper = subscriptionkeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(subscriptiontypes.ParamsSubspace), feeCollectorName)
		sessionKeeper      = sessionkeeper.NewKeeper(cdc, key, paramsKeeper.Subspace(sessiontypes.ParamsSubspace), feeCollectorName)
	)

	depositKeeper.WithBankKeeper(bankKeeper)

	providerKeeper.WithDistributionKeeper(distributionKeeper)

	nodeKeeper.WithDistributionKeeper(distributionKeeper)
	nodeKeeper.WithProviderKeeper(&providerKeeper)
	nodeKeeper.WithPlanKeeper(&planKeeper)

	planKeeper.WithProviderKeeper(&providerKeeper)
	planKeeper.WithNodeKeeper(&nodeKeeper)

	subscriptionKeeper.WithBankKeeper(bankKeeper)
	subscriptionKeeper.WithDepositKeeper(&depositKeeper)
	subscriptionKeeper.WithProviderKeeper(&providerKeeper)
	subscriptionKeeper.WithNodeKeeper(&nodeKeeper)
	subscriptionKeeper.WithPlanKeeper(&planKeeper)
	subscriptionKeeper.WithSessionKeeper(&sessionKeeper)

	sessionKeeper.WithAccountKeeper(accountKeeper)
	sessionKeeper.WithBankKeeper(bankKeeper)
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
