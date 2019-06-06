// nolint
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"

	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

func TestCreateInput() (csdkTypes.Context, *codec.Codec, deposit.Keeper, Keeper, auth.AccountKeeper, bank.BaseKeeper) {
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

	depositKeeper := deposit.NewKeeper(cdc, keyDeposits, bankKeeper)

	vpnKeeper := NewKeeper(cdc, keyNode, keySubscription, keySession, paramsKeeper.Subspace(DefaultParamspace), depositKeeper)

	params := types.Params{
		FreeNodesCount:          types.DefaultFreeNodesCount,
		Deposit:                 types.DefaultDeposit,
		NodeInactiveInterval:    types.DefaultNodeInactiveInterval,
		SessionInactiveInterval: types.DefaultSessionInactiveInterval,
	}

	vpnKeeper.SetParams(ctx, params)

	return ctx, cdc, depositKeeper, vpnKeeper, accountKeeper, bankKeeper
}

func TestMakeCodec() *codec.Codec {
	var cdc = codec.New()
	types.RegisterCodec(cdc)
	auth.RegisterBaseAccount(cdc)
	return cdc
}
