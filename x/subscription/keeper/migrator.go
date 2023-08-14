package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	v1types "github.com/sentinel-official/hub/x/subscription/legacy/v1/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := k.deleteSubscriptionForNodeKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteSubscriptionForPlanKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteActiveSubscriptionForAddressKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveSubscriptionForAddressKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveSubscriptionAtKeys(ctx); err != nil {
		return err
	}

	if err := k.setParams(ctx); err != nil {
		return err
	}
	if err := k.migrateSubscriptions(ctx); err != nil {
		return err
	}
	if err := k.migrateQuotas(ctx); err != nil {
		return err
	}
	if err := k.updateAllocations(ctx); err != nil {
		return err
	}

	return nil
}

func (k Migrator) deleteKeys(ctx sdk.Context, keyPrefix []byte) error {
	store := prefix.NewStore(k.Store(ctx), keyPrefix)

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}

	return nil
}

func (k Migrator) deleteSubscriptionForNodeKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x11})
}

func (k Migrator) deleteSubscriptionForPlanKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x12})
}

func (k Migrator) deleteActiveSubscriptionForAddressKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x20})
}

func (k Migrator) deleteInactiveSubscriptionForAddressKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x21})
}

func (k Migrator) deleteInactiveSubscriptionAtKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x30})
}

func (k Migrator) migrateSubscriptions(ctx sdk.Context) error {
	statusChangeDelay := k.StatusChangeDelay(ctx)

	store := prefix.NewStore(k.Store(ctx), []byte{0x10})

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var value v1types.Subscription
		k.cdc.MustUnmarshal(iter.Value(), &value)
		store.Delete(iter.Key())

		if value.Plan != 0 {
			subscription := &types.PlanSubscription{
				BaseSubscription: &types.BaseSubscription{
					ID:         value.Id,
					Address:    value.Owner,
					InactiveAt: value.Expiry,
					Status:     value.Status,
					StatusAt:   value.StatusAt,
				},
				PlanID: value.Plan,
				Denom:  value.Denom,
			}

			if value.Status.Equal(hubtypes.StatusInactivePending) {
				subscription.SetInactiveAt(value.StatusAt.Add(statusChangeDelay))
			}

			k.SetSubscription(ctx, subscription)
			k.SetSubscriptionForAccount(ctx, subscription.GetAddress(), subscription.GetID())
			k.SetSubscriptionForPlan(ctx, subscription.PlanID, subscription.GetID())
			k.SetSubscriptionForInactiveAt(ctx, subscription.GetInactiveAt(), subscription.GetID())
		}
	}

	return nil
}

func (k Migrator) migrateQuotas(ctx sdk.Context) error {
	store := prefix.NewStore(k.Store(ctx), []byte{0x40})

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var value v1types.Quota
		k.cdc.MustUnmarshal(iter.Value(), &value)

		id := sdk.BigEndianToUint64(iter.Key()[:8])
		store.Delete(iter.Key())

		if _, found := k.GetSubscription(ctx, id); !found {
			continue
		}

		alloc := types.Allocation{
			ID:            id,
			Address:       value.Address,
			GrantedBytes:  value.Allocated,
			UtilisedBytes: value.Consumed,
		}

		k.SetAllocation(ctx, alloc)
	}

	return nil
}

func (k Migrator) updateAllocations(ctx sdk.Context) error {
	k.IterateSubscriptions(ctx, func(_ int, item types.Subscription) (stop bool) {
		subscription, ok := item.(*types.PlanSubscription)
		if !ok {
			panic(fmt.Errorf("invalid subscription type %T", item))
		}

		grantedBytes := sdk.ZeroInt()
		k.IterateAllocationsForSubscription(ctx, item.GetID(), func(_ int, item types.Allocation) (stop bool) {
			grantedBytes = grantedBytes.Add(item.GrantedBytes)
			return false
		})

		plan, found := k.GetPlan(ctx, subscription.PlanID)
		if !found {
			panic(fmt.Errorf("plan %d does not exist", subscription.PlanID))
		}

		accAddr := subscription.GetAddress()
		diffBytes := hubtypes.Gigabyte.MulRaw(plan.Gigabytes).Sub(grantedBytes)

		alloc, found := k.GetAllocation(ctx, subscription.GetID(), accAddr)
		if !found {
			panic(fmt.Errorf("allocation %d/%s does not exist", subscription.GetID(), accAddr))
		}

		alloc.GrantedBytes = alloc.GrantedBytes.Add(diffBytes)
		k.SetAllocation(ctx, alloc)

		return false
	})

	return nil
}

func (k Migrator) setParams(ctx sdk.Context) error {
	k.SetParams(
		ctx,
		types.Params{
			StatusChangeDelay: 4 * 60 * time.Minute,
		},
	)

	return nil
}
