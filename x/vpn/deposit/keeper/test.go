package keeper

import (
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
	tdb "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub/x/vpn/deposit/types"
)

func CreateTestInput(t *testing.T, isCheckTx bool) (sdk.Context, Keeper, bank.Keeper) {
	var (
		db                 = tdb.NewMemDB()
		cms                = store.NewCommitMultiStore(db)
		ctx                = sdk.NewContext(cms, abci.Header{ChainID: "test"}, isCheckTx, log.NewNopLogger())
		cdc                = MakeTestCodec()
		keyParams          = sdk.NewKVStoreKey(params.StoreKey)
		keyAuth            = sdk.NewKVStoreKey(auth.StoreKey)
		keySupply          = sdk.NewKVStoreKey(supply.StoreKey)
		keyDeposit         = sdk.NewKVStoreKey(types.StoreKey)
		transientKeyParams = sdk.NewTransientStoreKey(params.TStoreKey)
		depositAccount     = supply.NewEmptyModuleAccount(types.ModuleName)
		blacklist          = map[string]bool{depositAccount.String(): true}
		permissions        = map[string][]string{types.ModuleName: nil}
	)

	cms.MountStoreWithDB(keyAuth, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(keyDeposit, sdk.StoreTypeIAVL, db)
	require.Nil(t, cms.LoadLatestVersion())

	var (
		paramsKeeper = params.NewKeeper(cdc,
			keyParams,
			transientKeyParams,
			params.DefaultCodespace)
		accountKeeper = auth.NewAccountKeeper(cdc,
			keyAuth,
			paramsKeeper.Subspace(auth.DefaultParamspace),
			auth.ProtoBaseAccount)
		bankKeeper = bank.NewBaseKeeper(accountKeeper,
			paramsKeeper.Subspace(bank.DefaultParamspace),
			bank.DefaultCodespace,
			blacklist)
		supplyKeeper = supply.NewKeeper(cdc,
			keySupply,
			accountKeeper,
			bankKeeper,
			permissions)
		keeper = NewKeeper(cdc,
			keyDeposit)
	)

	keeper.WithSupplyKeeper(supplyKeeper)
	return ctx, keeper, bankKeeper
}

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)

	return cdc
}
