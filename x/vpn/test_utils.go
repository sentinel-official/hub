package vpn

import (
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var addr1 = csdkTypes.AccAddress([]byte("some-address1"))
var addr2 = csdkTypes.AccAddress([]byte("some-address2"))
var count = uint64(5)
var status = "STATUS"

func setupMultiStore() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey) {
	db := dbm.NewMemDB()
	accountStoreKey := csdkTypes.NewKVStoreKey("acc")
	vpnStoreKey := csdkTypes.NewKVStoreKey("vpn")
	sessionStoreKey := csdkTypes.NewKVStoreKey("session")

	ms := store.NewCommitMultiStore(db)

	ms.MountStoreWithDB(accountStoreKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(vpnStoreKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(sessionStoreKey, csdkTypes.StoreTypeIAVL, db)

	ms.LoadLatestVersion()
	return ms, accountStoreKey, vpnStoreKey, sessionStoreKey
}

func TestGetVPNDetails(vpnId string) *sdkTypes.VPNDetails {

	return &sdkTypes.VPNDetails{
		Address: addr1,
		APIPort: 8080,
		Location: sdkTypes.Location{
			City:      "city",
			Country:   "country",
			Latitude:  1234,
			Longitude: 4321,
		},
		NetSpeed: sdkTypes.NetSpeed{
			Upload:   1000,
			Download: 1000,
		},
		EncMethod:  "AES-256",
		PricePerGB: 100,
		Version:    "1.0",
		Status:     "REGISTERED",
		LockerID:   "vpn" + "/" + vpnId,
	}

}

func TestGetSessionDetails(vpnId string) *sdkTypes.SessionDetails {
	return &sdkTypes.SessionDetails{
		VPNID:         vpnId,
		ClientAddress: addr1,
		GBToProvide:   10,
		PricePerGB:    100,
		Upload:        100,
		Download:      100,
		StartTime:     nil,
		EndTime:       nil,
		Status:        "STARTED",
	}
}
