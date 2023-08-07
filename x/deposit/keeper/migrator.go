package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit/types"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := k.refund(ctx); err != nil {
		return err
	}

	return nil
}

func (k Migrator) refund(ctx sdk.Context) error {
	k.IterateDeposits(ctx, func(_ int, item types.Deposit) (stop bool) {
		if err := k.Subtract(ctx, item.GetAddress(), item.Coins); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}
