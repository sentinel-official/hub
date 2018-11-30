package ibc

import (
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/db"
)

var (
	privateKey1 = ed25519.GenPrivKey()
	pubKey1     = privateKey1.PubKey()
	accAddress1 = csdkTypes.AccAddress(pubKey1.Address())

	privateKey2 = ed25519.GenPrivKey()
	pubKey2     = privateKey2.PubKey()
	accAddress2 = csdkTypes.AccAddress(pubKey2.Address())

	privateKey3 = ed25519.GenPrivKey()
	pubKey3     = privateKey3.PubKey()
)

func DefaultSetup() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey) {
	db := db.NewMemDB()
	vpnKey := csdkTypes.NewKVStoreKey("vpn")
	ibcKey := csdkTypes.NewKVStoreKey("ibc")
	hubKey := csdkTypes.NewKVStoreKey("coin_locker")
	sessionKey := csdkTypes.NewKVStoreKey("session")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(vpnKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(ibcKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(hubKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(sessionKey, csdkTypes.StoreTypeIAVL, db)
	ms.LoadLatestVersion()

	return ms, ibcKey
}

func TestNewIBCPacket() sdkTypes.IBCPacket {
	return sdkTypes.IBCPacket{
		"sentinel-vpn",
		"sentinel-hub",
		nil,
	}

}
