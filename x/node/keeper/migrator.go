package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	v1types "github.com/sentinel-official/hub/x/node/legacy/v1/types"
	"github.com/sentinel-official/hub/x/node/types"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := k.deleteActiveNodeKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveNodeKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteActiveNodeForProviderKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveNodeForProviderKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveNodeAtKeys(ctx); err != nil {
		return err
	}

	if err := k.migrateNodes(ctx); err != nil {
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

func (k Migrator) deleteActiveNodeKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x20})
}

func (k Migrator) deleteInactiveNodeKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x21})
}

func (k Migrator) deleteActiveNodeForProviderKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x30})
}

func (k Migrator) deleteInactiveNodeForProviderKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x31})
}

func (k Migrator) deleteInactiveNodeAtKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x41})
}

func (k Migrator) migrateNodes(ctx sdk.Context) error {
	activeDuration := k.ActiveDuration(ctx)

	store := prefix.NewStore(k.Store(ctx), []byte{0x10})

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var value v1types.Node
		k.cdc.MustUnmarshal(iter.Value(), &value)
		store.Delete(iter.Key())

		node := types.Node{
			Address:        value.Address,
			GigabytePrices: value.Price,
			HourlyPrices:   nil,
			RemoteURL:      value.RemoteURL,
			InactiveAt:     time.Time{},
			Status:         value.Status,
			StatusAt:       value.StatusAt,
		}

		if node.Status.Equal(hubtypes.StatusActive) {
			node.InactiveAt = node.StatusAt.Add(activeDuration)
			k.SetNodeForInactiveAt(ctx, node.InactiveAt, node.GetAddress())
		}

		k.SetNode(ctx, node)
	}

	return nil
}
