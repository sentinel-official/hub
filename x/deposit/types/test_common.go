// nolint
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	TestPrivKey1 = ed25519.GenPrivKey()
	TestPrivKey2 = ed25519.GenPrivKey()
	TestPubKey1  = TestPrivKey1.PubKey()
	TestPubKey2  = TestPrivKey2.PubKey()
	TestAddress1 = sdk.AccAddress(TestPubKey1.Address())
	TestAddress2 = sdk.AccAddress(TestPubKey2.Address())
)
