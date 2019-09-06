package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit/types"
)

func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	key := types.DepositKey(deposit.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(deposit)

	store := ctx.KVStore(k.key)
	store.Set(key, value)
}

func (k Keeper) GetDeposit(ctx sdk.Context, address sdk.AccAddress) (deposit types.Deposit, found bool) {
	store := ctx.KVStore(k.key)

	key := types.DepositKey(address)
	value := store.Get(key)
	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &deposit)
	return deposit, true
}

func (k Keeper) GetAllDeposits(ctx sdk.Context) (deposits []types.Deposit) {
	store := ctx.KVStore(k.key)

	iter := sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &deposit)
		deposits = append(deposits, deposit)
	}

	return deposits
}

func (k Keeper) Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) (err sdk.Error) {
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
		return types.ErrorInsufficientDepositFunds(deposit.Coins, coins)
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) (err sdk.Error) {
	deposit, found := k.GetDeposit(ctx, address)
	if !found {
		return types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	if err := k.supply.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) SendFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	deposit, found := k.GetDeposit(ctx, from)
	if !found {
		return types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	if err := k.supply.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) ReceiveFromAccountToDeposit(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
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
		return types.ErrorInsufficientDepositFunds(deposit.Coins, coins)
	}

	k.SetDeposit(ctx, deposit)
	return nil
}

func (k Keeper) IterateDeposits(ctx sdk.Context, fn func(index int64, deposit types.Deposit) (stop bool)) {
	store := ctx.KVStore(k.key)

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
