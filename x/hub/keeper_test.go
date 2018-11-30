package hub

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func TestKeeper(t *testing.T) {
	multiStore, accKey, _, coinLockerKey := setupMultiStore()
	ctx := csdkTypes.NewContext(multiStore, abci.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	csdkTypes.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	RegisterCodec(cdc)

	accountKeeper := auth.NewAccountKeeper(cdc, accKey, auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	hubKeeper := NewBaseKeeper(cdc, coinLockerKey, bankKeeper)

	var err csdkTypes.Error
	var locker *sdkTypes.CoinLocker

	account1 := auth.NewBaseAccountWithAddress(accAddress1)

	if err := account1.SetCoins(csdkTypes.Coins{
		coin(100, "x"),
	}); err != nil {
		panic(err)
	}

	accountKeeper.SetAccount(ctx, &account1)

	err = hubKeeper.LockCoins(ctx, "locker_id_2", accAddress1, csdkTypes.Coins{coin(10, "x")})
	require.Nil(t, err)

	require.Equal(t, accountKeeper.GetAccount(ctx, accAddress1).GetCoins(), csdkTypes.Coins{coin(90, "x")})

	locker, err = hubKeeper.GetLocker(ctx, "locker_id_2")
	require.Nil(t, err)
	require.Equal(t, locker.Address, accAddress1)
	require.Equal(t, locker.Coins, csdkTypes.Coins{coin(10, "x")})
	require.Equal(t, locker.Status, sdkTypes.StatusLock)

	err = hubKeeper.ReleaseCoins(ctx, "locker_id_2")
	require.Nil(t, err)
	require.Equal(t, accountKeeper.GetAccount(ctx, accAddress1).GetCoins(), csdkTypes.Coins{coin(100, "x")})

	locker, err = hubKeeper.GetLocker(ctx, "locker_id_2")
	require.Nil(t, err)
	require.Equal(t, locker.Address, accAddress1)
	require.Equal(t, locker.Coins, csdkTypes.Coins{coin(10, "x")})
	require.Equal(t, locker.Status, sdkTypes.StatusRelease)

	err = hubKeeper.LockCoins(ctx, "locker_id_3", accAddress1, csdkTypes.Coins{coin(10, "unknown")})
	require.NotNil(t, err)

	locker, err = hubKeeper.GetLocker(ctx, "locker_id_3")
	require.Nil(t, locker)
	require.Nil(t, err)

	err = hubKeeper.LockCoins(ctx, "locker_id_4", accAddress1, csdkTypes.Coins{coin(6, "x")})
	require.Nil(t, err)

	err = hubKeeper.ReleaseCoinsToMany(ctx, "locker_id_4",
		[]csdkTypes.AccAddress{accAddress1, accAddress2, accAddress3},
		[]csdkTypes.Coins{{coin(2, "x")}, {coin(2, "x")}, {coin(2, "x")}})
	require.Nil(t, err)
	require.Equal(t, accountKeeper.GetAccount(ctx, accAddress1).GetCoins(), csdkTypes.Coins{coin(96, "x")})
	require.Equal(t, accountKeeper.GetAccount(ctx, accAddress2).GetCoins(), csdkTypes.Coins{coin(2, "x")})
	require.Equal(t, accountKeeper.GetAccount(ctx, accAddress3).GetCoins(), csdkTypes.Coins{coin(2, "x")})

	locker, err = hubKeeper.GetLocker(ctx, "locker_id_4")
	require.Nil(t, err)
	require.Equal(t, locker.Address, accAddress1)
	require.Equal(t, locker.Coins, csdkTypes.Coins{coin(6, "x")})
	require.Equal(t, locker.Status, sdkTypes.StatusRelease)
}
