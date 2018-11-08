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

func (k Keeper) SetLocker(ctx sdkTypes.Context, lockerId string, locker *hubTypes.CoinLocker) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockerId)
	valueBytes, err := json.Marshal(&locker)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetLocker(ctx sdkTypes.Context, lockerId string) *hubTypes.CoinLocker {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockerId)
	valueBytes := store.Get(keyBytes)

	var locker hubTypes.CoinLocker

	if err := json.Unmarshal(valueBytes, &locker); err != nil {
		panic(err)
	}

	return &locker
}

func (k Keeper) LockCoins(ctx sdkTypes.Context, lockerId string, addr sdkTypes.AccAddress, coins sdkTypes.Coins) {
	_, _, err := k.bankKeeper.SubtractCoins(ctx, addr, coins)

	if err != nil {
		panic(err)
	}

	locker := hubTypes.CoinLocker{
		Address: addr,
		Coins:   coins,
		Locked:  true,
	}

	k.SetLocker(ctx, lockerId, &locker)
}

func (k Keeper) UnlockCoins(ctx sdkTypes.Context, lockerId string) {
	locker := k.GetLocker(ctx, lockerId)
	addr := locker.Address
	coins := locker.Coins

	_, _, err := k.bankKeeper.AddCoins(ctx, addr, coins)

	if err != nil {
		panic(err)
	}

	locker.Locked = false
	k.SetLocker(ctx, lockerId, locker)
}

func (k Keeper) UnlockAndShareCoins(ctx sdkTypes.Context, lockerId string, addrs []sdkTypes.AccAddress, shares []sdkTypes.Coins) {
	locker := k.GetLocker(ctx, lockerId)

	for i := range addrs {
		_, _, err := k.bankKeeper.AddCoins(ctx, addrs[i], shares[i])

		if err != nil {
			panic(err)
		}
	}

	locker.Locked = false
	k.SetLocker(ctx, lockerId, locker)
}
