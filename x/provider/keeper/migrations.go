package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/provider/types"
)

type Migrator struct {
	k Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k: k}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	store := m.k.Store(ctx)
	return migrateProviderKeys(store)
}

func migrateProviderKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.ProviderKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := types.ProviderKey(iterator.Key())

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}
