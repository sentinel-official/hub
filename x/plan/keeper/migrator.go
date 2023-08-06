package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	v1types "github.com/sentinel-official/hub/x/plan/legacy/v1/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := k.deleteActivePlanKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactivePlanKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteActivePlanForProviderKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactivePlanForProviderKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteCountForNodeByProviderKeys(ctx); err != nil {
		return err
	}

	if err := k.migratePlans(ctx); err != nil {
		return err
	}
	if err := k.migrateNodeForPlan(ctx); err != nil {
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

func (k Migrator) deleteActivePlanKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x20})
}

func (k Migrator) deleteInactivePlanKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x21})
}

func (k Migrator) deleteActivePlanForProviderKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x30})
}

func (k Migrator) deleteInactivePlanForProviderKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x31})
}

func (k Migrator) deleteCountForNodeByProviderKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x50})
}

func (k Migrator) migratePlans(ctx sdk.Context) error {
	store := prefix.NewStore(k.Store(ctx), []byte{0x10})

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var value v1types.Plan
		k.cdc.MustUnmarshal(iter.Value(), &value)
		store.Delete(iter.Key())

		plan := types.Plan{
			ID:              value.Id,
			ProviderAddress: value.Provider,
			Duration:        value.Validity,
			Gigabytes:       value.Bytes.ToDec().QuoInt(hubtypes.Gigabyte).Ceil().TruncateInt64(),
			Prices:          value.Price,
			Status:          value.Status,
			StatusAt:        value.StatusAt,
		}

		k.SetPlan(ctx, plan)
		k.SetPlanForProvider(ctx, plan.GetProviderAddress(), plan.ID)
	}

	return nil
}

func (k Migrator) migrateNodeForPlan(ctx sdk.Context) error {
	store := prefix.NewStore(k.Store(ctx), []byte{0x40})

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		id := sdk.BigEndianToUint64(iter.Key()[:8])
		nodeAddr := iter.Key()[9:]

		store.Delete(iter.Key())
		k.SetNodeForPlan(ctx, id, nodeAddr)
	}

	return nil
}
