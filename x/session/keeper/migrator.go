package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := k.deleteSessions(ctx); err != nil {
		return err
	}
	if err := k.deleteActiveSessionForAddressKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveSessionForAddressKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveSessionAtKeys(ctx); err != nil {
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

func (k Migrator) deleteSessions(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x11})
}

func (k Migrator) deleteActiveSessionForAddressKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x30})
}

func (k Migrator) deleteInactiveSessionForAddressKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x31})
}

func (k Migrator) deleteInactiveSessionAtKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x40})
}
