package hub

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper interface {
	GetLocker(ctx csdkTypes.Context, lockerId string) *sdkTypes.CoinLocker

	LockCoins(ctx csdkTypes.Context, lockerId string, address csdkTypes.AccAddress, coins csdkTypes.Coins)
	ReleaseCoins(ctx csdkTypes.Context, lockerId string)
	ReleaseCoinsToMany(ctx csdkTypes.Context, lockerId string, addresses []csdkTypes.AccAddress, shares []csdkTypes.Coins)
}

var _ Keeper = (*BaseKeeper)(nil)

type BaseKeeper struct {
	cdc           *codec.Codec
	coinLockerKey csdkTypes.StoreKey
	bankKeeper    bank.Keeper
}

func NewBaseKeeper(cdc *codec.Codec, coinLockerKey csdkTypes.StoreKey, bankKeeper bank.Keeper) BaseKeeper {
	return BaseKeeper{
		cdc:           cdc,
		coinLockerKey: coinLockerKey,
		bankKeeper:    bankKeeper,
	}
}

func (k BaseKeeper) SetLocker(ctx csdkTypes.Context, lockerId string, locker *sdkTypes.CoinLocker) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := k.cdc.MustMarshalBinary(lockerId)
	valueBytes := k.cdc.MustMarshalBinary(&locker)
	store.Set(keyBytes, valueBytes)
}

func (k BaseKeeper) GetLocker(ctx csdkTypes.Context, lockerId string) *sdkTypes.CoinLocker {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes := []byte(lockerId)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil
	}

	var locker sdkTypes.CoinLocker

	k.cdc.MustUnmarshalBinary(valueBytes, &locker)

	return &locker
}

func (k BaseKeeper) LockCoins(ctx csdkTypes.Context, lockerId string,
	address csdkTypes.AccAddress, coins csdkTypes.Coins) {
	_, _, err := k.bankKeeper.SubtractCoins(ctx, address, coins)

	if err != nil {
		panic(err)
	}

	locker := sdkTypes.CoinLocker{
		Address: address,
		Coins:   coins,
		Locked:  true,
	}

	k.SetLocker(ctx, lockerId, &locker)
}

func (k BaseKeeper) ReleaseCoins(ctx csdkTypes.Context, lockerId string) {
	locker := k.GetLocker(ctx, lockerId)

	_, _, err := k.bankKeeper.AddCoins(ctx, locker.Address, locker.Coins)

	if err != nil {
		panic(err)
	}

	locker.Locked = false
	k.SetLocker(ctx, lockerId, locker)
}

func (k BaseKeeper) ReleaseCoinsToMany(ctx csdkTypes.Context, lockerId string,
	addresses []csdkTypes.AccAddress, shares []csdkTypes.Coins) {
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
