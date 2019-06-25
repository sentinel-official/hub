package querier

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/sentinel-hub/x/deposit/keeper"
	"github.com/sentinel-official/sentinel-hub/x/deposit/types"
)

var (
	testPrivKey1     = ed25519.GenPrivKey()
	testPrivKey2     = ed25519.GenPrivKey()
	testPubKey1      = testPrivKey1.PubKey()
	testPubKey2      = testPrivKey2.PubKey()
	testAddressEmpty = sdk.AccAddress([]byte(""))
	testAddress1     = sdk.AccAddress(testPubKey1.Address())
	testAddress2     = sdk.AccAddress(testPubKey2.Address())

	testCoinPos  = sdk.NewInt64Coin("stake", 10)
	testCoinsPos = sdk.Coins{testCoinPos}

	testDepositPos  = types.Deposit{Address: testAddress1, Coins: testCoinsPos}
	testDepositsPos = []types.Deposit{testDepositPos}
)

func testCreateInput() (sdk.Context, keeper.Keeper, bank.BaseKeeper) {
	keyParams := sdk.NewKVStoreKey("params")
	keyAccount := sdk.NewKVStoreKey("acc")
	keyDeposits := sdk.NewKVStoreKey("deposits")
	tkeyParams := sdk.NewTransientStoreKey("tparams")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAccount, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyDeposits, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := testMakeCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, false, log.NewNopLogger())

	paramsKeeper := params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAccount, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	depositKeeper := keeper.NewKeeper(cdc, keyDeposits, bankKeeper)

	return ctx, depositKeeper, bankKeeper
}

func testMakeCodec() *codec.Codec {
	var cdc = codec.New()
	auth.RegisterBaseAccount(cdc)
	return cdc
}

var (
	TestDepositOfAddressParamsEmpty = QueryDepositOfAddressPrams{testAddressEmpty}
	TestDepositOfAddressParams1     = QueryDepositOfAddressPrams{testAddress1}
	TestDepositOfAddressParams2     = QueryDepositOfAddressPrams{testAddress2}
)

func TestNewQueryDepositOfAddressParams(t *testing.T) {
	params_ := NewQueryDepositOfAddressParams(testAddressEmpty)
	require.Equal(t, TestDepositOfAddressParamsEmpty, params_)

	params_ = NewQueryDepositOfAddressParams(testAddress1)
	require.Equal(t, TestDepositOfAddressParams1, params_)

	params_ = NewQueryDepositOfAddressParams(testAddress2)
	require.Equal(t, TestDepositOfAddressParams2, params_)
}

func Test_queryDepositOfAddress(t *testing.T) {
	ctx, depositKeeper, _ := testCreateInput()
	cdc := testMakeCodec()
	var err error
	var deposit types.Deposit

	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, QueryDepositOfAddress),
		Data: []byte{},
	}

	res, _err := queryDepositOfAddress(ctx, cdc, req, depositKeeper)
	require.NotNil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	req.Data, err = cdc.MarshalJSON(NewQueryDepositOfAddressParams(testAddressEmpty))
	require.Nil(t, err)

	res, _err = queryDepositOfAddress(ctx, cdc, req, depositKeeper)
	require.Nil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	err = cdc.UnmarshalJSON(res, &deposit)
	require.NotNil(t, err)
	require.NotEqual(t, testDepositPos, deposit)

	depositKeeper.SetDeposit(ctx, testDepositPos)

	req.Data, err = cdc.MarshalJSON(NewQueryDepositOfAddressParams(testAddressEmpty))
	require.Nil(t, err)

	res, _err = queryDepositOfAddress(ctx, cdc, req, depositKeeper)
	require.Nil(t, _err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	err = cdc.UnmarshalJSON(res, &deposit)
	require.NotNil(t, err)
	require.NotEqual(t, testDepositPos, deposit)

	req.Data, err = cdc.MarshalJSON(NewQueryDepositOfAddressParams(testAddress1))
	require.Nil(t, err)

	res, _err = queryDepositOfAddress(ctx, cdc, req, depositKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &deposit)
	require.Nil(t, err)
	require.Equal(t, testDepositPos, deposit)

	req.Data, err = cdc.MarshalJSON(NewQueryDepositOfAddressParams(testAddress2))
	require.Nil(t, err)

	res, _err = queryDepositOfAddress(ctx, cdc, req, depositKeeper)
	require.Nil(t, _err)
	require.Equal(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &deposit)
	require.NotNil(t, err)
}

func Test_queryAllDeposits(t *testing.T) {
	ctx, depositKeeper, _ := testCreateInput()
	cdc := testMakeCodec()
	var err error
	var deposits []types.Deposit

	res, _err := queryAllDeposits(ctx, cdc, depositKeeper)
	require.Nil(t, _err)
	require.Equal(t, []byte("null"), res)

	err = cdc.UnmarshalJSON(res, &deposits)
	require.Nil(t, err)
	require.NotEqual(t, testDepositsPos, deposits)

	depositKeeper.SetDeposit(ctx, testDepositPos)

	res, _err = queryAllDeposits(ctx, cdc, depositKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &deposits)
	require.Nil(t, err)
	require.Equal(t, testDepositsPos, deposits)

	deposit := testDepositPos
	deposit.Address = testAddress2

	depositKeeper.SetDeposit(ctx, deposit)

	res, _err = queryAllDeposits(ctx, cdc, depositKeeper)
	require.Nil(t, _err)
	require.NotEqual(t, []byte(nil), res)

	err = cdc.UnmarshalJSON(res, &deposits)
	require.Nil(t, err)
	require.Len(t, deposits, 2)
}
