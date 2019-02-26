package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	TestPrivKey1 = ed25519.GenPrivKey()
	TestPubKey1  = TestPrivKey1.PubKey()
	TestAddress1 = csdkTypes.AccAddress(TestPubKey1.Address())
)
