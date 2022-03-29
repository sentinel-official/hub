package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

type Migrator struct {
	k Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k: k}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	store := m.k.Store(ctx)

	if err := migrateActiveSubscriptionForAddressKeys(store); err != nil {
		return err
	}
	if err := migrateInactiveSubscriptionForAddressKeys(store); err != nil {
		return err
	}
	if err := migrateQuotaKeys(store); err != nil {
		return err
	}

	count := m.k.GetCount(ctx)
	for id := uint64(0); id <= count; id++ {
		item, found := m.k.GetSubscription(ctx, id)
		if !found {
			continue
		}
		if !item.Status.Equal(hubtypes.StatusInactive) {
			continue
		}

		m.k.DeleteSubscription(ctx, item.Id)
		m.k.IterateQuotas(ctx, item.Id, func(_ int, quota types.Quota) bool {
			address := quota.GetAddress()
			m.k.DeleteQuota(ctx, item.Id, address)
			m.k.DeleteInactiveSubscriptionForAddress(ctx, address, item.Id)

			return false
		})
	}

	return nil
}

func migrateActiveSubscriptionForAddressKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.ActiveSubscriptionForAddressKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			addr = iterator.Key()[:20]
			id   = sdk.BigEndianToUint64(iterator.Key()[20:])
		)

		key := types.ActiveSubscriptionForAddressKey(addr, id)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateInactiveSubscriptionForAddressKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.InactiveSubscriptionForAddressKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			addr = iterator.Key()[:20]
			id   = sdk.BigEndianToUint64(iterator.Key()[20:])
		)

		key := types.InactiveSubscriptionForAddressKey(addr, id)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}

func migrateQuotaKeys(parent sdk.KVStore) error {
	child := prefix.NewStore(parent, types.QuotaKeyPrefix)

	iterator := child.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var (
			id   = sdk.BigEndianToUint64(iterator.Key()[:8])
			addr = iterator.Key()[8:]
		)

		key := types.QuotaKey(id, addr)

		parent.Set(key, iterator.Value())
		child.Delete(iterator.Key())
	}

	return nil
}
