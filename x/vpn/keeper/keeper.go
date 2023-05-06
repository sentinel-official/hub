package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

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
	"github.com/sentinel-official/hub/x/vpn/types"
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
	key sdk.StoreKey,
	paramsKeeper expected.ParamsKeeper,
	accountKeeper expected.AccountKeeper,
	bankKeeper expected.BankKeeper,
	distributionKeeper expected.DistributionKeeper,
	feeCollectorName string,
) (k Keeper) {
	k.Deposit = depositkeeper.NewKeeper(cdc, key)
	k.Deposit.WithBankKeeper(bankKeeper)

	k.Provider = providerkeeper.NewKeeper(
		cdc, key, paramsKeeper.Subspace(fmt.Sprintf("%s/%s", types.ModuleName, providertypes.ModuleName)),
	)
	k.Provider.WithDistributionKeeper(distributionKeeper)

	k.Node = nodekeeper.NewKeeper(
		cdc, key, paramsKeeper.Subspace(fmt.Sprintf("%s/%s", types.ModuleName, nodetypes.ModuleName)),
	)
	k.Node.WithDistributionKeeper(distributionKeeper)
	k.Node.WithProviderKeeper(&k.Provider)

	k.Plan = plankeeper.NewKeeper(cdc, key)
	k.Plan.WithProviderKeeper(&k.Provider)
	k.Plan.WithNodeKeeper(&k.Node)

	k.Subscription = subscriptionkeeper.NewKeeper(
		cdc, key, paramsKeeper.Subspace(fmt.Sprintf("%s/%s", types.ModuleName, subscriptiontypes.ModuleName)),
		feeCollectorName,
	)
	k.Subscription.WithBankKeeper(bankKeeper)
	k.Subscription.WithDepositKeeper(&k.Deposit)
	k.Subscription.WithProviderKeeper(&k.Provider)
	k.Subscription.WithNodeKeeper(&k.Node)
	k.Subscription.WithPlanKeeper(&k.Plan)
	k.Subscription.WithSessionKeeper(&k.Session)

	k.Session = sessionkeeper.NewKeeper(
		cdc, key, paramsKeeper.Subspace(fmt.Sprintf("%s/%s", types.ModuleName, sessiontypes.ModuleName)),
		feeCollectorName,
	)
	k.Session.WithAccountKeeper(accountKeeper)
	k.Session.WithBankKeeper(bankKeeper)
	k.Session.WithDepositKeeper(&k.Deposit)
	k.Session.WithNodeKeeper(&k.Node)
	k.Session.WithSubscriptionKeeper(&k.Subscription)

	return k
}
