// nolint
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto/ed25519"

	hub "github.com/sentinel-official/hub/types"
)

var (
	TestPrivKey1     = ed25519.GenPrivKey()
	TestPrivKey2     = ed25519.GenPrivKey()
	TestPubkey1      = TestPrivKey1.PubKey()
	TestPubkey2      = TestPrivKey2.PubKey()
	TestAddress1     = sdk.AccAddress(TestPubkey1.Address())
	TestAddress2     = sdk.AccAddress(TestPubkey2.Address())
	TestAddressEmpty = sdk.AccAddress([]byte(""))

	TestCoinEmpty   = sdk.Coin{}
	TestCoinNeg     = sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)}
	TestCoinZero    = sdk.NewInt64Coin("stake", 0)
	TestCoinPos     = sdk.NewInt64Coin("stake", 100)
	TestCoinInvalid = sdk.NewInt64Coin("invalid", 100)

	TestCoinsEmpty   = sdk.Coins{}
	TestCoinsNil     = sdk.Coins(nil)
	TestCoinsNeg     = sdk.Coins{TestCoinNeg}
	TestCoinsZero    = sdk.Coins{TestCoinZero}
	TestCoinsPos     = sdk.Coins{TestCoinPos}
	TestCoinsInvalid = sdk.Coins{TestCoinInvalid}

	TestIDZero = hub.NewIDFromUInt64(0)
	TestIDPos  = hub.NewIDFromUInt64(1)

	TestIDsEmpty = hub.IDs{}
	TestIDsNil   = hub.IDs(nil)
	TestIDsValid = hub.IDs{TestIDZero}

	TestNodeType    = "NODE_TYPE"
	TestNewNodeType = "NEW_NODE_TYPE"

	TestVersion    = "VERSION"
	TestNewVersion = "NEW_VERSION"

	TestMonikerLenZero     = ""
	TestMonikerValid       = "MONIKER"
	TestNewMonikerValid    = "NEW_MONIKER"
	TestMonikerLengthGT128 = "MONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKER" +
		"MONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKER"

	TestUploadNeg  = sdk.NewInt(-500000000)
	TestUploadZero = sdk.NewInt(0)
	TestUploadPos1 = sdk.NewInt(500000000)
	TestUploadPos2 = TestUploadPos1.Mul(sdk.NewInt(2))

	TestDownloadNeg  = sdk.NewInt(-500000000)
	TestDownloadZero = sdk.NewInt(0)
	TestDownloadPos1 = sdk.NewInt(500000000)
	TestDownloadPos2 = TestDownloadPos1.Mul(sdk.NewInt(2))

	TestBandwidthNeg  = hub.NewBandwidth(TestUploadNeg, TestDownloadNeg)
	TestBandwidthZero = hub.NewBandwidth(TestUploadZero, TestDownloadZero)
	TestBandwidthPos1 = hub.NewBandwidth(TestUploadPos1, TestDownloadPos1)
	TestBandwidthPos2 = hub.NewBandwidth(TestUploadPos2, TestDownloadPos2)

	TestEncryption    = "ENCRYPTION"
	TestNewEncryption = "NEW_ENCRYPTION"

	TestStatusActive   = "ACTIVE"
	TestStatusInActive = "INACTIVE"
	TestStatusInValid  = "STATUS"

	TestNodeEmpty = Node{}
	TestNodeValid = Node{
		ID:               TestIDZero,
		Owner:            TestAddress1,
		Deposit:          TestCoinPos,
		Type:             TestNodeType,
		Version:          TestVersion,
		Moniker:          TestMonikerValid,
		PricesPerGB:      TestCoinsPos,
		InternetSpeed:    TestBandwidthPos1,
		Encryption:       TestEncryption,
		Status:           TestStatusInActive,
		StatusModifiedAt: 1,
	}

	TestNodesEmpty = []Node{}
	TestNodesNil   = []Node(nil)
	TestNodesValid = []Node{TestNodeValid}

	TestSubscriptionEmpty = Subscription{}
	TestSubscriptionValid = Subscription{
		ID:                 TestIDZero,
		NodeID:             TestIDZero,
		Client:             TestAddress2,
		PricePerGB:         TestCoinPos,
		TotalDeposit:       TestCoinPos,
		RemainingDeposit:   TestCoinPos,
		RemainingBandwidth: TestBandwidthPos1,
		Status:             TestStatusActive,
		StatusModifiedAt:   0,
	}

	TestSubscriptionsEmpty = []Subscription{}
	TestSubscriptionsNil   = []Subscription(nil)
	TestSubscriptionsValid = []Subscription{TestSubscriptionValid}

	TestSessionEmpty = Session{}
	TestSessionValid = Session{
		ID:               TestIDZero,
		SubscriptionID:   TestIDZero,
		Bandwidth:        TestBandwidthPos1,
		Status:           TestStatusActive,
		StatusModifiedAt: 0,
	}

	TestSessionsEmpty = []Session{}
	TestSessionsNil   = []Session(nil)
	TestSessionsValid = []Session{TestSessionValid}

	TestBandWidthSignDataNeg  = NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthNeg)
	TestBandWidthSignDataZero = NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthZero)
	TestBandWidthSignDataPos1 = NewBandwidthSignatureData(TestIDZero, 1, TestBandwidthPos1)
	TestBandWidthSignDataPos2 = NewBandwidthSignatureData(TestIDZero, 1, TestBandwidthPos2)

	TestNodeOwnerSignBandWidthNeg, _  = TestPrivKey1.Sign(TestBandWidthSignDataNeg.Bytes())
	TestNodeOwnerSignBandWidthZero, _ = TestPrivKey1.Sign(TestBandWidthSignDataZero.Bytes())
	TestNodeOwnerSignBandWidthPos1, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos1.Bytes())
	TestNodeOwnerSignBandWidthPos2, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos2.Bytes())

	TestClientSignBandWidthNeg, _  = TestPrivKey2.Sign(TestBandWidthSignDataNeg.Bytes())
	TestClientSignBandWidthZero, _ = TestPrivKey2.Sign(TestBandWidthSignDataZero.Bytes())
	TestClientSignBandWidthPos1, _ = TestPrivKey2.Sign(TestBandWidthSignDataPos1.Bytes())
	TestClientSignBandWidthPos2, _ = TestPrivKey2.Sign(TestBandWidthSignDataPos2.Bytes())

	TestStdSignatureEmpty         = auth.StdSignature{}
	TestNodeOwnerStdSignatureNeg  = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthNeg}
	TestNodeOwnerStdSignatureZero = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthZero}
	TestNodeOwnerStdSignaturePos1 = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos1}
	TestNodeOwnerStdSignaturePos2 = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos2}

	TestClientStdSignatureNeg  = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthNeg}
	TestClientStdSignatureZero = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthZero}
	TestClientStdSignaturePos1 = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos1}
	TestClientStdSignaturePos2 = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos2}
)
