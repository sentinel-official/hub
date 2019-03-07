package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	testPrivKey1 = ed25519.GenPrivKeyFromSecret([]byte("priv_key_1"))
	testPubKey1  = testPrivKey1.PubKey()
	testAddress1 = csdkTypes.AccAddress(testPubKey1.Address())

	testPrivKey2 = ed25519.GenPrivKeyFromSecret([]byte("priv_key_2"))
	testPubKey2  = testPrivKey2.PubKey()
	testAddress2 = csdkTypes.AccAddress(testPubKey2.Address())

	testUpload    = csdkTypes.NewInt(1000000000)
	testDownload  = csdkTypes.NewInt(1000000000)
	testBandwidth = NewBandwidth(testUpload, testDownload)

	testID = IDFromOwnerAndCount(testAddress1, 0)
)
