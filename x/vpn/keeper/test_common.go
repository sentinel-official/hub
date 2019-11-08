package keeper

import (
	"math/rand"
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
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func CreateTestInput(t *testing.T, isCheckTx bool) (sdk.Context, Keeper, deposit.Keeper, bank.Keeper) {
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	keyAccount := sdk.NewKVStoreKey(auth.StoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	keyDeposit := sdk.NewKVStoreKey(deposit.StoreKey)
	keyNode := sdk.NewKVStoreKey(types.StoreKeyNode)
	keySubscription := sdk.NewKVStoreKey(types.StoreKeySubscription)
	keySession := sdk.NewKVStoreKey(types.StoreKeySession)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)

	mdb := db.NewMemDB()
	ms := store.NewCommitMultiStore(mdb)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyAccount, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyDeposit, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyNode, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySubscription, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySession, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, mdb)
	require.Nil(t, ms.LoadLatestVersion())

	depositAccount := supply.NewEmptyModuleAccount(types.ModuleName)
	blacklist := make(map[string]bool)
	blacklist[depositAccount.String()] = true
	accountPermissions := map[string][]string{
		deposit.ModuleName: nil,
	}

	cdc := MakeTestCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, isCheckTx, log.NewNopLogger())

	pk := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	ak := auth.NewAccountKeeper(cdc, keyAccount, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bk := bank.NewBaseKeeper(ak, pk.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, blacklist)
	sk := supply.NewKeeper(cdc, keySupply, ak, bk, accountPermissions)
	dk := deposit.NewKeeper(cdc, keyDeposit, sk)
	vk := NewKeeper(cdc, keyNode, keySubscription, keySession, pk.Subspace(DefaultParamspace), dk)

	sk.SetModuleAccount(ctx, depositAccount)
	vk.SetParams(ctx, types.DefaultParams())

	return ctx, vk, dk, bk
}

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)
	return cdc
}

func RandomNode(r *rand.Rand, ctx sdk.Context, keeper Keeper) types.Node {
	nodes := keeper.GetAllNodes(ctx)
	i := r.Intn(len(nodes))

	return nodes[i]
}

func RandomSubscription(r *rand.Rand, ctx sdk.Context, keeper Keeper) types.Subscription {
	subscriptions := keeper.GetAllSubscriptions(ctx)
	i := r.Intn(len(subscriptions))

	return subscriptions[i]
}

func RandomSession(r *rand.Rand, ctx sdk.Context, keeper Keeper) types.Session {
	sessions := keeper.GetAllSessions(ctx)
	i := r.Intn(len(sessions))

	return sessions[i]
}
