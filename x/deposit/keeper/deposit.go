package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/deposit/types"
)

// SetDeposit is for inserting a deposit into KVStore.
func (k *Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	key := types.DepositKey(deposit.GetAddress())
	value := k.cdc.MustMarshal(&deposit)

	store := k.Store(ctx)
	store.Set(key, value)
}

// GetDeposit is for getting the deposit of an address from KVStore.
func (k *Keeper) GetDeposit(ctx sdk.Context, address sdk.AccAddress) (deposit types.Deposit, found bool) {
	store := k.Store(ctx)

	key := types.DepositKey(address)
	value := store.Get(key)
	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshal(value, &deposit)
	return deposit, true
}

// GetDeposits is for getting the deposits from KVStore.
func (k *Keeper) GetDeposits(ctx sdk.Context, skip, limit int64) (items types.Deposits) {
	var (
		store = k.Store(ctx)
		iter  = hubtypes.NewPaginatedIterator(
			sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix),
		)
	)

	defer iter.Close()

	iter.Skip(skip)
	iter.Limit(limit, func(iter sdk.Iterator) {
		var item types.Deposit
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	})

	return items
}

// Add is for adding the amount to the deposit account from the bank account of an address.
func (k *Keeper) Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error {
	if err := k.SendCoinsFromAccountToDeposit(ctx, address, address, coins); err != nil {
		return err
	}

	return nil
}

// Subtract is for adding the amount to the bank account from the deposit account of an address.
func (k *Keeper) Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error {
	if err := k.SendCoinsFromDepositToAccount(ctx, address, address, coins); err != nil {
		return err
	}

	return nil
}

// SendCoinsFromAccountToDeposit is for sending the amount
// from the bank account of from address to the deposit account of to address.
func (k *Keeper) SendCoinsFromAccountToDeposit(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error {
	if err := k.bank.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, coins); err != nil {
		return err
	}

	deposit, found := k.GetDeposit(ctx, to)
	if !found {
		deposit = types.Deposit{
			Address: to.String(),
			Coins:   sdk.NewCoins(),
		}
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventAdd{
			Address:  deposit.Address,
			Previous: deposit.Coins,
			Coins:    coins,
		},
	)

	deposit.Coins = deposit.Coins.Add(coins...)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

// SendCoinsFromDepositToAccount is for sending the amount
// from the deposit account of from address to the bank account of to address.
func (k *Keeper) SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) error {
	deposit, found := k.GetDeposit(ctx, from)
	if !found {
		return types.ErrorDepositDoesNotExist
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventAdd{
			Address:  deposit.Address,
			Previous: deposit.Coins,
			Coins:    coins,
		},
	)

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds
	}

	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

// IterateDeposits is for iterating over all the deposits to perform an action.
func (k *Keeper) IterateDeposits(ctx sdk.Context, fn func(index int, item types.Deposit) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshal(iterator.Value(), &deposit)

		if stop := fn(i, deposit); stop {
			break
		}
		i++
	}
}
