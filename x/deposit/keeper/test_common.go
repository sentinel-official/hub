// nolint
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/crypto/ed25519"

	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
	vpnTypes "github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

var (
	TestCoinPos   = csdkTypes.NewInt64Coin("stake", 10)
	TestCoinNeg   = csdkTypes.Coin{"stake", csdkTypes.NewInt(-10)}
	TestCoinZero  = csdkTypes.NewInt64Coin("stake", 0)
	TestCoinEmpty = csdkTypes.NewInt64Coin("empty", 0)
	TestCoinNil   = csdkTypes.Coin{}

	TestCoinsPos     = csdkTypes.Coins{TestCoinPos}
	TestCoinsNeg     = csdkTypes.Coins{TestCoinNeg, csdkTypes.Coin{"stake", csdkTypes.NewInt(-100)}}
	TestCoinsZero    = csdkTypes.Coins{TestCoinZero, csdkTypes.NewInt64Coin("stake", 0)}
	TestCoinsInvalid = csdkTypes.Coins{csdkTypes.NewInt64Coin("stake", 100), TestCoinZero}
	TestCoinsEmpty   = csdkTypes.Coins{}
	TestCoinsNil     = csdkTypes.Coins(nil)

	TestPrivKey1 = ed25519.GenPrivKey()
	TestPrivKey2 = ed25519.GenPrivKey()

	TestPubkey1 = TestPrivKey1.PubKey()
	TestPubkey2 = TestPrivKey2.PubKey()

	TestAddress1 = csdkTypes.AccAddress(TestPubkey1.Address())
	TestAddress2 = csdkTypes.AccAddress(TestPubkey2.Address())

	TestAddressEmpty = csdkTypes.AccAddress([]byte(""))
)

var (
	TestDepositPos   = types.Deposit{Address: TestAddress1, Coins: TestCoinsPos}
	TestDepositZero  = types.Deposit{Address: TestAddress1, Coins: TestCoinsZero}
	TestDepositEmpty = types.Deposit{}

	TestDepositsPos   = []types.Deposit{TestDepositPos}
	TestDepositsZero  = []types.Deposit{TestDepositZero}
	TestDepositsEmpty = []types.Deposit{}
	TestDepositsNil   = []types.Deposit(nil)
)

func TestCreateInput() (csdkTypes.Context, *codec.Codec, Keeper, auth.AccountKeeper, bank.BaseKeeper) {

	keyDeposits := csdkTypes.NewKVStoreKey("deposits")
	keyNode := csdkTypes.NewKVStoreKey("node")
	keySession := csdkTypes.NewKVStoreKey("session")
	keySubscription := csdkTypes.NewKVStoreKey("subscription")
	keyAccount := csdkTypes.NewKVStoreKey("acc")
	keyParams := csdkTypes.NewKVStoreKey("params")
	tkeyParams := csdkTypes.NewTransientStoreKey("tparams")

	paramsKeeper := params.NewKeeper(TestMakeCodec(), keyParams, tkeyParams)

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyDeposits, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyNode, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySubscription, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySession, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAccount, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, csdkTypes.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := TestMakeCodec()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{ChainID: "chain-id"}, false, log.NewNopLogger())

	paramsKeeper = params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAccount, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)

	depositKeeper := NewKeeper(cdc, keyDeposits, bankKeeper)

	return ctx, cdc, depositKeeper, accountKeeper, bankKeeper
}

func TestMakeCodec() *codec.Codec {
	var cdc = codec.New()
	vpnTypes.RegisterCodec(cdc)
	auth.RegisterBaseAccount(cdc)
	return cdc
}
