// nolint
package types

import (
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto/ed25519"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

var (
	TestMonikerLenZero     = ""
	TestMonikerValid       = "MONIKER"
	TestNewMonikerValid    = "NEW_MONIKER"
	TestMonikerLengthGT128 = "MONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKER"
	TestCoinPos            = csdk.NewInt64Coin("stake", 100)
	TestCoinNeg            = csdk.Coin{Denom: "stake", Amount: csdk.NewInt(-100)}
	TestCoinZero           = csdk.NewInt64Coin("stake", 0)
	TestCoinEmpty          = csdk.Coin{}
	TestCoinsPos           = csdk.Coins{TestCoinPos}
	TestCoinsNeg           = csdk.Coins{TestCoinNeg, csdk.Coin{Denom: "stake", Amount: csdk.NewInt(-100)}}
	TestCoinsZero          = csdk.Coins{TestCoinZero, csdk.NewInt64Coin("stake", 0)}
	TestCoinsInvalid       = csdk.Coins{csdk.NewInt64Coin("stake", 100), TestCoinZero}
	TestCoinsEmpty         = csdk.Coins{}
	TestCoinsNil           = csdk.Coins(nil)
	TestPrivKey1           = ed25519.GenPrivKey()
	TestPrivKey2           = ed25519.GenPrivKey()
	TestPubkey1            = TestPrivKey1.PubKey()
	TestPubkey2            = TestPrivKey2.PubKey()
	TestAddress1           = csdk.AccAddress(TestPubkey1.Address())
	TestAddress2           = csdk.AccAddress(TestPubkey2.Address())
	TestAddressEmpty       = csdk.AccAddress([]byte(""))
	TestUploadNeg          = csdk.NewInt(-1000000)
	TestUploadZero         = csdk.NewInt(0)
	TestUploadPos1         = csdk.NewInt(1000000)
	TestUploadPos2         = TestUploadPos1.Mul(csdk.NewInt(2))
	TestDownloadNeg        = csdk.NewInt(-1000000000)
	TestDownloadZero       = csdk.NewInt(0)
	TestDownloadPos1       = csdk.NewInt(1000000000)
	TestDownloadPos2       = TestDownloadPos1.Mul(csdk.NewInt(2))
	TestBandwidthNeg       = sdk.NewBandwidth(TestUploadNeg, TestDownloadNeg)
	TestBandwidthZero      = sdk.NewBandwidth(TestUploadZero, TestDownloadZero)
	TestBandwidthPos1      = sdk.NewBandwidth(TestUploadPos1, TestDownloadPos1)
	TestBandwidthPos2      = sdk.NewBandwidth(TestUploadPos2, TestDownloadPos2)
	TestEncryption         = "ENCRYPTION"
	TestNodeType           = "NODE_TYPE"
	TestVersion            = "VERSION"
	TestStatusActive       = "ACTIVE"
	TestStatusInActive     = "INACTIVE"
	TestStatusInValid      = "STATUS"
	TestNewNodeType        = "NEW_NODE_TYPE"
	TestNewVersion         = "NEW_VERSION"
	TestNewEncryption      = "NEW_ENCRYPTION"
	TestIDPos              = sdk.NewIDFromUInt64(1)
	TestIDZero             = sdk.NewIDFromUInt64(0)
	TestIDsNil             = sdk.IDs(nil)
	TestIDsEmpty           = sdk.IDs{}
	TestIDsValid           = sdk.IDs{TestIDZero}
	TestNodeValid          = Node{
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
	TestNodeEmpty         = Node{}
	TestNodesValid        = []Node{TestNodeValid}
	TestNodesEmpty        = []Node{}
	TestNodesNil          = []Node(nil)
	TestSubscriptionValid = Subscription{
		ID:                 TestIDZero,
		NodeID:             TestIDZero,
		Client:             TestAddress2,
		PricePerGB:         TestCoinPos,
		TotalDeposit:       TestCoinPos,
		RemainingDeposit:   TestCoinPos,
		RemainingBandwidth: TestBandwidthPos1,
		Status:             TestStatusInActive,
		StatusModifiedAt:   1,
	}
	TestSubscriptionEmpty  = Subscription{}
	TestSubscriptionsValid = []Subscription{TestSubscriptionValid}
	TestSubscriptionsEmpty = []Subscription{}
	TestSubscriptionsNil   = []Subscription(nil)
	TestSessionValid       = Session{
		ID:               TestIDZero,
		SubscriptionID:   TestIDZero,
		Bandwidth:        TestBandwidthPos1,
		Status:           TestStatusInActive,
		StatusModifiedAt: 1,
	}
	TestSessionEmpty                  = Session{}
	TestSessionsValid                 = []Session{TestSessionValid}
	TestSessionsEmpty                 = []Session{}
	TestSessionsNil                   = []Session(nil)
	TestBandWidthSignDataPos1         = NewBandwidthSignatureData(TestIDZero, 1, TestBandwidthPos1)
	TestBandWidthSignDataPos2         = NewBandwidthSignatureData(TestIDPos, 2, TestBandwidthPos2)
	TestBandWidthSignDataNeg          = NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthNeg)
	TestBandWidthSignDataZero         = NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthZero)
	TestNodeOwnerSignBandWidthPos1, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos1.Bytes())
	TestNodeOwnerSignBandWidthPos2, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos2.Bytes())
	TestNodeOwnerSignBandWidthNeg, _  = TestPrivKey1.Sign(TestBandWidthSignDataNeg.Bytes())
	TestNodeOwnerSignBandWidthZero, _ = TestPrivKey1.Sign(TestBandWidthSignDataZero.Bytes())
	TestClientSignBandWidthPos1, _    = TestPrivKey2.Sign(TestBandWidthSignDataPos1.Bytes())
	TestClientSignBandWidthPos2, _    = TestPrivKey2.Sign(TestBandWidthSignDataPos2.Bytes())
	TestClientSignBandWidthNeg, _     = TestPrivKey2.Sign(TestBandWidthSignDataNeg.Bytes())
	TestClientSignBandWidthZero, _    = TestPrivKey2.Sign(TestBandWidthSignDataZero.Bytes())
	TestNodeOwnerStdSignaturePos1     = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos1}
	TestNodeOwnerStdSignaturePos2     = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos2}
	TestNodeOwnerStdSignatureNeg      = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthNeg}
	TestNodeOwnerStdSignatureZero     = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthZero}
	TestClientStdSignaturePos1        = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos1}
	TestClientStdSignaturePos2        = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos2}
	TestClientStdSignatureNeg         = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthNeg}
	TestClientStdSignatureZero        = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthZero}
	TestStdSignatureEmpty             = auth.StdSignature{}
)
