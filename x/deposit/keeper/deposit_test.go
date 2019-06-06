package keeper

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
)

var (
	testCoinPos      = csdk.NewInt64Coin("stake", 10)
	testCoinNeg      = csdk.Coin{Denom: "stake", Amount: csdk.NewInt(-10)}
	testCoinZero     = csdk.NewInt64Coin("stake", 0)
	testCoinsPos     = csdk.Coins{testCoinPos}
	testCoinsNeg     = csdk.Coins{testCoinNeg, csdk.Coin{Denom: "stake", Amount: csdk.NewInt(-100)}}
	testCoinsZero    = csdk.Coins{testCoinZero, csdk.NewInt64Coin("stake", 0)}
	testCoinsEmpty   = csdk.Coins{}
	testCoinsNil     = csdk.Coins(nil)
	testPrivKey1     = ed25519.GenPrivKey()
	testPrivKey2     = ed25519.GenPrivKey()
	testPubKey1      = testPrivKey1.PubKey()
	testPubKey2      = testPrivKey2.PubKey()
	testAddress1     = csdk.AccAddress(testPubKey1.Address())
	testAddress2     = csdk.AccAddress(testPubKey2.Address())
	testAddressEmpty = csdk.AccAddress([]byte(""))
	testDepositPos   = types.Deposit{Address: testAddress1, Coins: testCoinsPos}
	testDepositZero  = types.Deposit{Address: testAddress1, Coins: testCoinsZero}
	testDepositEmpty = types.Deposit{}
	testDepositNil   = types.Deposit{Coins: csdk.Coins(nil)}
	testDepositsPos  = []types.Deposit{testDepositPos}
	testDepositsZero = []types.Deposit{testDepositZero}
	testDepositsNil  = []types.Deposit(nil)
)

func testCreateInput() (csdk.Context, Keeper, bank.BaseKeeper) {
	keyDeposits := csdk.NewKVStoreKey("deposits")
	keyAccount := csdk.NewKVStoreKey("acc")
	keyParams := csdk.NewKVStoreKey("params")
	tkeyParams := csdk.NewTransientStoreKey("tparams")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyDeposits, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAccount, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, csdk.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := testMakeCodec()
	ctx := csdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, false, log.NewNopLogger())

	paramsKeeper := params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAccount, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	depositKeeper := NewKeeper(cdc, keyDeposits, bankKeeper)

	return ctx, depositKeeper, bankKeeper
}

func testMakeCodec() *codec.Codec {
	var cdc = codec.New()
	auth.RegisterBaseAccount(cdc)
	return cdc
}

func TestKeeper_SetDeposit(t *testing.T) {
	ctx, depositKeeper, _ := testCreateInput()

	_, found := depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, false, found)

	depositKeeper.SetDeposit(ctx, testDepositEmpty)
	deposit, found := depositKeeper.GetDeposit(ctx, testDepositEmpty.Address)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	depositKeeper.SetDeposit(ctx, testDepositZero)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositZero, deposit)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	deposit, found = depositKeeper.GetDeposit(ctx, testDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)
}

func TestKeeper_GetDeposit(t *testing.T) {
	ctx, depositKeeper, _ := testCreateInput()

	_, found := depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, false, found)

	depositKeeper.SetDeposit(ctx, testDepositEmpty)
	deposit, found := depositKeeper.GetDeposit(ctx, testAddressEmpty)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	depositKeeper.SetDeposit(ctx, testDepositZero)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositZero, deposit)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	deposit, found = depositKeeper.GetDeposit(ctx, testDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)
}

func TestKeeper_GetAllDeposits(t *testing.T) {
	ctx, depositKeeper, _ := testCreateInput()

	deposits := depositKeeper.GetAllDeposits(ctx)
	require.Equal(t, testDepositsNil, deposits)

	depositKeeper.SetDeposit(ctx, testDepositZero)
	deposits = depositKeeper.GetAllDeposits(ctx)
	require.Equal(t, testDepositsZero, deposits)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	deposits = depositKeeper.GetAllDeposits(ctx)
	require.Equal(t, testDepositsPos, deposits)
}

func TestKeeper_Add(t *testing.T) {
	ctx, depositKeeper, bankKeeper := testCreateInput()

	coins, _, err := bankKeeper.AddCoins(ctx, testAddress1, testCoinsNeg)
	require.NotNil(t, err)

	coins, _, err = bankKeeper.AddCoins(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	require.Equal(t, testCoinsPos, coins)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	deposit, found := depositKeeper.GetDeposit(ctx, testDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Add(ctx, testDepositPos.Address, testCoinsPos)
	require.Nil(t, err)

	deposit, found = depositKeeper.GetDeposit(ctx, testDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsPos.Add(testCoinsPos), deposit.Coins)
}

func TestKeeper_Subtract(t *testing.T) {
	ctx, depositKeeper, _ := testCreateInput()

	_, err := depositKeeper.Subtract(ctx, testAddress1, testCoinsPos)
	require.NotNil(t, err)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	deposit, found := depositKeeper.GetDeposit(ctx, testDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsPos)
	require.NotNil(t, err)

	deposit, found = depositKeeper.GetDeposit(ctx, testDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
}

func TestKeeper_Send(t *testing.T) {
	ctx, depositKeeper, bankKeeper := testCreateInput()

	_, err := depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)

	coins, _, err := bankKeeper.AddCoins(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	require.Equal(t, testCoinsPos, coins)

	depositKeeper.SetDeposit(ctx, testDepositEmpty)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsEmpty, coins)

	depositKeeper.SetDeposit(ctx, testDepositZero)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsEmpty, coins)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.Nil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsPos, coins)

	deposit, found := depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)

	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsPos, coins)
}

func TestKeeper_Receive(t *testing.T) {
	ctx, depositKeeper, bankKeeper := testCreateInput()

	_, err := depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)
	deposit, found := depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, false, found)

	coins, _, err := bankKeeper.AddCoins(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	require.Equal(t, testCoinsPos, coins)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsPos.Add(testCoinsPos))
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, false, found)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsPos)
	require.Nil(t, err)

	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	reflect.DeepEqual(testDepositPos, deposit)
}
