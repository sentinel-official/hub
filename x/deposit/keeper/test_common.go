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
	db "github.com/tendermint/tm-db"
	
	"github.com/sentinel-official/hub/x/deposit/types"
)

func CreateTestInput(t *testing.T, isCheckTx bool) (sdk.Context, Keeper, bank.Keeper) {
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	keyAccount := sdk.NewKVStoreKey(auth.StoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	keyDeposits := sdk.NewKVStoreKey(types.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)
	
	mdb := db.NewMemDB()
	ms := store.NewCommitMultiStore(mdb)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyAccount, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(keyDeposits, sdk.StoreTypeIAVL, mdb)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, mdb)
	require.Nil(t, ms.LoadLatestVersion())
	
	depositAccount := supply.NewEmptyModuleAccount(types.ModuleName)
	blacklist := make(map[string]bool)
	blacklist[depositAccount.String()] = true
	accountPermissions := map[string][]string{
		types.ModuleName: nil,
	}
	
	cdc := MakeTestCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, isCheckTx, log.NewNopLogger())
	
	pk := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)
	ak := auth.NewAccountKeeper(cdc, keyAccount, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bk := bank.NewBaseKeeper(ak, pk.Subspace(bank.DefaultParamspace), bank.DefaultCodespace, blacklist)
	sk := supply.NewKeeper(cdc, keySupply, ak, bk, accountPermissions)
	dk := NewKeeper(cdc, keyDeposits, sk)
	
	sk.SetModuleAccount(ctx, depositAccount)
	
	return ctx, dk, bk
}

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)
	return cdc
}
