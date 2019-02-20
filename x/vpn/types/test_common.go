package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/ironman0x7b2/sentinel-sdk/types"
)

var (
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
	TestUploadPos  = csdkTypes.NewInt(1000000000)

	TestDownloadNeg  = csdkTypes.NewInt(-1000000000)
	TestDownloadZero = csdkTypes.NewInt(0)
	TestDownloadPos  = csdkTypes.NewInt(1000000000)

	TestBandwidthNeg  = types.NewBandwidth(TestUploadNeg, TestDownloadNeg)
	TestBandwidthZero = types.NewBandwidth(TestUploadZero, TestDownloadZero)
	TestBandwidthPos  = types.NewBandwidth(TestUploadPos, TestDownloadPos)

	TestAPIPortValid   = NewAPIPort(8000)
	TestAPIPortInvalid = NewAPIPort(0)

	TestEncMethod = "enc_method"
	TestNodeType  = "node_type"
	TestVersion   = "version"

	TestNodeIDValid   = NewNodeID(fmt.Sprintf("%s/%d", TestAddress1.String(), 0))
	TestNodeIDInvalid = NewNodeID("invalid")
	TestNodeIDEmpty   = NewNodeID("")

	TestStatusInvalid = "invalid"

	TestSessionIDValid   = NewSessionID(fmt.Sprintf("%s/%d", TestAddress2.String(), 0))
	TestSessionIDInvalid = NewSessionID("invalid")
	TestSessionIDEmpty   = NewSessionID("")

	TestClientSign    = []byte("client_sign")
	TestNodeOwnerSign = []byte("node_owner_sign")
)
