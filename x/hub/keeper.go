package hub

import (
	"encoding/json"

	ccsdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	csdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper interface {
	GetLocker(ctx ccsdkTypes.Context, lockerId string) *csdkTypes.CoinLocker

	LockCoins(ctx ccsdkTypes.Context, lockerId string, address ccsdkTypes.AccAddress, coins ccsdkTypes.Coins)
	ReleaseCoins(ctx ccsdkTypes.Context, lockerId string)
	ReleaseCoinsToMany(ctx ccsdkTypes.Context, lockerId string, addresses []ccsdkTypes.AccAddress, shares []ccsdkTypes.Coins)
}

var _ Keeper = (*BaseKeeper)(nil)

type BaseKeeper struct {
	coinLockerKey ccsdkTypes.StoreKey
	bankKeeper    bank.Keeper
}

func NewBaseKeeper(coinLockerKey ccsdkTypes.StoreKey, bankKeeper bank.Keeper) BaseKeeper {
	return BaseKeeper{
		coinLockerKey: coinLockerKey,
		bankKeeper:    bankKeeper,
	}
}

func (k BaseKeeper) SetLocker(ctx ccsdkTypes.Context, lockerId string, locker *csdkTypes.CoinLocker) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockerId)
	valueBytes, err := json.Marshal(&locker)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k BaseKeeper) GetLocker(ctx ccsdkTypes.Context, lockerId string) *csdkTypes.CoinLocker {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockerId)
	valueBytes := store.Get(keyBytes)

	var locker csdkTypes.CoinLocker

	if err := json.Unmarshal(valueBytes, &locker); err != nil {
		panic(err)
	}

	return &locker
}

func (k BaseKeeper) LockCoins(ctx ccsdkTypes.Context, lockerId string,
	address ccsdkTypes.AccAddress, coins ccsdkTypes.Coins) {
	_, _, err := k.bankKeeper.SubtractCoins(ctx, address, coins)

	if err != nil {
		panic(err)
	}

	locker := csdkTypes.CoinLocker{
		Address: address,
		Coins:   coins,
		Locked:  true,
	}

	k.SetLocker(ctx, lockerId, &locker)
}

func (k BaseKeeper) ReleaseCoins(ctx ccsdkTypes.Context, lockerId string) {
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

func (k BaseKeeper) ReleaseCoinsToMany(ctx ccsdkTypes.Context, lockerId string,
	addresses []ccsdkTypes.AccAddress, shares []ccsdkTypes.Coins) {
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
