package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto/ed25519"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var (
	TestMonikerLenZero     = ""
	TestMonikerValid       = "MONIKER"
	TestMonikerLengthGT128 = "MONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKER" +
		"MONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKERMONIKER"

	TestCoinPos   = csdkTypes.NewInt64Coin("stake", 10)
	TestCoinNeg   = csdkTypes.Coin{"stake", csdkTypes.NewInt(-10)}
	TestCoinZero  = csdkTypes.NewInt64Coin("stake", 0)
	TestCoinEmpty = csdkTypes.NewInt64Coin("empty", 0)
	TestCoinNil   = csdkTypes.Coin{}

	TestCoinsPos     = csdkTypes.Coins{TestCoinPos}
	TestCoinsNeg     = csdkTypes.Coins{TestCoinNeg, csdkTypes.Coin{"stake", csdkTypes.NewInt(-100)}}
	TestCoinsZero    = csdkTypes.Coins{TestCoinZero, csdkTypes.NewInt64Coin("stake", 0)}
	TestCoinsInvalid = csdkTypes.Coins{csdkTypes.NewInt64Coin("stake", 100), TestCoinZero}
	TestCoinsEmpty   = csdkTypes.Coins{}

	TestPrivKey1 = ed25519.GenPrivKey()
	TestPrivKey2 = ed25519.GenPrivKey()

	TestPubkey1 = TestPrivKey1.PubKey()
	TestPubkey2 = TestPrivKey2.PubKey()

	TestAddress1 = csdkTypes.AccAddress(TestPubkey1.Address())
	TestAddress2 = csdkTypes.AccAddress(TestPubkey2.Address())

	TestAddressEmpty = csdkTypes.AccAddress([]byte(""))

	TestUploadNeg  = csdkTypes.NewInt(-1000000)
	TestUploadZero = csdkTypes.NewInt(0)
	TestUploadPos1 = csdkTypes.NewInt(1000000)
	TestUploadPos2 = TestUploadPos1.Mul(csdkTypes.NewInt(2))

	TestDownloadNeg  = csdkTypes.NewInt(-1000000000)
	TestDownloadZero = csdkTypes.NewInt(0)
	TestDownloadPos1 = csdkTypes.NewInt(1000000000)
	TestDownloadPos2 = TestDownloadPos1.Mul(csdkTypes.NewInt(2))

	TestBandwidthNeg  = sdkTypes.NewBandwidth(TestUploadNeg, TestDownloadNeg)
	TestBandwidthZero = sdkTypes.NewBandwidth(TestUploadZero, TestDownloadZero)
	TestBandwidthPos1 = sdkTypes.NewBandwidth(TestUploadPos1, TestDownloadPos1)
	TestBandwidthPos2 = sdkTypes.NewBandwidth(TestUploadPos2, TestDownloadPos2)

	TestEncryption    = "encryption"
	TestNodeType      = "node_type"
	TestVersion       = "version"
	TestStatusInvalid = "invalid"

	TestIDPos    = sdkTypes.NewIDFromUInt64(1)
	TestIDZero   = sdkTypes.NewIDFromUInt64(0)
	TestIDsEmpty = sdkTypes.IDs(nil)
	TestIDsValid = sdkTypes.IDs{TestIDPos}
)
var (
	TestBandWidthSignDataPos1 = NewBandwidthSignatureData(TestIDZero, 1, TestBandwidthPos1)
	TestBandWidthSignDataPos2 = NewBandwidthSignatureData(TestIDPos, 2, TestBandwidthPos2)
	TestBandWidthSignDataNeg  = NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthNeg)
	TestBandWidthSignDataZero = NewBandwidthSignatureData(TestIDPos, 0, TestBandwidthZero)

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
