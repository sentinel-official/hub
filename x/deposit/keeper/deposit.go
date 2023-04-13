package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit/types"
)

// SetDeposit is for inserting a deposit into the KVStore.
func (k *Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	var (
		store = k.Store(ctx)
		key   = types.DepositKey(deposit.GetAddress())
		value = k.cdc.MustMarshal(&deposit)
	)

	store.Set(key, value)
}

// GetDeposit is for getting a deposit of an address from the KVStore.
func (k *Keeper) GetDeposit(ctx sdk.Context, addr sdk.AccAddress) (deposit types.Deposit, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.DepositKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshal(value, &deposit)
	return deposit, true
}

// GetDeposits is for getting the deposits from the KVStore.
func (k *Keeper) GetDeposits(ctx sdk.Context) (items types.Deposits) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Deposit
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

// IterateDeposits is for iterating over all the deposits to perform an action.
func (k *Keeper) IterateDeposits(ctx sdk.Context, fn func(index int, item types.Deposit) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item types.Deposit
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SendCoinsFromAccountToDeposit is for sending an amount
// from a bank account of an address to a deposit account of an address.
func (k *Keeper) SendCoinsFromAccountToDeposit(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error {
	if err := k.bank.SendCoinsFromAccountToModule(ctx, fromAddr, types.ModuleName, coins); err != nil {
		return err
	}

	deposit, found := k.GetDeposit(ctx, toAddr)
	if !found {
		deposit = types.Deposit{
			Address: toAddr.String(),
			Coins:   sdk.NewCoins(),
		}
	}

	deposit.Coins = deposit.Coins.Add(coins...)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds
	}

	k.SetDeposit(ctx, deposit)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAdd{
			Address: toAddr.String(),
			Coins:   coins,
		},
	)

	return nil
}

// SendCoinsFromDepositToAccount is for sending an amount
// from a deposit account of an address to a bank account of an address.
func (k *Keeper) SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error {
	deposit, found := k.GetDeposit(ctx, fromAddr)
	if !found {
		return types.ErrorDepositDoesNotExist
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds
	}

	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubtract{
			Address: fromAddr.String(),
			Coins:   coins,
		},
	)

	return nil
}

// SendCoinsFromDepositToModule is for sending an amount
// from a deposit account of an address to a module account.
func (k *Keeper) SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error {
	deposit, found := k.GetDeposit(ctx, fromAddr)
	if !found {
		return types.ErrorDepositDoesNotExist
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds
	}

	if err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, toModule, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubtract{
			Address: fromAddr.String(),
			Coins:   coins,
		},
	)

	return nil
}

// Add is for adding an amount to a deposit account from a bank account of an address.
func (k *Keeper) Add(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	return k.SendCoinsFromAccountToDeposit(ctx, addr, addr, coins)
}

// Subtract is for adding an amount to a bank account from a deposit account of an address.
func (k *Keeper) Subtract(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	return k.SendCoinsFromDepositToAccount(ctx, addr, addr, coins)
}
