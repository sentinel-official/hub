package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
)

func (k Keeper) SetDeposit(ctx csdkTypes.Context, deposit types.Deposit) {
	key := types.DepositKey(deposit.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(deposit)

	store := ctx.KVStore(k.storeKey)
	store.Set(key, value)
}

func (k Keeper) GetDeposit(ctx csdkTypes.Context, address csdkTypes.AccAddress) (deposit types.Deposit, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.DepositKey(address)
	value := store.Get(key)
	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &deposit)
	return deposit, true
}

func (k Keeper) GetAllDeposits(ctx csdkTypes.Context) (deposits []types.Deposit) {
	store := ctx.KVStore(k.storeKey)

	iter := csdkTypes.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &deposit)
		deposits = append(deposits, deposit)
	}

	return deposits
}

func (k Keeper) Add(ctx csdkTypes.Context, address csdkTypes.AccAddress,
	coins csdkTypes.Coins) (tags csdkTypes.Tags, err csdkTypes.Error) {

	_, tags, err = k.bankKeeper.SubtractCoins(ctx, address, coins)
	if err != nil {
		return nil, err
	}

	deposit, found := k.GetDeposit(ctx, address)
	if !found {
		deposit = types.Deposit{
			Address: address,
			Coins:   csdkTypes.Coins{},
		}
	}

	deposit.Coins = deposit.Coins.Add(coins)
	if deposit.Coins.IsAnyNegative() {
		return nil, types.ErrorInsufficientDepositFunds(deposit.Coins, coins)
	}

	k.SetDeposit(ctx, deposit)
	return tags, nil
}

func (k Keeper) Subtract(ctx csdkTypes.Context, address csdkTypes.AccAddress,
	coins csdkTypes.Coins) (tags csdkTypes.Tags, err csdkTypes.Error) {

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

func (k Keeper) Send(ctx csdkTypes.Context, from, toAddress csdkTypes.AccAddress,
	coins csdkTypes.Coins) (tags csdkTypes.Tags, err csdkTypes.Error) {

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

func (k Keeper) Receive(ctx csdkTypes.Context, fromAddress, to csdkTypes.AccAddress,
	coins csdkTypes.Coins) (tags csdkTypes.Tags, err csdkTypes.Error) {

	_, tags, err = k.bankKeeper.SubtractCoins(ctx, fromAddress, coins)
	if err != nil {
		return nil, err
	}

	deposit, found := k.GetDeposit(ctx, to)
	if !found {
		deposit = types.Deposit{
			Address: to,
			Coins:   csdkTypes.Coins{},
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
func (k Keeper) IterateDeposits(ctx csdkTypes.Context, fn func(index int64, deposit types.Deposit) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := csdkTypes.KVStorePrefixIterator(store, types.DepositKeyPrefix)
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
