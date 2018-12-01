package ibc

import (
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmDB "github.com/tendermint/tendermint/libs/db"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var (
	privateKey = ed25519.GenPrivKey()
	pubKey     = privateKey.PubKey()
	accAddress = csdkTypes.AccAddress(pubKey.Address())
)

func DefaultSetup() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey) {
	db := tmDB.NewMemDB()
	vpnKey := csdkTypes.NewKVStoreKey("vpn")
	ibcKey := csdkTypes.NewKVStoreKey("ibc")
	hubKey := csdkTypes.NewKVStoreKey("coin_locker")
	sessionKey := csdkTypes.NewKVStoreKey("session")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(vpnKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(ibcKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(hubKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(sessionKey, csdkTypes.StoreTypeIAVL, db)

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	return ms, ibcKey
}

var msgIBCTransaction = MsgIBCTransaction{
	csdkTypes.AccAddress([]byte("acc_address")),
	0,
	sdkTypes.IBCPacket{
		"src_chain_id",
		"dest_chain_id",
		nil,
	},
}
