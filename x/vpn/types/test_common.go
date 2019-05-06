package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var (
	TestMonikerLenZero  = ""
	TestMonikerValid    = "MONIKER"
	TestMonikerLenGT128 = "MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_" +
		"MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_MONIKER_"

	TestCoinPos   = csdkTypes.NewInt64Coin("sent", 10)
	TestCoinNeg   = csdkTypes.Coin{"sent", csdkTypes.NewInt(-10)}
	TestCoinZero  = csdkTypes.NewInt64Coin("sent", 0)
	TestCoinEmpty = csdkTypes.NewInt64Coin("empty", 0)
	TestCoinNil   = csdkTypes.Coin{}

	TestCoinsPos     = csdkTypes.Coins{TestCoinPos, csdkTypes.NewInt64Coin("sut", 100)}
	TestCoinsNeg     = csdkTypes.Coins{TestCoinNeg, csdkTypes.Coin{"sut", csdkTypes.NewInt(-100)}}
	TestCoinsZero    = csdkTypes.Coins{TestCoinZero, csdkTypes.NewInt64Coin("sut", 0)}
	TestCoinsInvalid = csdkTypes.Coins{csdkTypes.NewInt64Coin("sut", 100), TestCoinZero}
	TestCoinsEmpty   = csdkTypes.Coins{}

	TestPrivKey1 = ed25519.GenPrivKey()
	TestPrivKey2 = ed25519.GenPrivKey()

	TestPubkey1 = TestPrivKey1.PubKey()
	TestPubkey2 = TestPrivKey2.PubKey()

	TestAddress1 = csdkTypes.AccAddress(TestPubkey1.Address())
	TestAddress2 = csdkTypes.AccAddress(TestPubkey2.Address())

	TestAddressEmpty = csdkTypes.AccAddress([]byte(""))

	TestUploadNeg  = csdkTypes.NewInt(-1000000000)
	TestUploadZero = csdkTypes.NewInt(0)
	TestUploadPos1 = csdkTypes.NewInt(1000000000)
	TestUploadPos2 = TestUploadPos1.Mul(csdkTypes.NewInt(2))

	TestDownloadNeg  = csdkTypes.NewInt(-1000000000)
	TestDownloadZero = csdkTypes.NewInt(0)
	TestDownloadPos1 = csdkTypes.NewInt(1000000000)
	TestDownloadPos2 = TestDownloadPos1.Mul(csdkTypes.NewInt(2))

	TestBandwidthNeg  = sdkTypes.NewBandwidth(TestUploadNeg, TestDownloadNeg)
	TestBandwidthZero = sdkTypes.NewBandwidth(TestUploadZero, TestDownloadZero)
	TestBandwidthPos1 = sdkTypes.NewBandwidth(TestUploadPos1, TestDownloadPos1)
	TestBandwidthPos2 = sdkTypes.NewBandwidth(TestUploadPos2, TestDownloadPos2)

	TestSessionBandwidthValid = SessionBandwidthInfo{TestBandwidthPos1, TestBandwidthZero,
		TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthPos1, 0}
	TestSessionBandwidthValidEqual = SessionBandwidthInfo{TestBandwidthPos1, TestBandwidthPos1,
		TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthPos1, 0}
	TestSessionBandwidthInvalid = SessionBandwidthInfo{TestBandwidthPos1, TestBandwidthZero,
		TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthPos1, 0}

	TestAPIPortValid   = uint16(8000)
	TestAPIPortInvalid = uint16(0)

	TestEncryptionMethod = "encryption-method"
	TestNodeType         = "node_type"
	TestVersion          = "version"

	TestNodeIDValid   = sdkTypes.NewID(fmt.Sprintf("%s/%d", TestAddress1.String(), 0))
	TestNodeIDInvalid = sdkTypes.NewID("invalid")
	TestNodeIDEmpty   = sdkTypes.NewID("")

	TestStatusInvalid = "invalid"

	TestSessionIDValid   = sdkTypes.NewID(fmt.Sprintf("%s/%d", TestAddress2.String(), 0))
	TestSessionIDInvalid = sdkTypes.NewID("invalid")
	TestSessionIDEmpty   = sdkTypes.NewID("")

	TestBandWidthSignDataPos1 = sdkTypes.NewBandwidthSignData(TestSessionIDValid, TestBandwidthPos1,
		TestAddress1, TestAddress2)
	TestBandWidthSignDataPos2 = sdkTypes.NewBandwidthSignData(TestSessionIDValid, TestBandwidthPos2,
		TestAddress1, TestAddress2)
	TestBandWidthSignDataNeg = sdkTypes.NewBandwidthSignData(TestSessionIDValid, TestBandwidthNeg,
		TestAddress1, TestAddress2)
	TestBandWidthSignDataZero = sdkTypes.NewBandwidthSignData(TestSessionIDValid, TestBandwidthZero,
		TestAddress1, TestAddress2)

	TestNodeOwnerSignBandWidthPos1, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos1.GetBytes())
	TestNodeOwnerSignBandWidthPos2, _ = TestPrivKey1.Sign(TestBandWidthSignDataPos2.GetBytes())
	TestNodeOwnerSignBandWidthNeg, _  = TestPrivKey1.Sign(TestBandWidthSignDataNeg.GetBytes())
	TestNodeOwnerSignBandWidthZero, _ = TestPrivKey1.Sign(TestBandWidthSignDataZero.GetBytes())

	TestClientSignBandWidthPos1, _ = TestPrivKey2.Sign(TestBandWidthSignDataPos1.GetBytes())
	TestClientSignBandWidthPos2, _ = TestPrivKey2.Sign(TestBandWidthSignDataPos2.GetBytes())
	TestClientSignBandWidthNeg, _  = TestPrivKey2.Sign(TestBandWidthSignDataNeg.GetBytes())
	TestClientSignBandWidthZero, _ = TestPrivKey2.Sign(TestBandWidthSignDataZero.GetBytes())
)

