package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto/ed25519"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

var (
	privKey1 = ed25519.GenPrivKey()
	privKey2 = ed25519.GenPrivKey()
	pubkey1  = privKey1.PubKey()
	pubkey2  = privKey2.PubKey()
	address1 = sdk.AccAddress(pubkey1.Address())
	address2 = sdk.AccAddress(pubkey2.Address())

	bandWidthSignDataPos1 = types.NewBandwidthSignatureData(hub.NewIDFromUInt64(0), 1, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)))

	nodeOwnerSignBandWidthPos1, _ = privKey1.Sign(bandWidthSignDataPos1.Bytes())
	nodeOwnerStdSignaturePos1     = auth.StdSignature{PubKey: pubkey1, Signature: nodeOwnerSignBandWidthPos1}
	clientSignBandWidthPos1, _    = privKey2.Sign(bandWidthSignDataPos1.Bytes())
	clientStdSignaturePos1        = auth.StdSignature{PubKey: pubkey2, Signature: clientSignBandWidthPos1}
)
