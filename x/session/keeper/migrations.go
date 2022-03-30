package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

type Migrator struct {
	k Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k: k}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	store := m.k.Store(ctx)

	if err := migrateInactiveSessionForAddressKeys(store); err != nil {
		return err
	}
	if err := migrateActiveSessionForAddressKeys(store); err != nil {
		return err
	}

	count := m.k.GetCount(ctx)
	for id := uint64(0); id <= count; id++ {
		item, found := m.k.GetSession(ctx, id)
		if !found {
			continue
		}
		if !item.Status.Equal(hubtypes.StatusInactive) {
			continue
		}

		m.k.DeleteSession(ctx, item.Id)
		m.k.DeleteInactiveSessionForAddress(ctx, item.GetAddress(), item.Id)
	}

	return nil
}

func migrateInactiveSessionForAddressKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.InactiveSessionForAddressKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			addr = iterator.Key()[:20]
			id   = sdk.BigEndianToUint64(iterator.Key()[20:])
		)

		key := types.InactiveSessionForAddressKey(addr, id)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateActiveSessionForAddressKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.ActiveSessionForAddressKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			addr = iterator.Key()[:20]
			id   = sdk.BigEndianToUint64(iterator.Key()[20:])
		)

		key := types.ActiveSessionForAddressKey(addr, id)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}
