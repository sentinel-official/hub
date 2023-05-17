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

func (k Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := k.migrateProviders(ctx); err != nil {
		return err
	}

	return nil
}

func (k Migrator) migrateProviders(ctx sdk.Context) error {
	store := prefix.NewStore(k.Store(ctx), []byte{0x10})

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var value v1types.Provider
		k.cdc.MustUnmarshal(iter.Value(), &value)
		store.Delete(iter.Key())

		provider := types.Provider{
			Address:     value.Address,
			Name:        value.Name,
			Identity:    value.Identity,
			Website:     value.Website,
			Description: value.Description,
			Status:      hubtypes.StatusActive,
		}

		k.SetProvider(ctx, provider)
	}

	return nil
}
