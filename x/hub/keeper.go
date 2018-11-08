package hub

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
	"github.com/ironman0x7b2/sentinel-hub/x/ibc"
)

type Keeper struct {
	coinLockerKey sdkTypes.StoreKey

	bankKeeper bank.Keeper
	ibcKeeper  ibc.Keeper
}

func NewKeeper(coinLockerKey sdkTypes.StoreKey, bankKeeper bank.Keeper) Keeper {
	return Keeper{
		coinLockerKey: coinLockerKey,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) SetLockedCoins(ctx sdkTypes.Context, lockId string, coins hubTypes.LockedCoins) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockId)
	valueBytes, err := json.Marshal(coins)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetLockedCoins(ctx sdkTypes.Context, lockId string) hubTypes.LockedCoins {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockId)
	valueBytes := store.Get(keyBytes)

	var lockedCoins hubTypes.LockedCoins

	if err := json.Unmarshal(valueBytes, &lockedCoins); err != nil {
		panic(err)
	}

	return lockedCoins
}

func (k Keeper) DeleteLockedCoins(ctx sdkTypes.Context, lockId string) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockId)
	store.Delete(keyBytes)
}

func (k Keeper) LockCoins(ctx sdkTypes.Context, lockId string, addr sdkTypes.AccAddress, coins sdkTypes.Coins) {
	_, _, err := k.bankKeeper.SubtractCoins(ctx, addr, coins)

	if err != nil {
		panic(err)
	}

	lockedCoins := hubTypes.LockedCoins{
		Address: addr,
		Coins:   coins,
	}

	k.SetLockedCoins(ctx, lockId, lockedCoins)
}

func (k Keeper) UnlockCoins(ctx sdkTypes.Context, lockId string) {
	lockedCoins := k.GetLockedCoins(ctx, lockId)
	addr := lockedCoins.Address
	coins := lockedCoins.Coins

	_, _, err := k.bankKeeper.AddCoins(ctx, addr, coins)

	if err != nil {
		panic(err)
	}

	k.DeleteLockedCoins(ctx, lockId)
}

func (k Keeper) SplitUnlockCoins(ctx sdkTypes.Context, lockId string, splits []hubTypes.LockedCoins) {
	for _, split := range splits {
		addr := split.Address
		coins := split.Coins
		_, _, err := k.bankKeeper.AddCoins(ctx, addr, coins)

		if err != nil {
			panic(err)
		}
	}

	k.DeleteLockedCoins(ctx, lockId)
}
