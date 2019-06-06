package keeper

import (
	csdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
)

func (k Keeper) SetDeposit(ctx csdk.Context, deposit types.Deposit) {
	key := types.DepositKey(deposit.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(deposit)

	store := ctx.KVStore(k.storeKey)
	store.Set(key, value)
}

func (k Keeper) GetDeposit(ctx csdk.Context, address csdk.AccAddress) (deposit types.Deposit, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.DepositKey(address)
	value := store.Get(key)
	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &deposit)
	return deposit, true
}

func (k Keeper) GetAllDeposits(ctx csdk.Context) (deposits []types.Deposit) {
	store := ctx.KVStore(k.storeKey)

	iter := csdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &deposit)
		deposits = append(deposits, deposit)
	}

	return deposits
}

func (k Keeper) Add(ctx csdk.Context, address csdk.AccAddress,
	coins csdk.Coins) (tags csdk.Tags, err csdk.Error) {

	_, tags, err = k.bankKeeper.SubtractCoins(ctx, address, coins)
	if err != nil {
		return nil, err
	}

	deposit, found := k.GetDeposit(ctx, address)
	if !found {
		deposit = types.Deposit{
			Address: address,
			Coins:   csdk.Coins{},
		}
	}

	deposit.Coins = deposit.Coins.Add(coins)
	if deposit.Coins.IsAnyNegative() {
		return nil, types.ErrorInsufficientDepositFunds(deposit.Coins, coins)
	}

	k.SetDeposit(ctx, deposit)
	return tags, nil
}

func (k Keeper) Subtract(ctx csdk.Context, address csdk.AccAddress,
	coins csdk.Coins) (tags csdk.Tags, err csdk.Error) {

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

func (k Keeper) Send(ctx csdk.Context, from, toAddress csdk.AccAddress,
	coins csdk.Coins) (tags csdk.Tags, err csdk.Error) {

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

func (k Keeper) Receive(ctx csdk.Context, fromAddress, to csdk.AccAddress,
	coins csdk.Coins) (tags csdk.Tags, err csdk.Error) {

	_, tags, err = k.bankKeeper.SubtractCoins(ctx, fromAddress, coins)
	if err != nil {
		return nil, err
	}

	deposit, found := k.GetDeposit(ctx, to)
	if !found {
		deposit = types.Deposit{
			Address: to,
			Coins:   csdk.Coins{},
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
func (k Keeper) IterateDeposits(ctx csdk.Context, fn func(index int64, deposit types.Deposit) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := csdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
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
