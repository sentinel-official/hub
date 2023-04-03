package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	v1types "github.com/sentinel-official/hub/x/provider/legacy/v1/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := m.migrateProviders(ctx); err != nil {
		return err
	}

	return nil
}

func (m Migrator) migrateProviders(ctx sdk.Context) error {
	store := prefix.NewStore(m.Store(ctx), types.ProviderKeyPrefix)

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var value v1types.Provider
		m.cdc.MustUnmarshal(iter.Value(), &value)
		store.Delete(iter.Key())

		provider := types.Provider{
			Address:     value.Address,
			Name:        value.Name,
			Identity:    value.Identity,
			Website:     value.Website,
			Description: value.Description,
			Status:      hubtypes.StatusActive,
		}

		m.SetProvider(ctx, provider)
	}

	return nil
}
