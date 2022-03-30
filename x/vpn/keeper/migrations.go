package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	depositkeeper "github.com/sentinel-official/hub/x/deposit/keeper"
	nodekeeper "github.com/sentinel-official/hub/x/node/keeper"
	plankeeper "github.com/sentinel-official/hub/x/plan/keeper"
	providerkeeper "github.com/sentinel-official/hub/x/provider/keeper"
	sessionkeeper "github.com/sentinel-official/hub/x/session/keeper"
	subscriptionkeeper "github.com/sentinel-official/hub/x/subscription/keeper"
)

type Migrator struct {
	deposit      depositkeeper.Migrator
	provider     providerkeeper.Migrator
	node         nodekeeper.Migrator
	plan         plankeeper.Migrator
	subscription subscriptionkeeper.Migrator
	session      sessionkeeper.Migrator
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{
		deposit:      depositkeeper.NewMigrator(k.Deposit),
		provider:     providerkeeper.NewMigrator(k.Provider),
		node:         nodekeeper.NewMigrator(k.Node),
		plan:         plankeeper.NewMigrator(k.Plan),
		subscription: subscriptionkeeper.NewMigrator(k.Subscription),
		session:      sessionkeeper.NewMigrator(k.Session),
	}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	if err := m.deposit.Migrate1to2(ctx); err != nil {
		return err
	}
	if err := m.provider.Migrate1to2(ctx); err != nil {
		return err
	}
	if err := m.node.Migrate1to2(ctx); err != nil {
		return err
	}
	if err := m.plan.Migrate1to2(ctx); err != nil {
		return err
	}
	if err := m.subscription.Migrate1to2(ctx); err != nil {
		return err
	}
	if err := m.session.Migrate1to2(ctx); err != nil {
		return err
	}

	return nil
}
