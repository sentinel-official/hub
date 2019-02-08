package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/params"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestCreateInput() (csdkTypes.Context, Keeper, auth.AccountKeeper) {
	keyNode := csdkTypes.NewKVStoreKey("node")
	keySession := csdkTypes.NewKVStoreKey("session")
	keyAccount := csdkTypes.NewKVStoreKey("acc")
	keyParams := csdkTypes.NewKVStoreKey("params")
	tkeyParams := csdkTypes.NewTransientStoreKey("params")

	db := tmDB.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyNode, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySession, csdkTypes.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := TestMakeCodec()
	ctx := csdkTypes.NewContext(ms, abciTypes.Header{ChainID: "chain-id"}, false, log.NewNopLogger())
	keeper := NewKeeper(cdc, keyNode, keySession)
	paramsKeeper := params.NewKeeper(cdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(cdc, keyAccount, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)

	return ctx, keeper, accountKeeper
}

func TestMakeCodec() *codec.Codec {
	var cdc = codec.New()
	return cdc
}

var (
	TestNodeValid = types.NodeDetails{
		ID:              types.TestNodeIDValid,
		Owner:           types.TestAddress,
		LockedAmount:    types.TestCoinPos,
		PricesPerGB:     types.TestCoinsPos,
		NetSpeed:        sdkTypes.NewBandwidth(types.TestUploadPos, types.TestDownloadPos),
		APIPort:         types.TestAPIPortValid,
		EncMethod:       types.TestEncMethod,
		NodeType:        types.TestNodeType,
		Version:         types.TestVersion,
		Status:          types.StatusRegistered,
		StatusAtHeight:  1,
		DetailsAtHeight: 1,
	}
	TestNodeEmpty     = types.NodeDetails{}
	TestNodeIDsEmpty  = types.NodeIDs(nil)
	TestNodeIDsValid  = types.NodeIDs{types.TestNodeIDValid, types.TestNodeIDValid}
	TestNodesEmpty    = []*types.NodeDetails(nil)
	TestNodeTagsValid = csdkTypes.EmptyTags().AppendTag("node_id", types.TestNodeIDValid.Bytes())

	TestSessionValid = types.SessionDetails{
		ID:           types.TestSessionIDValid,
		NodeID:       types.TestNodeIDValid,
		Client:       types.TestAddress,
		LockedAmount: types.TestCoinPos,
		PricePerGB:   types.TestCoinPos,
		Bandwidth: types.SessionBandwidth{
			ToProvide:       sdkTypes.NewBandwidth(types.TestUploadPos, types.TestDownloadPos),
			Consumed:        sdkTypes.NewBandwidth(types.TestUploadPos, types.TestDownloadPos),
			NodeOwnerSign:   types.TestNodeOwnerSign,
			ClientSign:      types.TestClientSign,
			UpdatedAtHeight: 2,
		},
		Status:          types.StatusInit,
		StatusAtHeight:  1,
		StartedAtHeight: 2,
		EndedAtHeight:   3,
	}
	TestSessionEmpty     = types.SessionDetails{}
	TestSessionIDsEmpty  = types.SessionIDs(nil)
	TestSessionIDsValid  = types.SessionIDs{types.TestSessionIDValid, types.TestSessionIDValid}
	TestSessionsEmpty    = []*types.SessionDetails(nil)
	TestSessionTagsValid = csdkTypes.EmptyTags().AppendTag("session_id", types.TestSessionIDValid.Bytes())
)
