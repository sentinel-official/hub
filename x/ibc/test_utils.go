package ibc

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
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
	accAddress3 = csdkTypes.AccAddress(pubKey3.Address())

	emptyAccAddress csdkTypes.AccAddress

	coinsToLoad     = csdkTypes.Coins{csdkTypes.NewInt64Coin("coin_a", 100), csdkTypes.NewInt64Coin("coin_b", 10)}
	coinsToLock     = csdkTypes.Coins{csdkTypes.NewInt64Coin("coin_a", 10)}
	coninsToRemains = csdkTypes.Coins{csdkTypes.NewInt64Coin("coin_a", 90), csdkTypes.NewInt64Coin("coin_b", 10)}

	lockerID1 = "locker_id_1"
	lockerID2 = "locker_id_2"
	lockerID3 = "locker_id_3"

	sign1, _ = privateKey1.Sign(msgLockCoinsSignatureBytes())
)

func DefaultSetup() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey) {
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

	return ms, vpnKey, hubKey, ibcKey, sessionKey
}

func TestNewMsgRegisterNode() vpn.MsgRegisterNode {
	return vpn.MsgRegisterNode{
		From: accAddress1, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
		EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
		Coins: coinsToLock, PubKey: pubKey1, Signature: sign1}

}

func msgLockCoinsSignatureBytes() []byte {

	bytes, _ := json.Marshal(hub.MsgLockCoins{
		lockerID1,
		coinsToLock,
		pubKey1,
		nil,
	})
	return bytes
}
