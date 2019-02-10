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

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestCreateInput() (csdkTypes.Context, *codec.Codec, Keeper, auth.AccountKeeper, bank.BaseKeeper) {
	keyNode := csdkTypes.NewKVStoreKey("node")
	keySession := csdkTypes.NewKVStoreKey("session")
	keyAccount := csdkTypes.NewKVStoreKey("acc")
	keyParams := csdkTypes.NewKVStoreKey("params")
	tkeyParams := csdkTypes.NewTransientStoreKey("tparams")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyNode, csdkTypes.StoreTypeIAVL, db)
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

	keeper := NewKeeper(cdc, keyNode, keySession)
	paramsKeeper := params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAccount, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)

	types.RegisterCodec(cdc)
	auth.RegisterBaseAccount(cdc)

	return ctx, cdc, keeper, accountKeeper, bankKeeper
}

func TestMakeCodec() *codec.Codec {
	var cdc = codec.New()
	return cdc
}

var (
	TestNodeValid = types.NodeDetails{
		ID:              types.TestNodeIDValid,
		Owner:           types.TestAddress1,
		LockedAmount:    types.TestCoinPos,
		PricesPerGB:     types.TestCoinsPos,
		NetSpeed:        types.TestBandwidthPos,
		APIPort:         types.TestAPIPortValid,
		EncMethod:       types.TestEncMethod,
		NodeType:        types.TestNodeType,
		Version:         types.TestVersion,
		Status:          types.StatusRegistered,
		StatusAtHeight:  0,
		DetailsAtHeight: 0,
	}
	TestNodeEmpty     = types.NodeDetails{}
	TestNodeIDsEmpty  = types.NodeIDs(nil)
	TestNodeIDsValid  = types.NodeIDs{types.TestNodeIDValid, types.TestNodeIDValid}
	TestNodesEmpty    = []*types.NodeDetails(nil)
	TestNodeTagsValid = csdkTypes.EmptyTags().AppendTag("node_id", types.TestNodeIDValid.String())

	TestSessionValid = types.SessionDetails{
		ID:           types.TestSessionIDValid,
		NodeID:       types.TestNodeIDValid,
		Client:       types.TestAddress2,
		LockedAmount: types.TestCoinPos,
		PricePerGB:   types.TestCoinPos,
		Bandwidth: types.SessionBandwidth{
			ToProvide:       types.TestBandwidthPos,
			Consumed:        types.TestBandwidthZero,
			NodeOwnerSign:   types.TestNodeOwnerSign,
			ClientSign:      types.TestClientSign,
			UpdatedAtHeight: 0,
		},
		Status:          types.StatusInit,
		StatusAtHeight:  0,
		StartedAtHeight: 0,
		EndedAtHeight:   0,
	}
	TestSessionEmpty     = types.SessionDetails{}
	TestSessionIDsEmpty  = types.SessionIDs(nil)
	TestSessionIDsValid  = types.SessionIDs{types.TestSessionIDValid, types.TestSessionIDValid}
	TestSessionsEmpty    = []*types.SessionDetails(nil)
	TestSessionTagsValid = csdkTypes.EmptyTags().AppendTag("session_id", types.TestSessionIDValid.String())
)
