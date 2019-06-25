// nolint
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/sentinel-hub/x/deposit"
	"github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

func TestCreateInput() (sdk.Context, deposit.Keeper, Keeper, bank.BaseKeeper) {
	keyDeposits := sdk.NewKVStoreKey("deposits")
	keyNode := sdk.NewKVStoreKey("node")
	keySession := sdk.NewKVStoreKey("session")
	keySubscription := sdk.NewKVStoreKey("subscription")
	keyAccount := sdk.NewKVStoreKey("acc")
	keyParams := sdk.NewKVStoreKey("params")
	tkeyParams := sdk.NewTransientStoreKey("tparams")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyDeposits, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyNode, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySubscription, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySession, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAccount, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := TestMakeCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, false, log.NewNopLogger())

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
