package keeper_test

import (
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

func TestKeeper_SetDeposit(t *testing.T) {
	ctx, dk, _ := createTestInput(t, false)

	deposit, found := dk.GetDeposit(ctx, depositEmpty.Address)
	require.Equal(t, false, found)
	require.Equal(t, depositEmpty, deposit)
	dk.SetDeposit(ctx, depositEmpty)
	deposit, found = dk.GetDeposit(ctx, depositEmpty.Address)
	require.Equal(t, true, found)
	require.Equal(t, depositEmpty, deposit)

	dk.SetDeposit(ctx, depositNil)
	deposit, found = dk.GetDeposit(ctx, depositNil.Address)
	require.Equal(t, true, found)
	require.Equal(t, depositNil, deposit)

	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, depositEmpty, deposit)
	dk.SetDeposit(ctx, depositNeg)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositNeg, deposit)

	dk.SetDeposit(ctx, depositZero)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositZero, deposit)

	dk.SetDeposit(ctx, depositPos)
	deposit, found = dk.GetDeposit(ctx, depositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)
}

func TestKeeper_GetDeposit(t *testing.T) {
	TestKeeper_SetDeposit(t)
}

func TestKeeper_GetAllDeposits(t *testing.T) {
	ctx, dk, _ := createTestInput(t, false)

	deposits := dk.GetAllDeposits(ctx)
	require.Len(t, deposits, 0)
	require.Equal(t, depositsEmpty, deposits)

	dk.SetDeposit(ctx, depositPos)
	deposits = dk.GetAllDeposits(ctx)
	require.Len(t, deposits, 1)
	require.Equal(t, depositsPos, deposits)

	depositPos2 := types.Deposit{Address: address2, Coins: coinsPos}
	dk.SetDeposit(ctx, depositPos2)
	deposits = dk.GetAllDeposits(ctx)
	require.Len(t, deposits, 2)
}

func TestKeeper_Add(t *testing.T) {
	ctx, dk, bk := createTestInput(t, false)

	depositNil.Address = address1

	err := dk.Add(ctx, address1, coinsEmpty)
	require.Nil(t, err)
	deposit, found := dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositNil, deposit)

	err = dk.Add(ctx, address1, coinsNil)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositNil, deposit)

	err = dk.Add(ctx, address1, coinsNeg)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositNil, deposit)

	err = dk.Add(ctx, address1, coinsZero)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositNil, deposit)

	err = dk.Add(ctx, address1, coinsPos)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositNil, deposit)

	coins, err := bk.AddCoins(ctx, address1, coinsPos)
	require.Nil(t, err)
	require.Equal(t, coinsPos, coins)
	require.Equal(t, coinsPos, bk.GetCoins(ctx, address1))

	err = dk.Add(ctx, address1, coinsPos)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)
	require.Equal(t, coinsNil, bk.GetCoins(ctx, address1))

	err = dk.Add(ctx, address1, coinsPos)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)
	require.Equal(t, coinsNil, bk.GetCoins(ctx, address1))

	coinsPos2 := coinsPos.Add(coinsPos)
	depositPos2 := types.Deposit{Address: address1, Coins: coinsPos2}

	coins, err = bk.AddCoins(ctx, address1, coinsPos2)
	require.Nil(t, err)
	require.Equal(t, coinsPos2, bk.GetCoins(ctx, address1))

	err = dk.Add(ctx, address1, coinsPos)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos2, deposit)
	require.Equal(t, coinsPos, bk.GetCoins(ctx, address1))

	coinsPos3 := coinsPos2.Add(coinsPos)
	depositPos3 := types.Deposit{Address: address1, Coins: coinsPos3}

	err = dk.Add(ctx, address1, coinsPos)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos3, deposit)
	require.Equal(t, coinsNil, bk.GetCoins(ctx, address1))
}

func TestKeeper_Subtract(t *testing.T) {
	ctx, dk, bk := createTestInput(t, false)

	deposit, found := dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, depositEmpty, deposit)

	err := dk.Subtract(ctx, address1, coinsEmpty)
	require.NotNil(t, err)
	err = dk.Subtract(ctx, address1, coinsNil)
	require.NotNil(t, err)
	err = dk.Subtract(ctx, address1, coinsNeg)
	require.NotNil(t, err)
	err = dk.Subtract(ctx, address1, coinsZero)
	require.NotNil(t, err)
	err = dk.Subtract(ctx, address1, coinsPos)
	require.NotNil(t, err)

	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, depositEmpty, deposit)
	dk.SetDeposit(ctx, depositPos)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.Subtract(ctx, address1, coinsEmpty)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.Subtract(ctx, address1, coinsNil)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.Subtract(ctx, address1, coinsNeg)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.Subtract(ctx, address1, coinsZero)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.Subtract(ctx, address1, coinsPos)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, coinsPos, deposit.Coins)
	coins := bk.GetCoins(ctx, address1)
	require.Equal(t, coinsNil, coins)
}

func TestKeeper_SendFromDepositToAccount(t *testing.T) {
	ctx, dk, bk := createTestInput(t, false)

	deposit, found := dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, depositEmpty, deposit)

	err := dk.SendFromDepositToAccount(ctx, address1, address2, coinsEmpty)
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsNil)
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsNeg)
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsZero)
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsPos)
	require.NotNil(t, err)

	dk.SetDeposit(ctx, depositPos)
	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsEmpty)
	require.Nil(t, err)
	coins := bk.GetCoins(ctx, address2)
	require.Equal(t, coinsNil, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsNil)
	require.Nil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, coinsNil, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsNeg)
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, coinsNil, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsZero)
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, coinsNil, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos, deposit)

	err = dk.SendFromDepositToAccount(ctx, address1, address2, coinsPos)
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, coinsNil, coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsEmpty, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, coinsPos, deposit.Coins)
}

func TestKeeper_ReceiveFromAccountToDeposit(t *testing.T) {
	ctx, dk, bk := createTestInput(t, false)

	deposit, found := dk.GetDeposit(ctx, address2)
	require.Equal(t, false, found)
	coins := bk.GetCoins(ctx, address1)
	require.Equal(t, coinsEmpty, coins)

	err := dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsEmpty)
	require.Nil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsNil)
	require.Nil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsNeg)
	require.NotNil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsZero)
	require.NotNil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsPos)
	require.NotNil(t, err)

	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsNil, deposit.Coins)
	coins, err = bk.AddCoins(ctx, address1, coinsPos)
	require.Nil(t, err)
	require.Equal(t, coinsPos, coins)

	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsEmpty)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsNil, deposit.Coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsPos, coins)

	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsNil)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsNil, deposit.Coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsPos, coins)

	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsNeg)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsNil, deposit.Coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsPos, coins)

	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsZero)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsNil, deposit.Coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsPos, coins)

	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsPos)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsPos, deposit.Coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsNil, coins)

	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsPos)
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsPos, deposit.Coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsNil, coins)

	coins, err = bk.AddCoins(ctx, address1, coinsPos.Add(coinsPos))
	require.Nil(t, err)
	require.Equal(t, coinsPos.Add(coinsPos), coins)
	err = dk.ReceiveFromAccountToDeposit(ctx, address1, address2, coinsPos)
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address2)
	require.Equal(t, true, found)
	require.Equal(t, coinsPos.Add(coinsPos), deposit.Coins)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, coinsPos, coins)
}
