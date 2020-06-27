package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/deposit/types"
)

func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	key := types.DepositKey(deposit.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(deposit)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetDeposit(ctx sdk.Context, address sdk.AccAddress) (deposit types.Deposit, found bool) {
	store := k.Store(ctx)

	key := types.DepositKey(address)
	value := store.Get(key)
	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &deposit)
	return deposit, true
}

func (k Keeper) GetDeposits(ctx sdk.Context) (items types.Deposits) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Deposit
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k Keeper) Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) sdk.Error {
	if err := k.supply.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, coins); err != nil {
		return err
	}

	deposit, found := k.GetDeposit(ctx, address)
	if !found {
		deposit = types.Deposit{
			Address: address,
			Coins:   sdk.Coins{},
		}
	}

	deposit.Coins = deposit.Coins.Add(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds()
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) sdk.Error {
	deposit, found := k.GetDeposit(ctx, address)
	if !found {
		return types.ErrorDepositDoesNotExist()
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds()
	}

	if err := k.supply.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) SendCoinsFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	deposit, found := k.GetDeposit(ctx, from)
	if !found {
		return types.ErrorDepositDoesNotExist()
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds()
	}

	if err := k.supply.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) SendCoinsFromAccountToDeposit(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	if err := k.supply.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, coins); err != nil {
		return err
	}

	deposit, found := k.GetDeposit(ctx, to)
	if !found {
		deposit = types.Deposit{
			Address: to,
			Coins:   sdk.Coins{},
		}
	}

	deposit.Coins = deposit.Coins.Add(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds()
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) IterateDeposits(ctx sdk.Context, fn func(index int64, item types.Deposit) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iterator.Close()

	for i := int64(0); iterator.Valid(); iterator.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &deposit)

		if stop := fn(i, deposit); stop {
			break
		}
		i++
	}
}
