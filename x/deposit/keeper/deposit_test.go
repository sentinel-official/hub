package keeper

import (
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
	testPrivKey1     = ed25519.GenPrivKey()
	testPrivKey2     = ed25519.GenPrivKey()
	testPubKey1      = testPrivKey1.PubKey()
	testPubKey2      = testPrivKey2.PubKey()
	testAddressEmpty = csdk.AccAddress([]byte(""))
	testAddress1     = csdk.AccAddress(testPubKey1.Address())
	testAddress2     = csdk.AccAddress(testPubKey2.Address())

	testCoinEmpty = csdk.Coin{}
	testCoinNeg   = csdk.Coin{Denom: "stake", Amount: csdk.NewInt(-10)}
	testCoinZero  = csdk.NewInt64Coin("stake", 0)
	testCoinPos   = csdk.NewInt64Coin("stake", 10)

	testCoinsEmpty = csdk.Coins{}
	testCoinsNil   = csdk.Coins(nil)
	testCoinsNeg   = csdk.Coins{testCoinNeg}
	testCoinsZero  = csdk.Coins{testCoinZero}
	testCoinsPos   = csdk.Coins{testCoinPos}

	testDepositEmpty = types.Deposit{}
	testDepositNil   = types.Deposit{Coins: csdk.Coins(nil)}
	testDepositNeg   = types.Deposit{Address: testAddress1, Coins: testCoinsNeg}
	testDepositZero  = types.Deposit{Address: testAddress1, Coins: testCoinsZero}
	testDepositPos   = types.Deposit{Address: testAddress1, Coins: testCoinsPos}

	testDepositsEmpty []types.Deposit
	testDepositsNil   = []types.Deposit(nil)
	testDepositsNeg   = []types.Deposit{testDepositNeg}
	testDepositsZero  = []types.Deposit{testDepositZero}
	testDepositsPos   = []types.Deposit{testDepositPos}
)

func testCreateInput() (csdk.Context, Keeper, bank.BaseKeeper) {
	keyParams := csdk.NewKVStoreKey("params")
	keyAccount := csdk.NewKVStoreKey("acc")
	keyDeposits := csdk.NewKVStoreKey("deposits")
	tkeyParams := csdk.NewTransientStoreKey("tparams")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyParams, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAccount, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyDeposits, csdk.StoreTypeIAVL, db)
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

	deposit, found := depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, false, found)
	require.Equal(t, testDepositEmpty, deposit)

	depositKeeper.SetDeposit(ctx, testDepositEmpty)
	deposit, found = depositKeeper.GetDeposit(ctx, testDepositEmpty.Address)
	require.Equal(t, true, found)
	require.Equal(t, testDepositEmpty, deposit)

	depositKeeper.SetDeposit(ctx, testDepositNil)
	deposit, found = depositKeeper.GetDeposit(ctx, testDepositNil.Address)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	depositKeeper.SetDeposit(ctx, testDepositNeg)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNeg, deposit)

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
	TestKeeper_SetDeposit(t)
}

func TestKeeper_GetAllDeposits(t *testing.T) {
	ctx, depositKeeper, _ := testCreateInput()

	deposits := depositKeeper.GetAllDeposits(ctx)
	require.Len(t, deposits, 0)
	require.Equal(t, testDepositsEmpty, deposits)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	deposits = depositKeeper.GetAllDeposits(ctx)
	require.Len(t, deposits, 1)
	require.Equal(t, testDepositsPos, deposits)

	testDepositPos2 := types.Deposit{Address: testAddress2, Coins: testCoinsPos}
	depositKeeper.SetDeposit(ctx, testDepositPos2)
	deposits = depositKeeper.GetAllDeposits(ctx)
	require.Len(t, deposits, 2)
}

func TestKeeper_Add(t *testing.T) {
	ctx, depositKeeper, bankKeeper := testCreateInput()

	testDepositNil.Address = testAddress1

	_, err := depositKeeper.Add(ctx, testAddress1, testCoinsEmpty)
	require.Nil(t, err)
	deposit, found := depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsNil)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsNeg)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsZero)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsPos)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositNil, deposit)

	coins, _, err := bankKeeper.AddCoins(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	require.Equal(t, testCoinsPos, coins)
	require.Equal(t, testCoinsPos, bankKeeper.GetCoins(ctx, testAddress1))

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)
	require.Equal(t, testCoinsNil, bankKeeper.GetCoins(ctx, testAddress1))

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsPos)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)
	require.Equal(t, testCoinsNil, bankKeeper.GetCoins(ctx, testAddress1))

	testCoinsPos2 := testCoinsPos.Add(testCoinsPos)
	testDepositPos2 := types.Deposit{Address: testAddress1, Coins: testCoinsPos2}

	coins, _, err = bankKeeper.AddCoins(ctx, testAddress1, testCoinsPos2)
	require.Nil(t, err)
	require.Equal(t, testCoinsPos2, coins)

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos2, deposit)
	require.Equal(t, testCoinsPos, bankKeeper.GetCoins(ctx, testAddress1))

	testCoinsPos3 := testCoinsPos2.Add(testCoinsPos)
	testDepositPos3 := types.Deposit{Address: testAddress1, Coins: testCoinsPos3}

	_, err = depositKeeper.Add(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)

	require.Equal(t, testDepositPos3, deposit)
	require.Equal(t, testCoinsNil, bankKeeper.GetCoins(ctx, testAddress1))
}