var (
	TestNodeValid = Node{
		ID:                      TestNodeIDValid,
		Owner:                   TestAddress1,
		OwnerPubKey:             TestPubkey1,
		LockedAmount:            TestCoinPos,
		PricesPerGB:             TestCoinsPos,
		NetSpeed:                TestBandwidthPos1,
		APIPort:                 TestAPIPortValid,
		EncryptionMethod:        TestEncryptionMethod,
		Type:                    TestNodeType,
		Version:                 TestVersion,
		Status:                  StatusRegistered,
		StatusModifiedAtHeight:  0,
		DetailsModifiedAtHeight: 0,
	}
	TestNodeEmpty     = Node{}
	TestNodeIDsEmpty  = sdkTypes.IDs(nil)
	TestNodeIDsValid  = sdkTypes.IDs{TestNodeIDValid, TestNodeIDValid}
	TestNodesEmpty    = []*Node(nil)
	TestNodeTagsValid = csdkTypes.EmptyTags().AppendTag("node_id", TestNodeIDValid.String())

	TestSessionValid = Session{
		ID:              TestSessionIDValid,
		NodeID:          TestNodeIDValid,
		NodeOwner:       TestAddress1,
		NodeOwnerPubKey: TestPubkey1,
		Client:          TestAddress2,
		ClientPubKey:    TestPubkey2,
		LockedAmount:    TestCoinPos,
		PricePerGB:      TestCoinPos,
		BandwidthInfo: SessionBandwidthInfo{
			ToProvide:        TestBandwidthPos2,
			Consumed:         TestBandwidthZero,
			NodeOwnerSign:    nil,
			ClientSign:       nil,
			ModifiedAtHeight: 0,
		},
		Status:                 StatusInit,
		StatusModifiedAtHeight: 0,
		StartedAtHeight:        0,
		EndedAtHeight:          0,
	}
	TestSessionEmpty     = Session{}
	TestSessionIDsEmpty  = sdkTypes.IDs(nil)
	TestSessionIDsValid  = sdkTypes.IDs{TestSessionIDValid, TestSessionIDValid}
	TestSessionsEmpty    = []*Session(nil)
	TestSessionTagsValid = csdkTypes.EmptyTags().AppendTag("session_id", TestSessionIDValid.String())
)
