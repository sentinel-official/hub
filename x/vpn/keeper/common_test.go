package keeper_test

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
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

var (
	privKey1 = ed25519.GenPrivKey()
	privKey2 = ed25519.GenPrivKey()
	pubkey1  = privKey1.PubKey()
	pubkey2  = privKey2.PubKey()
	address1 = sdk.AccAddress(pubkey1.Address())
	address2 = sdk.AccAddress(pubkey2.Address())

	nodeValid = types.Node{
		ID:               hub.NewIDFromUInt64(0),
		Owner:            address1,
		Deposit:          sdk.NewInt64Coin("stake", 100),
		Type:             "node_type",
		Version:          "version",
		Moniker:          "moniker",
		PricesPerGB:      sdk.Coins{sdk.NewInt64Coin("stake", 100)},
		InternetSpeed:    hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)),
		Encryption:       "encryption",
		Status:           types.StatusInactive,
		StatusModifiedAt: 1,
	}
	subscriptionValid = types.Subscription{
		ID:                 hub.NewIDFromUInt64(0),
		NodeID:             hub.NewIDFromUInt64(0),
		Client:             address2,
		PricePerGB:         sdk.NewInt64Coin("stake", 100),
		TotalDeposit:       sdk.NewInt64Coin("stake", 100),
		RemainingDeposit:   sdk.NewInt64Coin("stake", 100),
		RemainingBandwidth: hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)),
		Status:             types.StatusActive,
		StatusModifiedAt:   0,
	}
	sessionValid = types.Session{
		ID:               hub.NewIDFromUInt64(0),
		SubscriptionID:   hub.NewIDFromUInt64(0),
		Bandwidth:        hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)),
		Status:           types.StatusActive,
		StatusModifiedAt: 0,
	}
)

func CreateTestInput(t *testing.T, isCheckTx bool) (sdk.Context, keeper.Keeper, deposit.Keeper, bank.Keeper) {
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
	vk := keeper.NewKeeper(cdc, keyNode, keySubscription, keySession, pk.Subspace(keeper.DefaultParamspace), dk)

	sk.SetModuleAccount(ctx, depositAccount)

	return ctx, vk, dk, bk
}

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	return cdc
}
