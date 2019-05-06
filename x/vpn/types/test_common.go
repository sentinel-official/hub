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
	TestMonikerLenGT128 = "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM" +
		                  "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM"

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

	TestBandwidthNeg  = sdkTypes.NewBandwidth(TestUploadNeg, TestDownloadNeg)
	TestBandwidthZero = sdkTypes.NewBandwidth(TestUploadZero, TestDownloadZero)
	TestBandwidthPos  = sdkTypes.NewBandwidth(TestUploadPos, TestDownloadPos)

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

	TestClientSign    = []byte("client_sign")
	TestNodeOwnerSign = []byte("node_owner_sign")
)
