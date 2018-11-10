package hub

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type Keeper interface {
	GetLocker(ctx sdkTypes.Context, lockerId string) *hubTypes.CoinLocker

	LockCoins(ctx sdkTypes.Context, lockerId string, address sdkTypes.AccAddress, coins sdkTypes.Coins)
	ReleaseCoins(ctx sdkTypes.Context, lockerId string)
	ReleaseCoinsToMany(ctx sdkTypes.Context, lockerId string, addresses []sdkTypes.AccAddress, shares []sdkTypes.Coins)
}

var _ Keeper = (*BaseKeeper)(nil)

type BaseKeeper struct {
	coinLockerKey sdkTypes.StoreKey
	bankKeeper    bank.Keeper
}

func NewBaseKeeper(coinLockerKey sdkTypes.StoreKey, bankKeeper bank.Keeper) BaseKeeper {
	return BaseKeeper{
		coinLockerKey: coinLockerKey,
		bankKeeper:    bankKeeper,
	}
}

func (k BaseKeeper) SetLocker(ctx sdkTypes.Context, lockerId string, locker *hubTypes.CoinLocker) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockerId)
	valueBytes, err := json.Marshal(&locker)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k BaseKeeper) GetLocker(ctx sdkTypes.Context, lockerId string) *hubTypes.CoinLocker {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockerId)
	valueBytes := store.Get(keyBytes)

	var locker hubTypes.CoinLocker

	if err := json.Unmarshal(valueBytes, &locker); err != nil {
		panic(err)
	}

	return &locker
}

func (k BaseKeeper) LockCoins(ctx sdkTypes.Context, lockerId string,
	address sdkTypes.AccAddress, coins sdkTypes.Coins) {
	_, _, err := k.bankKeeper.SubtractCoins(ctx, address, coins)

	if err != nil {
		panic(err)
	}

	locker := hubTypes.CoinLocker{
		Address: address,
		Coins:   coins,
		Locked:  true,
	}

	k.SetLocker(ctx, lockerId, &locker)
}

func (k BaseKeeper) ReleaseCoins(ctx sdkTypes.Context, lockerId string) {
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

func (k BaseKeeper) ReleaseCoinsToMany(ctx sdkTypes.Context, lockerId string,
	addresses []sdkTypes.AccAddress, shares []sdkTypes.Coins) {
	locker := k.GetLocker(ctx, lockerId)

	for i := range addresses {
		_, _, err := k.bankKeeper.AddCoins(ctx, addresses[i], shares[i])

		if err != nil {
			panic(err)
		}
	}

	locker.Locked = false
	k.SetLocker(ctx, lockerId, locker)
}
