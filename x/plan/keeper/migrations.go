package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/plan/types"
)

type Migrator struct {
	k Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k: k}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	store := m.k.Store(ctx)

	if err := migrateActivePlanForProviderKeys(store); err != nil {
		return err
	}
	if err := migrateInactivePlanForProviderKeys(store); err != nil {
		return err
	}
	if err := migrateNodeForPlanKeys(store); err != nil {
		return err
	}
	if err := migrateCountForNodeByProviderKeys(store); err != nil {
		return err
	}

	return nil
}

func migrateActivePlanForProviderKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.ActivePlanForProviderKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			addr = iterator.Key()[:20]
			id   = sdk.BigEndianToUint64(iterator.Key()[20:])
		)

		key := types.ActivePlanForProviderKey(addr, id)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateInactivePlanForProviderKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.InactivePlanForProviderKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			addr = iterator.Key()[:20]
			id   = sdk.BigEndianToUint64(iterator.Key()[20:])
		)

		key := types.InactivePlanForProviderKey(addr, id)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateNodeForPlanKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.NodeForPlanKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			id   = sdk.BigEndianToUint64(iterator.Key()[:8])
			addr = iterator.Key()[8:]
		)

		key := types.NodeForPlanKey(id, addr)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateCountForNodeByProviderKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.CountForNodeByProviderKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			provider = iterator.Key()[:20]
			node     = iterator.Key()[20:]
		)

		key := types.CountForNodeByProviderKey(provider, node)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}
