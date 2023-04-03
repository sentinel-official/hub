package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	providerkeeper "github.com/sentinel-official/hub/x/provider/keeper"
)

type Migrator struct {
	provider providerkeeper.Migrator
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{
		provider: providerkeeper.NewMigrator(k.Provider),
	}
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := m.provider.Migrate2to3(ctx); err != nil {
		return err
	}

	return nil
}
