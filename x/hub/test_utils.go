package hub

import (
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	ssdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"

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

	lockerId  = "locker1"
	lockerId2 = "locker2"
	lockerId3 = "locker3"

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
)
