// nolint
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestCreateInput() (csdk.Context, deposit.Keeper, Keeper, bank.BaseKeeper) {
	keyDeposits := csdk.NewKVStoreKey("deposits")
	keyNode := csdk.NewKVStoreKey("node")
	keySession := csdk.NewKVStoreKey("session")
	keySubscription := csdk.NewKVStoreKey("subscription")
	keyAccount := csdk.NewKVStoreKey("acc")
	keyParams := csdk.NewKVStoreKey("params")
	tkeyParams := csdk.NewTransientStoreKey("tparams")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyDeposits, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyNode, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySubscription, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySession, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAccount, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, csdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, csdk.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := TestMakeCodec()
	ctx := csdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, false, log.NewNopLogger())

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
