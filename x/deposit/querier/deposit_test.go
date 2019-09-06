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
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

var (
	privKey1     = ed25519.GenPrivKey()
	privKey2     = ed25519.GenPrivKey()
	pubKey1      = privKey1.PubKey()
	pubKey2      = privKey2.PubKey()
	addressEmpty = sdk.AccAddress([]byte(""))
	address1     = sdk.AccAddress(pubKey1.Address())
	address2     = sdk.AccAddress(pubKey2.Address())

	coinEmpty = sdk.Coin{}
	coinNeg   = sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-10)}
	coinZero  = sdk.NewInt64Coin("stake", 0)
	coinPos   = sdk.NewInt64Coin("stake", 10)

	coinsEmpty = sdk.Coins{}
	coinsNil   = sdk.Coins(nil)
	coinsNeg   = sdk.Coins{coinNeg}
	coinsZero  = sdk.Coins{coinZero}
	coinsPos   = sdk.Coins{coinPos}

	depositEmpty = types.Deposit{}
	depositNil   = types.Deposit{Coins: sdk.Coins(nil)}
	depositNeg   = types.Deposit{Address: address1, Coins: coinsNeg}
	depositZero  = types.Deposit{Address: address1, Coins: coinsZero}
	depositPos   = types.Deposit{Address: address1, Coins: coinsPos}

	depositsEmpty []types.Deposit
	depositsNil   = []types.Deposit(nil)
	depositsNeg   = []types.Deposit{depositNeg}
	depositsZero  = []types.Deposit{depositZero}
	depositsPos   = []types.Deposit{depositPos}
)

func createTestInput(t *testing.T, isCheckTx bool) (sdk.Context, keeper.Keeper, bank.Keeper) {
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	keyAccount := sdk.NewKVStoreKey(auth.StoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	keyDeposits := sdk.NewKVStoreKey(types.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)

	mdb := db.NewMemDB()
	ms := store.NewCommitMultiStore(mdb)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyAccount, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyDeposits, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, mdb)
	require.Nil(t, ms.LoadLatestVersion())

	depositAccount := supply.NewEmptyModuleAccount(types.ModuleName)
	blacklist := make(map[string]bool)
	blacklist[depositAccount.String()] = true
	accountPermissions := map[string][]string{
		types.ModuleName: nil,
	}

	cdc := makeTestCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, isCheckTx, log.NewNopLogger())

	pk := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	ak := auth.NewAccountKeeper(cdc, keyAccount, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bk := bank.NewBaseKeeper(ak, pk.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, blacklist)
	sk := supply.NewKeeper(cdc, keySupply, ak, bk, accountPermissions)
	dk := keeper.NewKeeper(cdc, keyDeposits, sk)

	sk.SetModuleAccount(ctx, depositAccount)

	return ctx, dk, bk
}

func makeTestCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)
	return cdc
}

func Test_queryDepositOfAddress(t *testing.T) {
	ctx, dk, _ := createTestInput(t, false)
	cdc := makeTestCodec()
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDepositOfAddress),
		Data: []byte{},
	}

	res, err := queryDepositOfAddress(ctx, req, dk)
	require.NotNil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams(addressEmpty))

	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	var deposit types.Deposit
	require.NotNil(t, cdc.UnmarshalJSON(res, &deposit))
	require.NotEqual(t, depositPos, deposit)
	dk.SetDeposit(ctx, depositPos)

	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams(addressEmpty))

	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)

	require.NotNil(t, cdc.UnmarshalJSON(res, &deposit))
	require.NotEqual(t, depositPos, deposit)

	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams(address1))
	require.Nil(t, err)

	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)

	cdc.MustUnmarshalJSON(res, &deposit)
	require.Nil(t, err)
	require.Equal(t, depositPos, deposit)

	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams(address2))
	require.Nil(t, err)

	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.NotNil(t, cdc.UnmarshalJSON(res, &deposit))
}

func Test_queryAllDeposits(t *testing.T) {
	ctx, dk, _ := createTestInput(t, false)
	cdc := makeTestCodec()

	res, err := queryAllDeposits(ctx, dk)
	require.Nil(t, err)
	require.Equal(t, []byte("null"), res)

	var deposits []types.Deposit
	cdc.MustUnmarshalJSON(res, &deposits)
	require.NotEqual(t, depositsPos, deposits)

	dk.SetDeposit(ctx, depositPos)

	res, err = queryAllDeposits(ctx, dk)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)

	cdc.MustUnmarshalJSON(res, &deposits)
	require.Equal(t, depositsPos, deposits)

	deposit := depositPos
	deposit.Address = address2
	dk.SetDeposit(ctx, deposit)

	res, err = queryAllDeposits(ctx, dk)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)

	cdc.MustUnmarshalJSON(res, &deposits)
	require.Nil(t, err)
	require.Len(t, deposits, 2)
}
