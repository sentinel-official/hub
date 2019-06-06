// nolint
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestCreateInput() (csdkTypes.Context, deposit.Keeper, Keeper, bank.BaseKeeper) {
	keyDeposits := csdkTypes.NewKVStoreKey("deposits")
	keyNode := csdkTypes.NewKVStoreKey("node")
	keySession := csdkTypes.NewKVStoreKey("session")
	keySubscription := csdkTypes.NewKVStoreKey("subscription")
	keyAccount := csdkTypes.NewKVStoreKey("acc")
	keyParams := csdkTypes.NewKVStoreKey("params")
	tkeyParams := csdkTypes.NewTransientStoreKey("tparams")

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

	paramsKeeper := params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAccount, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	depositKeeper := deposit.NewKeeper(cdc, keyDeposits, bankKeeper)
	vpnKeeper := NewKeeper(cdc, keyNode, keySubscription, keySession, paramsKeeper.Subspace(DefaultParamspace), depositKeeper)

	vpnKeeper.SetParams(ctx, types.DefaultParams())

	return ctx, depositKeeper, vpnKeeper, bankKeeper
}

func TestMakeCodec() *codec.Codec {
	var cdc = codec.New()
	types.RegisterCodec(cdc)
	auth.RegisterBaseAccount(cdc)
	return cdc
}
