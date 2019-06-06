package keeper

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
	vpnTypes "github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

var (
	testCoinPos      = csdkTypes.NewInt64Coin("stake", 10)
	testCoinNeg      = csdkTypes.Coin{Denom: "stake", Amount: csdkTypes.NewInt(-10)}
	testCoinZero     = csdkTypes.NewInt64Coin("stake", 0)
	testCoinsPos     = csdkTypes.Coins{testCoinPos}
	testCoinsNeg     = csdkTypes.Coins{testCoinNeg, csdkTypes.Coin{Denom: "stake", Amount: csdkTypes.NewInt(-100)}}
	testCoinsZero    = csdkTypes.Coins{testCoinZero, csdkTypes.NewInt64Coin("stake", 0)}
	testCoinsEmpty   = csdkTypes.Coins{}
	testCoinsNil     = csdkTypes.Coins(nil)
	testPrivKey1     = ed25519.GenPrivKey()
	testPrivKey2     = ed25519.GenPrivKey()
	testPubKey1      = testPrivKey1.PubKey()
	testPubKey2      = testPrivKey2.PubKey()
	testAddress1     = csdkTypes.AccAddress(testPubKey1.Address())
	testAddress2     = csdkTypes.AccAddress(testPubKey2.Address())
	testAddressEmpty = csdkTypes.AccAddress([]byte(""))
	testDepositPos   = types.Deposit{Address: testAddress1, Coins: testCoinsPos}
	testDepositZero  = types.Deposit{Address: testAddress1, Coins: testCoinsZero}
	testDepositEmpty = types.Deposit{}
	testDepositNil   = types.Deposit{Coins: csdkTypes.Coins(nil)}
	testDepositsPos  = []types.Deposit{testDepositPos}
	testDepositsZero = []types.Deposit{testDepositZero}
	testDepositsNil  = []types.Deposit(nil)
)

func testCreateInput() (csdkTypes.Context, Keeper, bank.BaseKeeper) {
	keyDeposits := csdkTypes.NewKVStoreKey("deposits")
	keyAccount := csdkTypes.NewKVStoreKey("acc")
	keyParams := csdkTypes.NewKVStoreKey("params")
	tkeyParams := csdkTypes.NewTransientStoreKey("tparams")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyDeposits, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAccount, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, csdkTypes.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := testMakeCodec()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{ChainID: "chain-id"}, false, log.NewNopLogger())

	paramsKeeper := params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAccount, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	depositKeeper := NewKeeper(cdc, keyDeposits, bankKeeper)

	return ctx, depositKeeper, bankKeeper
}

func testMakeCodec() *codec.Codec {
	var cdc = codec.New()
	vpnTypes.RegisterCodec(cdc)
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