func TestKeeper_Subtract(t *testing.T) {
	ctx, depositKeeper, bankKeeper := testCreateInput()
	coins := bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsEmpty, coins)

	deposit, found := depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, false, found)
	require.Equal(t, testDepositEmpty, deposit)

	_, err := depositKeeper.Subtract(ctx, testAddress1, testCoinsEmpty)
	require.NotNil(t, err)
	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsNil)
	require.NotNil(t, err)
	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsNeg)
	require.NotNil(t, err)
	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsZero)
	require.NotNil(t, err)
	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsPos)
	require.NotNil(t, err)

	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, false, found)
	require.Equal(t, testDepositEmpty, deposit)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsEmpty)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsNil)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsNeg)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsZero)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsPos)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)

	deposit2 := testDepositPos
	deposit2.Coins = testCoinsPos.Add(testCoinsPos)
	depositKeeper.SetDeposit(ctx, deposit2)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsPos.Add(testCoinsPos), deposit.Coins)

	_, err = depositKeeper.Subtract(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsPos, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos.Add(testCoinsPos), coins)
}

func TestKeeper_Send(t *testing.T) {
	ctx, depositKeeper, bankKeeper := testCreateInput()

	coins := bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsEmpty, coins)

	deposit, found := depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, false, found)
	require.Equal(t, testDepositEmpty, deposit)

	_, err := depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsEmpty)
	require.NotNil(t, err)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsNil)
	require.NotNil(t, err)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsNeg)
	require.NotNil(t, err)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsZero)
	require.NotNil(t, err)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)

	depositKeeper.SetDeposit(ctx, testDepositPos)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsEmpty)
	require.Nil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsNil, coins)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsNil)
	require.Nil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsNil, coins)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsNeg)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsNil, coins)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsZero)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsNil, coins)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)

	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.Nil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsPos, coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsEmpty, coins)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)

	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsPos, coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsEmpty, coins)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)

	deposit = testDepositPos
	deposit.Coins = testCoinsPos.Add(testCoinsPos)
	depositKeeper.SetDeposit(ctx, deposit)
	_, err = depositKeeper.Send(ctx, testAddress1, testAddress2, testCoinsPos)
	require.Nil(t, err)
	coins = bankKeeper.GetCoins(ctx, testAddress2)
	require.Equal(t, testCoinsPos.Add(testCoinsPos), coins)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress1)
	require.Equal(t, true, found)
	require.Equal(t, testDepositPos, deposit)
}

func TestKeeper_Receive(t *testing.T) {
	ctx, depositKeeper, bankKeeper := testCreateInput()

	deposit, found := depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, false, found)
	coins := bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsEmpty, coins)

	_, err := depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsEmpty)
	require.Nil(t, err)
	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsNil)
	require.Nil(t, err)
	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsNeg)
	require.NotNil(t, err)
	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsZero)
	require.NotNil(t, err)
	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)

	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
	coins, _, err = bankKeeper.AddCoins(ctx, testAddress1, testCoinsPos)
	require.Nil(t, err)
	require.Equal(t, testCoinsPos, coins)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsEmpty)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsNil)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsNeg)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsZero)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsNil, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsPos)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsPos, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsNil, coins)

	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsPos)
	require.NotNil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsPos, deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsNil, coins)

	coins, _, err = bankKeeper.AddCoins(ctx, testAddress1, testCoinsPos.Add(testCoinsPos))
	require.Nil(t, err)
	require.Equal(t, testCoinsPos.Add(testCoinsPos), coins)
	_, err = depositKeeper.Receive(ctx, testAddress1, testAddress2, testCoinsPos)
	require.Nil(t, err)
	deposit, found = depositKeeper.GetDeposit(ctx, testAddress2)
	require.Equal(t, true, found)
	require.Equal(t, testCoinsPos.Add(testCoinsPos), deposit.Coins)
	coins = bankKeeper.GetCoins(ctx, testAddress1)
	require.Equal(t, testCoinsPos, coins)
}
