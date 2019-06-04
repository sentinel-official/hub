// nolint
package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/crypto/ed25519"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"

	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

var (
	TestMonikerLenZero     = ""
	TestMonikerValid       = "MONIKER"
	TestNewMonikerValid    = "NEW_MONIKER"
	TestMonikerLengthGT128 = "MONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKER"
	TestCoinPos            = csdkTypes.NewInt64Coin("stake", 100)
	TestCoinNeg            = csdkTypes.Coin{"stake", csdkTypes.NewInt(-100)}
	TestCoinZero           = csdkTypes.NewInt64Coin("stake", 0)
	TestCoinEmpty          = csdkTypes.NewInt64Coin("empty", 0)
	TestCoinNil            = csdkTypes.Coin{}
	TestCoinsPos           = csdkTypes.Coins{TestCoinPos}
	TestCoinsNeg           = csdkTypes.Coins{TestCoinNeg, csdkTypes.Coin{"stake", csdkTypes.NewInt(-100)}}
	TestCoinsZero          = csdkTypes.Coins{TestCoinZero, csdkTypes.NewInt64Coin("stake", 0)}
	TestCoinsInvalid       = csdkTypes.Coins{csdkTypes.NewInt64Coin("stake", 100), TestCoinZero}
	TestCoinsEmpty         = csdkTypes.Coins{}
	TestCoinsNil           = csdkTypes.Coins(nil)
	TestPrivKey1           = ed25519.GenPrivKey()
	TestPrivKey2           = ed25519.GenPrivKey()
	TestPubkey1            = TestPrivKey1.PubKey()
	TestPubkey2            = TestPrivKey2.PubKey()
	TestAddress1           = csdkTypes.AccAddress(TestPubkey1.Address())
	TestAddress2           = csdkTypes.AccAddress(TestPubkey2.Address())
	TestAddressEmpty       = csdkTypes.AccAddress([]byte(""))
	TestUploadNeg          = csdkTypes.NewInt(-1000000)
	TestUploadZero         = csdkTypes.NewInt(0)
	TestUploadPos1         = csdkTypes.NewInt(1000000)
	TestUploadPos2         = TestUploadPos1.Mul(csdkTypes.NewInt(2))
	TestDownloadNeg        = csdkTypes.NewInt(-1000000000)
	TestDownloadZero       = csdkTypes.NewInt(0)
	TestDownloadPos1       = csdkTypes.NewInt(1000000000)
	TestDownloadPos2       = TestDownloadPos1.Mul(csdkTypes.NewInt(2))
	TestBandwidthNeg       = sdkTypes.NewBandwidth(TestUploadNeg, TestDownloadNeg)
	TestBandwidthZero      = sdkTypes.NewBandwidth(TestUploadZero, TestDownloadZero)
	TestBandwidthPos1      = sdkTypes.NewBandwidth(TestUploadPos1, TestDownloadPos1)
	TestBandwidthPos2      = sdkTypes.NewBandwidth(TestUploadPos2, TestDownloadPos2)
	TestEncryption         = "encryption"
	TestNodeType           = "node_type"
	TestVersion            = "version"
	TestStatusInvalid      = "invalid"
	TestNewNodeType        = "NEW_TYPE"
	TestNewVersion         = "NEW_VERSION"
	TestNewEncryption      = "NEW_ENCRYPTION"

	TestIDPos    = sdkTypes.NewIDFromUInt64(1)
	TestIDZero   = sdkTypes.NewIDFromUInt64(0)
	TestIDsEmpty = sdkTypes.IDs(nil)
	TestIDsValid = sdkTypes.IDs{TestIDPos}
)

