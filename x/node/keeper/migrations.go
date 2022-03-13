package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/node/types"
)

type Migrator struct {
	k Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k: k}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	store := m.k.Store(ctx)

	if err := migrateNodeKeys(store); err != nil {
		return err
	}
	if err := migrateActiveNodeKeys(store); err != nil {
		return err
	}
	if err := migrateInactiveNodeKeys(store); err != nil {
		return err
	}
	if err := migrateActiveNodeForProviderKeys(store); err != nil {
		return err
	}
	if err := migrateInactiveNodeForProviderKeys(store); err != nil {
		return err
	}
	if err := migrateInactiveNodeAtKeys(store); err != nil {
		return err
	}

	return nil
}

func migrateNodeKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.NodeKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := types.NodeKey(iterator.Key())

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateActiveNodeKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.ActiveNodeKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := types.ActiveNodeKey(iterator.Key())

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateInactiveNodeKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.InactiveNodeKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := types.InactiveNodeKey(iterator.Key())

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateActiveNodeForProviderKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.ActiveNodeForProviderKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			providerAddr = iterator.Key()[:20]
			nodeAddr     = iterator.Key()[20:]
		)

		key := types.ActiveNodeForProviderKey(providerAddr, nodeAddr)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateInactiveNodeForProviderKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.InactiveNodeForProviderKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			providerAddr = iterator.Key()[:20]
			nodeAddr     = iterator.Key()[20:]
		)

		key := types.InactiveNodeForProviderKey(providerAddr, nodeAddr)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateInactiveNodeAtKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.InactiveNodeAtKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		at, err := sdk.ParseTimeBytes(iterator.Key()[:29])
		if err != nil {
			return err
		}

		addr := iterator.Key()[29:]
		key := types.InactiveNodeAtKey(at, addr)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}
