package hub

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper interface {
	GetLocker(ctx csdkTypes.Context, lockerID string) *sdkTypes.CoinLocker

	LockCoins(ctx csdkTypes.Context, lockerID string, address csdkTypes.AccAddress, coins csdkTypes.Coins) csdkTypes.Error
	ReleaseCoins(ctx csdkTypes.Context, lockerID string) csdkTypes.Error
	ReleaseCoinsToMany(ctx csdkTypes.Context, lockerID string, addresses []csdkTypes.AccAddress, shares []csdkTypes.Coins) csdkTypes.Error
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

func (k BaseKeeper) SetLocker(ctx csdkTypes.Context, lockerID string, locker *sdkTypes.CoinLocker) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(lockerID)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(locker)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k BaseKeeper) GetLocker(ctx csdkTypes.Context, lockerID string) *sdkTypes.CoinLocker {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(lockerID)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil
	}

	var locker sdkTypes.CoinLocker

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &locker); err != nil {
		panic(err)
	}

	return &locker
}

func (k BaseKeeper) LockCoins(ctx csdkTypes.Context, lockerID string,
	address csdkTypes.AccAddress, coins csdkTypes.Coins) csdkTypes.Error {
	_, _, err := k.bankKeeper.SubtractCoins(ctx, address, coins)

	if err != nil {
		return err
	}

	locker := sdkTypes.CoinLocker{
		Address: address,
		Coins:   coins,
		Status:  "LOCKED",
	}

	k.SetLocker(ctx, lockerID, &locker)

	return nil
}

func (k BaseKeeper) ReleaseCoins(ctx csdkTypes.Context, lockerID string) csdkTypes.Error {
	locker := k.GetLocker(ctx, lockerID)
	_, _, err := k.bankKeeper.AddCoins(ctx, locker.Address, locker.Coins)

	if err != nil {
		return err
	}

	locker.Status = "RELEASED"
	k.SetLocker(ctx, lockerID, locker)

	return nil
}

func (k BaseKeeper) ReleaseCoinsToMany(ctx csdkTypes.Context, lockerID string,
	addresses []csdkTypes.AccAddress, shares []csdkTypes.Coins) csdkTypes.Error {
	locker := k.GetLocker(ctx, lockerID)

	for i := range addresses {
		_, _, err := k.bankKeeper.AddCoins(ctx, addresses[i], shares[i])

		if err != nil {
			return err
		}
	}

	locker.Status = "RELEASED"
	k.SetLocker(ctx, lockerID, locker)

	return nil
}
