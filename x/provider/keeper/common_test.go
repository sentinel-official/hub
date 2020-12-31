package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintdb "github.com/tendermint/tm-db"

	app "github.com/sentinel-official/hub"
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	keeper   keeper.Keeper
	storeKey sdk.StoreKey
	cdc      *codec.Codec
	ctx      sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {

	suite.storeKey = sdk.NewKVStoreKey(types.StoreKey)

	memDB := tendermintdb.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)

	ms.MountStoreWithDB(suite.storeKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, abci.Header{ChainID: "testchainid"}, false, log.NewNopLogger())
	suite.cdc = app.MakeCodec()
	suite.keeper = keeper.NewKeeper(suite.cdc, suite.storeKey)

}
