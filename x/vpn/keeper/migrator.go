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

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := m.deposit.Migrate2to3(ctx); err != nil {
		panic(err)
	}
	if err := m.provider.Migrate2to3(ctx); err != nil {
		return err
	}
	if err := m.node.Migrate2to3(ctx); err != nil {
		return err
	}
	if err := m.plan.Migrate2to3(ctx); err != nil {
		return err
	}
	if err := m.subscription.Migrate2to3(ctx); err != nil {
		return err
	}
	if err := m.session.Migrate2to3(ctx); err != nil {
		return err
	}

	return nil
}
