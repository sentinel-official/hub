package vpn

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/tendermint/tendermint/crypto/ed25519"
	dbm "github.com/tendermint/tendermint/libs/db"
	"strconv"
)

var count = uint64(0)

var addr1 = csdkTypes.AccAddress([]byte("some-address1"))
var addr2 = csdkTypes.AccAddress([]byte("some-address2"))

var coin1 = csdkTypes.NewCoin("SENT", csdkTypes.NewInt(100))
var coin2 = csdkTypes.NewCoin("sent", csdkTypes.NewInt(100))

var pvk1 = ed25519.GenPrivKey()

var pubKey1 = pvk1.PubKey()

var vpnRoute = "vpn"
var sessionRoute = "session"

var vpnRegisterType = "msg_register_node"

func msgLockCoinsSignatureBytes() []byte {
	bytes, _ := json.Marshal(hub.MsgLockCoins{
		lockerID1,
		csdkTypes.Coins{coin1},
		pubKey1,
		nil,
	})
	return bytes
}

var sign1, _ = pvk1.Sign(msgLockCoinsSignatureBytes())

var addr3 = csdkTypes.AccAddress(pubKey1.Address())

var vpnID = addr3.String() + "/" + strconv.Itoa(int(0))

var lockerID1 = "vpn" + "/" + vpnID

var sessionID1 = addr3.String() + "/" + strconv.Itoa(int(count))

var lockerID2 = "session" + "/" + sessionID1

var status = "STATUS"

func setupMultiStore() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey) {
	db := dbm.NewMemDB()
	ibcStoreKey := csdkTypes.NewKVStoreKey("ibc")
	vpnStoreKey := csdkTypes.NewKVStoreKey("vpn")
	sessionStoreKey := csdkTypes.NewKVStoreKey("session")

	ms := store.NewCommitMultiStore(db)

	ms.MountStoreWithDB(ibcStoreKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(vpnStoreKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(sessionStoreKey, csdkTypes.StoreTypeIAVL, db)

	ms.LoadLatestVersion()
	return ms, ibcStoreKey, vpnStoreKey, sessionStoreKey
}

func TestGetVPNDetails() *sdkTypes.VPNDetails {

	return &sdkTypes.VPNDetails{
		Address: addr3,
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
		LockerID:   lockerID1,
	}

}

func TestGetSessionDetails() *sdkTypes.SessionDetails {
	return &sdkTypes.SessionDetails{
		VPNID:         vpnID,
		ClientAddress: addr1,
		GBToProvide:   10,
		PricePerGB:    100,
		Upload:        100,
		Download:      100,
		StartTime:     nil,
		EndTime:       nil,
		Status:        "LOCKED",
	}
}

func TestGetMsgRegisterNode() *MsgRegisterNode {
	return &MsgRegisterNode{
		From: addr3, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
		EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
		Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}

}

func TestGetMsgPayVpnService() *MsgPayVPNService {

	return &MsgPayVPNService{
		From:  addr3,
		VPNID: vpnID,

		LockerID:  lockerID2,
		Coins:     csdkTypes.Coins{coin1},
		PubKey:    pubKey1,
		Signature: sign1,
	}

}

func TestGetMsgUpdateSessionSTatus() *MsgUpdateSessionStatus {
	return &MsgUpdateSessionStatus{
		From:      addr3,
		SessionID: sessionID1,
		Status:    status,
	}

}

func TestGetMsgDeregisterNode() *MsgDeregisterNode {
	return &MsgDeregisterNode{
		From:      addr3,
		VPNID:     vpnID,
		LockerID:  lockerID1,
		PubKey:    pubKey1,
		Signature: sign1,
	}
}

func TestGetIBCPacket() *sdkTypes.IBCPacket {
	return &sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockerStatus{
			LockerID: lockerID1,
			Status:   "LOCKED",
		},
	}

}

func TestGetSessionIBCPacket() *sdkTypes.IBCPacket {
	return &sdkTypes.IBCPacket{
		SrcChainID:  "sentinel-vpn",
		DestChainID: "sentinel-hub",
		Message: hub.MsgLockerStatus{
			LockerID: lockerID2,
			Status:   "LOCKED",
		},
	}

}
