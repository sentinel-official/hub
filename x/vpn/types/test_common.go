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
	TestPubkey1  = TestPrivKey1.PubKey()
	TestPubkey2  = TestPrivKey2.PubKey()
	TestAddress1 = sdk.AccAddress(TestPubkey1.Address())
	TestAddress2 = sdk.AccAddress(TestPubkey2.Address())
	TestNode     = Node{
		ID:               hub.NewIDFromUInt64(0),
		Owner:            TestAddress1,
		Deposit:          sdk.NewInt64Coin("stake", 100),
		Type:             "node_type",
		Version:          "version",
		Moniker:          "moniker",
		PricesPerGB:      sdk.Coins{sdk.NewInt64Coin("stake", 100)},
		InternetSpeed:    hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)),
		Encryption:       "encryption",
		Status:           StatusInactive,
		StatusModifiedAt: 1,
	}
	TestSubscription = Subscription{
		ID:                 hub.NewIDFromUInt64(0),
		NodeID:             hub.NewIDFromUInt64(0),
		Client:             TestAddress2,
		PricePerGB:         sdk.NewInt64Coin("stake", 100),
		TotalDeposit:       sdk.NewInt64Coin("stake", 100),
		RemainingDeposit:   sdk.NewInt64Coin("stake", 100),
		RemainingBandwidth: hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)),
		Status:             StatusActive,
		StatusModifiedAt:   0,
	}
	TestSession = Session{
		ID:               hub.NewIDFromUInt64(0),
		SubscriptionID:   hub.NewIDFromUInt64(0),
		Bandwidth:        hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)),
		Status:           StatusActive,
		StatusModifiedAt: 0,
	}
	bandWidthSignDataPos1         = NewBandwidthSignatureData(hub.NewIDFromUInt64(0), 1, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)))
	nodeOwnerSignBandWidthPos1, _ = TestPrivKey1.Sign(bandWidthSignDataPos1.Bytes())
	nodeOwnerStdSignaturePos1     = auth.StdSignature{PubKey: TestPubkey1, Signature: nodeOwnerSignBandWidthPos1}
	clientSignBandWidthPos1, _    = TestPrivKey2.Sign(bandWidthSignDataPos1.Bytes())
	clientStdSignaturePos1        = auth.StdSignature{PubKey: TestPubkey2, Signature: clientSignBandWidthPos1}
)
