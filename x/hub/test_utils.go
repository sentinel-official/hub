package hub

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	ssdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/tendermint/tendermint/crypto/ed25519"
	dbm "github.com/tendermint/tendermint/libs/db"
)

func setupMultiStore() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey) {
	db := dbm.NewMemDB()
	authKey := csdkTypes.NewKVStoreKey("acc")
	hubKey := csdkTypes.NewKVStoreKey("hub")
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(hubKey, csdkTypes.StoreTypeIAVL, db)
	ms.LoadLatestVersion()
	return ms, authKey, hubKey
}

var (
	addr1 = csdkTypes.AccAddress("sentinel-hub")
	addr2 = csdkTypes.AccAddress("sentinel-vpn")
	addr3 = csdkTypes.AccAddress("sentinel-ibc")

	coins1   = csdkTypes.Coins{csdkTypes.NewInt64Coin("sut", 100), csdkTypes.NewInt64Coin("cdex", 120)}
	coins2   = csdkTypes.Coins{csdkTypes.NewInt64Coin("sut", 10)}
	coinsNeg = csdkTypes.Coins{csdkTypes.NewInt64Coin("sut", -1)}
	coin1    = csdkTypes.NewCoin("sut", csdkTypes.NewInt(90))
	coin2    = csdkTypes.NewCoin("cdex", csdkTypes.NewInt(120))
	coin3    = csdkTypes.NewCoin("sent", csdkTypes.NewInt(120))

	lockerId      = "locker1"
	lockerId2     = "locker2"
	lockerId3     = "locker3"
	emptyLockerId = ""

	locker1 = &ssdkTypes.CoinLocker{
		Address: addr1,
		Coins:   coins1,
		Status:  "LOCKED",
	}

	emptyLocker1 = &ssdkTypes.CoinLocker{
		Address: addr2,
		Coins:   coinsNeg,
		Status:  "LOCKED",
	}

	pvk1 = ed25519.GenPrivKey()
	pvk2 = ed25519.GenPrivKey()
	pvk3 = ed25519.GenPrivKey()

	pk1         = pvk1.PubKey()
	pk2         = pvk2.PubKey()
	pk3         = pvk3.PubKey()
	emptyPubKey crypto.PubKey

	sign1, _ = pvk1.Sign(msgLockCoinsSignatureBytes())
	sign2, _ = pvk2.Sign(msgReleaseCoinsSignatureBytes())
	sign3, _ = pvk3.Sign(msgReleaseCoinsToManySignatureBytes())
)

func msgLockCoinsSignatureBytes() []byte {
	bz, _ := json.Marshal(MsgLockCoins{
		lockerId,
		coins1,
		pk1,
		nil,
	})
	return bz
}

func msgReleaseCoinsSignatureBytes() []byte {
	bz, _ := json.Marshal(MsgReleaseCoins{
		lockerId2,
		pk2,
		nil,
	})
	return bz
}

func msgReleaseCoinsToManySignatureBytes() []byte {
	bz, _ := json.Marshal(MsgReleaseCoinsToMany{
		lockerId3,
		[]csdkTypes.AccAddress{csdkTypes.AccAddress(pk1.Address()), csdkTypes.AccAddress(pk2.Address())},
		[]csdkTypes.Coins{csdkTypes.Coins{coin1}},
		pk3,
		nil,
	})
	return bz
}