var (
	TestNodeValid = types.Node{
		ID:               TestIDZero,
		Owner:            TestAddress1,
		Deposit:          TestCoinPos,
		Type:             TestNodeType,
		Version:          TestVersion,
		Moniker:          TestMonikerValid,
		PricesPerGB:      TestCoinsPos,
		InternetSpeed:    TestBandwidthPos1,
		Encryption:       TestEncryption,
		Status:           TestStatusInvalid,
		StatusModifiedAt: 1,
	}
	TestNodeEmpty         = types.Node{}
	TestNodesValid        = []types.Node{TestNodeValid}
	TestNodesEmpty        = []types.Node{}
	TestNodesNil          = []types.Node(nil)
	TestNodeTagsValid     = csdkTypes.EmptyTags().AppendTag("node_id", TestIDPos.String())
	TestSubscriptionValid = types.Subscription{
		ID:                 TestIDZero,
		NodeID:             TestIDZero,
		Client:             TestAddress2,
		PricePerGB:         TestCoinPos,
		TotalDeposit:       TestCoinPos,
		RemainingDeposit:   TestCoinPos,
		RemainingBandwidth: TestBandwidthPos1,
		Status:             TestStatusInvalid,
		StatusModifiedAt:   1,
	}
	TestSubscriptionEmpty  = types.Subscription{}
	TestSubscriptionsValid = []types.Subscription{TestSubscriptionValid}
	TestSubscriptionsEmpty = []types.Subscription{}
	TestSubscriptionsNil   = []types.Subscription(nil)
	TestSessionValid       = types.Session{
		ID:               TestIDZero,
		SubscriptionID:   TestIDZero,
		Bandwidth:        TestBandwidthPos1,
		Status:           TestStatusInvalid,
		StatusModifiedAt: 1,
	}
	TestSessionEmpty  = types.Session{}
	TestSessionsValid = []types.Session{TestSessionValid}
	TestSessionsEmpty = []types.Session{}
	TestSessionsNil   = []types.Session(nil)
)

var (
	TestBandWidthSignDataPos1 = types.NewBandwidthSignatureData(TestIDZero, 1, TestBandwidthPos1)
	TestBandWidthSignDataPos2 = types.NewBandwidthSignatureData(TestIDPos, 2, TestBandwidthPos2)
	TestBandWidthSignDataNeg  = types.NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthNeg)
	TestBandWidthSignDataZero = types.NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthZero)

	TestNodeOwnerSignBandWidthPos1, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos1.Bytes())
	TestNodeOwnerSignBandWidthPos2, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos2.Bytes())
	TestNodeOwnerSignBandWidthNeg, _  = TestPrivKey1.Sign(TestBandWidthSignDataNeg.Bytes())
	TestNodeOwnerSignBandWidthZero, _ = TestPrivKey1.Sign(TestBandWidthSignDataZero.Bytes())

	TestClientSignBandWidthPos1, _ = TestPrivKey2.Sign(TestBandWidthSignDataPos1.Bytes())
	TestClientSignBandWidthPos2, _ = TestPrivKey2.Sign(TestBandWidthSignDataPos2.Bytes())
	TestClientSignBandWidthNeg, _  = TestPrivKey2.Sign(TestBandWidthSignDataNeg.Bytes())
	TestClientSignBandWidthZero, _ = TestPrivKey2.Sign(TestBandWidthSignDataZero.Bytes())

	TestNodeOwnerstdSignaturePos1 = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos1,}
	TestNodeOwnerstdSignaturePos2 = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos2,}
	TestNodeOwnerstdSignatureNeg  = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthNeg,}
	TestNodeOwnerstdSignatureZero = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthZero,}

	TestClientstdSignaturePos1 = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos1,}
	TestClientstdSignaturePos2 = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos2,}
	TestClientstdSignatureNeg  = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthNeg,}
	TestClientstdSignatureZero = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthZero,}

	TeststdSignatureEmpty = auth.StdSignature{}
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

	return ctx, cdc, depositKeeper, vpnKeeper, accountKeeper, bankKeeper
}

func TestMakeCodec() *codec.Codec {
	var cdc = codec.New()
	types.RegisterCodec(cdc)
	auth.RegisterBaseAccount(cdc)
	return cdc
}
