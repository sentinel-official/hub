package hub

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper interface {
	GetLocker(ctx csdkTypes.Context, lockerID string) (*sdkTypes.CoinLocker, csdkTypes.Error)

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

func (k BaseKeeper) SetLocker(ctx csdkTypes.Context, lockerID string, locker *sdkTypes.CoinLocker) csdkTypes.Error {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(lockerID)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(locker)

	if err != nil {
		return errorMarshal()
	}

	store.Set(keyBytes, valueBytes)

	return nil
}

func (k BaseKeeper) GetLocker(ctx csdkTypes.Context, lockerID string) (*sdkTypes.CoinLocker, csdkTypes.Error) {
	store := ctx.KVStore(k.coinLockerKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(lockerID)

	if err != nil {
		return nil, errorMarshal()
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil, nil
	}

	var locker sdkTypes.CoinLocker

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &locker); err != nil {
		return nil, errorUnmarshal()
	}

	return &locker, nil
}

func (k BaseKeeper) LockCoins(ctx csdkTypes.Context, lockerID string,
	address csdkTypes.AccAddress, coins csdkTypes.Coins) csdkTypes.Error {

	if _, _, err := k.bankKeeper.SubtractCoins(ctx, address, coins); err != nil {
		return err
	}

	locker := sdkTypes.CoinLocker{
		Address: address,
		Coins:   coins,
		Status:  "LOCKED",
	}

	if err := k.SetLocker(ctx, lockerID, &locker); err != nil {
		return err
	}

	return nil
}

func (k BaseKeeper) ReleaseCoins(ctx csdkTypes.Context, lockerID string) csdkTypes.Error {
	locker, err := k.GetLocker(ctx, lockerID)

	if err != nil {
		return err
	}

	if _, _, err := k.bankKeeper.AddCoins(ctx, locker.Address, locker.Coins); err != nil {
		return err
	}

	locker.Status = "RELEASED"

	if err := k.SetLocker(ctx, lockerID, locker); err != nil {
		return err
	}

	return nil
}

func (k BaseKeeper) ReleaseCoinsToMany(ctx csdkTypes.Context, lockerID string,
	addresses []csdkTypes.AccAddress, shares []csdkTypes.Coins) csdkTypes.Error {
	locker, err := k.GetLocker(ctx, lockerID)

	if err != nil {
		return err
	}

	for index := range addresses {
		if _, _, err := k.bankKeeper.AddCoins(ctx, addresses[index], shares[index]); err != nil {
			return err
		}
	}

	locker.Status = "RELEASED"

	if err := k.SetLocker(ctx, lockerID, locker); err != nil {
		return err
	}

	return nil
}
