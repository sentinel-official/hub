package vpn

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmDB "github.com/tendermint/tendermint/libs/db"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func setupMultiStore() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey) {
	db := tmDB.NewMemDB()

	ibcKey := csdkTypes.NewKVStoreKey(sdkTypes.KeyIBC)
	vpnKey := csdkTypes.NewKVStoreKey(sdkTypes.KeyVPN)
	sessionKey := csdkTypes.NewKVStoreKey(sdkTypes.KeySession)

	ms := store.NewCommitMultiStore(db)

	ms.MountStoreWithDB(ibcKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(vpnKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(sessionKey, csdkTypes.StoreTypeIAVL, db)

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	return ms, ibcKey, vpnKey, sessionKey
}

var (
	privKey1 = ed25519.GenPrivKey()
	privKey2 = ed25519.GenPrivKey()

	pubKey1 = privKey1.PubKey()
	pubKey2 = privKey2.PubKey()

	accAddress1 = csdkTypes.AccAddress(pubKey1.Address())
	accAddress2 = csdkTypes.AccAddress(pubKey2.Address())

	now        = time.Now().UTC()
	vpnDetails = sdkTypes.VPNDetails{
		accAddress1,
		3000,
		sdkTypes.Location{0, 0, "city", "country"},
		sdkTypes.NetSpeed{1024, 1024},
		"enc_method",
		0,
		"version",
		sdkTypes.StatusRegister,
		"locker_id",
	}
	sessionDetails = sdkTypes.SessionDetails{
		"vpn_id",
		accAddress1,
		1024,
		0,
		1024,
		1024,
		&now,
		&now,
		sdkTypes.StatusStart,
	}

	msgRegisterNode = NewMsgRegisterNode(
		accAddress1,
		3000,
		0, 0, "city", "country",
		1024, 1024,
		"enc_method",
		0,
		"version",
		sdkTypes.KeyVPN+"/"+accAddress1.String()+"/"+strconv.Itoa(0),
		csdkTypes.Coins{coin(100, "x")},
		pubKey1,
		getMsgLockCoinsSignature(sdkTypes.KeyVPN+"/"+accAddress1.String()+"/"+strconv.Itoa(0), csdkTypes.Coins{coin(100, "x")}),
	)
	msgPayVPNService = NewMsgPayVPNService(
		accAddress1,
		accAddress1.String()+"/"+strconv.Itoa(0),
		sdkTypes.KeySession+"/"+accAddress1.String()+"/"+strconv.Itoa(0),
		csdkTypes.Coins{coin(100, "x")},
		pubKey1,
		getMsgLockCoinsSignature(sdkTypes.KeySession+"/"+accAddress1.String()+"/"+strconv.Itoa(0), csdkTypes.Coins{coin(100, "x")}),
	)
	msgUpdateSessionStatus = NewMsgUpdateSessionStatus(
		accAddress1,
		"session_id",
		sdkTypes.StatusInactive,
	)
	msgDeregisterNode = NewMsgDeregisterNode(
		accAddress1,
		accAddress1.String()+"/"+strconv.Itoa(0),
		sdkTypes.KeyVPN+"/"+accAddress1.String()+"/"+strconv.Itoa(0),
		pubKey1,
		getMsgReleaseCoinsSignature(sdkTypes.KeyVPN+"/"+accAddress1.String()+"/"+strconv.Itoa(0)),
	)

	msgIBCTransaction = ibc.MsgIBCTransaction{
		accAddress2,
		0,
		sdkTypes.IBCPacket{
			"src_chain_id",
			"dest_chain_id",
			hub.MsgLockerStatus{
				sdkTypes.KeyVPN + "/" + accAddress1.String() + "/" + strconv.Itoa(0),
				sdkTypes.StatusLock,
			},
		},
	}
)

func coin(value int64, name string) csdkTypes.Coin {
	return csdkTypes.Coin{name, csdkTypes.NewInt(value)}
}

func getMsgLockCoinsSignature(lockerID string, coins csdkTypes.Coins) []byte {
	msg := hub.MsgLockCoins{
		LockerID: lockerID,
		Coins:    coins,
		PubKey:   pubKey1,
	}

	sign, err := privKey1.Sign(msg.GetUnSignBytes())

	if err != nil {
		panic(err)
	}

	return sign
}

func getMsgReleaseCoinsSignature(lockerID string) []byte {
	msg := hub.MsgLockCoins{
		LockerID: lockerID,
		PubKey:   pubKey1,
	}

	sign, err := privKey1.Sign(msg.GetUnSignBytes())

	if err != nil {
		panic(err)
	}

	return sign
}
