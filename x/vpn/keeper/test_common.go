package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"

	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type nodeList []types.NodeDetails

type sessionList []types.SessionDetails

func CreateTestInput() (Keeper, csdkTypes.Context) {
	keyNode := csdkTypes.NewKVStoreKey("store1")
	keySession := csdkTypes.NewKVStoreKey("store2")

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyNode, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySession, csdkTypes.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	cdc := MakeCdc()

	ctx := csdkTypes.NewContext(ms, abci.Header{ChainID: "chaind-id"}, false, log.NewNopLogger())

	keeper := NewKeeper(cdc, keyNode, keySession)

	return keeper, ctx
}

func MakeCdc() *codec.Codec {
	var cdc = codec.New()
	return cdc
}

func ParamsOfNodeDetails() []types.NodeDetails {

	id1 := types.NodeDetails{
		ID:           types.NewNodeID("new-node-id-1"),
		Owner:        types.NodeAddress1,
		LockedAmount: types.Coin,
		PricesPerGB:  types.Coins,
		NetSpeed: sdkTypes.Bandwidth{
			csdkTypes.NewInt(12),
			csdkTypes.NewInt(12),
		},
		APIPort:         types.NewAPIPort(uint32(1000)),
		EncMethod:       "EncMethod-1",
		NodeType:        "NodeType-1",
		Version:         "0.01",
		Status:          types.StatusRegistered,
		StatusAtHeight:  8,
		DetailsAtHeight: 8,
	}

	id2 := types.NodeDetails{
		ID:           types.NewNodeID("new-node-id-2"),
		Owner:        types.NodeAddress2,
		LockedAmount: types.Coin,
		PricesPerGB:  types.Coins,
		NetSpeed: sdkTypes.Bandwidth{
			csdkTypes.NewInt(12),
			csdkTypes.NewInt(12),
		},
		APIPort:         types.NewAPIPort(uint32(1000)),
		EncMethod:       "EncMethod-2",
		NodeType:        "NodeType-1",
		Version:         "0.01",
		Status:          types.StatusRegistered,
		StatusAtHeight:  10,
		DetailsAtHeight: 10,
	}

	id3 := types.NodeDetails{}

	return nodeList{id1, id2, id3}
}

func ParamsOfSessionDetails() []types.SessionDetails {

	id1 := types.SessionDetails{
		ID:              types.NewSessionID("new-session-id-1"),
		NodeID:          types.NewNodeID("new-node-id-1"),
		Client:          types.ClientAddress1,
		LockedAmount:    types.Coin,
		PricePerGB:      types.Coin,
		Bandwidth:       types.SessionBandwidth{},
		Status:          types.StatusActive,
		StatusAtHeight:  12,
		StartedAtHeight: 10,
		EndedAtHeight:   20,
	}

	id2 := types.SessionDetails{
		ID:              types.NewSessionID("new-session-id-2"),
		NodeID:          types.NewNodeID("new-node-id-1"),
		Client:          types.ClientAddress1,
		LockedAmount:    types.Coin,
		PricePerGB:      types.Coin,
		Bandwidth:       types.SessionBandwidth{},
		Status:          types.StatusRegistered,
		StatusAtHeight:  12,
		StartedAtHeight: 10,
		EndedAtHeight:   20,
	}

	id3 := types.SessionDetails{}

	return sessionList{id1, id2, id3}
}
