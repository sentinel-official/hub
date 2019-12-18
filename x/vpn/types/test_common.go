// nolint
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto/ed25519"

	hub "github.com/sentinel-official/hub/types"
)

var (
	TestPrivKey1 = ed25519.GenPrivKey()
	TestPrivKey2 = ed25519.GenPrivKey()
	TestPrivKey3 = ed25519.GenPrivKey()
	TestPubkey1  = TestPrivKey1.PubKey()
	TestPubkey2  = TestPrivKey2.PubKey()
	TestPubkey3  = TestPrivKey3.PubKey()
	TestAddress1 = sdk.AccAddress(TestPubkey1.Address())
	TestAddress2 = sdk.AccAddress(TestPubkey2.Address())
	TestAddress3 = sdk.AccAddress(TestPubkey3.Address())
	TestNode     = Node{
		ID:               hub.NewNodeID(0),
		Owner:            TestAddress1,
		Deposit:          sdk.NewInt64Coin("stake", 100),
		Type:             "node_type",
		Version:          "version",
		Moniker:          "moniker",
		PricesPerGB:      sdk.Coins{sdk.NewInt64Coin("stake", 100)},
		InternetSpeed:    TestBandwidthPos1,
		Encryption:       "encryption",
		Status:           StatusDeRegistered,
		StatusModifiedAt: 1,
	}
	TestSubscription = Subscription{
		ID:                 hub.NewSubscriptionID(0),
		NodeID:             hub.NewNodeID(0),
		Client:             TestAddress2,
		Resolver:           TestAddress3,
		PricePerGB:         sdk.NewInt64Coin("stake", 100),
		TotalDeposit:       sdk.NewInt64Coin("stake", 100),
		RemainingDeposit:   sdk.NewInt64Coin("stake", 100),
		RemainingBandwidth: TestBandwidthPos1,
		Status:             StatusActive,
		StatusModifiedAt:   0,
	}
	TestSession = Session{
		ID:               hub.NewSessionID(0),
		SubscriptionID:   hub.NewSubscriptionID(0),
		Bandwidth:        TestBandwidthPos1,
		Status:           StatusActive,
		StatusModifiedAt: 0,
	}
	TestResolver = Resolver{
		Owner:      TestAddress3,
		Commission: sdk.NewDecWithPrec(12, 2),
		Status:     StatusRegistered,
	}
	TestFreeClient = FreeClient{
		NodeID: hub.NewNodeID(0),
		Client: TestAddress2,
	}

	TestBandwidthNeg                  = hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000))
	TestBandwidthZero                 = hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0))
	TestBandwidthPos1                 = hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000))
	TestBandwidthPos2                 = TestBandwidthPos1.Add(TestBandwidthPos1)
	TestBandWidthSignDataNeg          = hub.NewBandwidthSignatureData(hub.NewSubscriptionID(0), 0, TestBandwidthNeg)
	TestNodeOwnerSignBandWidthNeg, _  = TestPrivKey1.Sign(TestBandWidthSignDataNeg.Bytes())
	TestNodeOwnerStdSignatureNeg      = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthNeg}
	TestClientSignBandWidthNeg, _     = TestPrivKey2.Sign(TestBandWidthSignDataNeg.Bytes())
	TestClientStdSignatureNeg         = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthNeg}
	TestBandWidthSignDataZero         = hub.NewBandwidthSignatureData(hub.NewSubscriptionID(0), 0, TestBandwidthZero)
	TestNodeOwnerSignBandWidthZero, _ = TestPrivKey1.Sign(TestBandWidthSignDataZero.Bytes())
	TestNodeOwnerStdSignatureZero     = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthZero}
	TestClientSignBandWidthZero, _    = TestPrivKey2.Sign(TestBandWidthSignDataZero.Bytes())
	TestClientStdSignatureZero        = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthZero}
	TestBandWidthSignDataPos1         = hub.NewBandwidthSignatureData(hub.NewSubscriptionID(0), 1, TestBandwidthPos1)
	TestNodeOwnerSignBandWidthPos1, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos1.Bytes())
	TestNodeOwnerStdSignaturePos1     = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos1}
	TestClientSignBandWidthPos1, _    = TestPrivKey2.Sign(TestBandWidthSignDataPos1.Bytes())
	TestClientStdSignaturePos1        = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos1}
	TestBandWidthSignDataPos2         = hub.NewBandwidthSignatureData(hub.NewSubscriptionID(0), 1, TestBandwidthPos2)
	TestNodeOwnerSignBandWidthPos2, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos2.Bytes())
	TestNodeOwnerStdSignaturePos2     = auth.StdSignature{PubKey: TestPubkey1, Signature: TestNodeOwnerSignBandWidthPos2}
	TestClientSignBandWidthPos2, _    = TestPrivKey2.Sign(TestBandWidthSignDataPos2.Bytes())
	TestClientStdSignaturePos2        = auth.StdSignature{PubKey: TestPubkey2, Signature: TestClientSignBandWidthPos2}
)
