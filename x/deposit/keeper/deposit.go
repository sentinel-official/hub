package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit/types"
)

func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	key := types.DepositKey(deposit.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(deposit)

	store := ctx.KVStore(k.storeKey)
	store.Set(key, value)
}

func (k Keeper) GetDeposit(ctx sdk.Context, address sdk.AccAddress) (deposit types.Deposit, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.DepositKey(address)
	value := store.Get(key)
	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &deposit)
	return deposit, true
}

func (k Keeper) GetAllDeposits(ctx sdk.Context) (deposits []types.Deposit) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &deposit)
		deposits = append(deposits, deposit)
	}

	return deposits
}

func (k Keeper) Add(ctx sdk.Context, address sdk.AccAddress,
	coins sdk.Coins) (tags sdk.Tags, err sdk.Error) {

	_, tags, err = k.bankKeeper.SubtractCoins(ctx, address, coins)
	if err != nil {
		return nil, err
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
		return nil, types.ErrorInsufficientDepositFunds(deposit.Coins, coins)
	}

	k.SetDeposit(ctx, deposit)
	return tags, nil
}

func (k Keeper) Subtract(ctx sdk.Context, address sdk.AccAddress,
	coins sdk.Coins) (tags sdk.Tags, err sdk.Error) {

	deposit, found := k.GetDeposit(ctx, address)
	if !found {
		return nil, types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return nil, types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	_, tags, err = k.bankKeeper.AddCoins(ctx, address, coins)
	if err != nil {
		return nil, err
	}

	k.SetDeposit(ctx, deposit)
	return tags, nil
}

func (k Keeper) Send(ctx sdk.Context, from, toAddress sdk.AccAddress,
	coins sdk.Coins) (tags sdk.Tags, err sdk.Error) {

	deposit, found := k.GetDeposit(ctx, from)
	if !found {
		return nil, types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return nil, types.ErrorInsufficientDepositFunds(coins, deposit.Coins)
	}

	_, tags, err = k.bankKeeper.AddCoins(ctx, toAddress, coins)
	if err != nil {
		return nil, err
	}

	k.SetDeposit(ctx, deposit)
	return tags, nil
}

func (k Keeper) Receive(ctx sdk.Context, fromAddress, to sdk.AccAddress,
	coins sdk.Coins) (tags sdk.Tags, err sdk.Error) {

	_, tags, err = k.bankKeeper.SubtractCoins(ctx, fromAddress, coins)
	if err != nil {
		return nil, err
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
		return nil, types.ErrorInsufficientDepositFunds(deposit.Coins, coins)
	}

	k.SetDeposit(ctx, deposit)
	return tags, nil
}

// nolint: dupl
func (k Keeper) IterateDeposits(ctx sdk.Context, fn func(index int64, deposit types.Deposit) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

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
