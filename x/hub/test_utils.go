package hub

import (
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	dbm "github.com/tendermint/tendermint/libs/db"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func setupMultiStore() (csdkTypes.MultiStore, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey, *csdkTypes.KVStoreKey) {
	db := dbm.NewMemDB()

	accKey := csdkTypes.NewKVStoreKey(sdkTypes.KeyAccount)
	ibcKey := csdkTypes.NewKVStoreKey(sdkTypes.KeyIBC)
	coinLockerKey := csdkTypes.NewKVStoreKey(sdkTypes.KeyCoinLocker)

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(accKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(ibcKey, csdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(coinLockerKey, csdkTypes.StoreTypeIAVL, db)

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	return ms, accKey, ibcKey, coinLockerKey
}

var (
	privKey1 = ed25519.GenPrivKey()
	privKey2 = ed25519.GenPrivKey()
	privKey3 = ed25519.GenPrivKey()

	pubKey1 = privKey1.PubKey()
	pubKey2 = privKey2.PubKey()
	pubKey3 = privKey3.PubKey()

	accAddress1 = csdkTypes.AccAddress(pubKey1.Address())
	accAddress2 = csdkTypes.AccAddress(pubKey2.Address())
	accAddress3 = csdkTypes.AccAddress(pubKey3.Address())
)

func coin(value int64, name string) csdkTypes.Coin {
	return csdkTypes.Coin{name, csdkTypes.NewInt(value)}
}

func getMsgLockCoinsSignature(lockerID string, coins csdkTypes.Coins) []byte {
	msg := MsgLockCoins{
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
	msg := MsgReleaseCoins{
		LockerID: lockerID,
		PubKey:   pubKey1,
	}

	sign, err := privKey1.Sign(msg.GetUnSignBytes())

	if err != nil {
		panic(err)
	}

	return sign
}

func getMsgReleaseCoinsToManySignature(lockerID string, addresses []csdkTypes.AccAddress, shares []csdkTypes.Coins) []byte {
	msg := MsgReleaseCoinsToMany{
		LockerID:  lockerID,
		Addresses: addresses,
		Shares:    shares,
		PubKey:    pubKey1,
	}

	sign, err := privKey1.Sign(msg.GetUnSignBytes())

	if err != nil {
		panic(err)
	}

	return sign
}

func getIBCPacketMsgLockCoins(lockerID string) sdkTypes.IBCPacket {
	return sdkTypes.IBCPacket{
		"src-chain-id",
		"dest-chain-id",
		MsgLockCoins{
			lockerID,
			csdkTypes.Coins{coin(10, "x")},
			pubKey1,
			getMsgLockCoinsSignature(lockerID, csdkTypes.Coins{coin(10, "x")}),
		},
	}
}

func getIBCPacketMsgReleaseCoins(lockerID string) sdkTypes.IBCPacket {
	return sdkTypes.IBCPacket{
		"src-chain-id",
		"dest-chain-id",
		MsgReleaseCoins{
			lockerID,
			pubKey1,
			getMsgReleaseCoinsSignature(lockerID),
		},
	}
}

func getIBCPacketMsgReleaseCoinsToMany(lockerID string) sdkTypes.IBCPacket {
	return sdkTypes.IBCPacket{
		"src-chain-id",
		"dest-chain-id",
		MsgReleaseCoinsToMany{
			lockerID,
			[]csdkTypes.AccAddress{accAddress1, accAddress2, accAddress3},
			[]csdkTypes.Coins{{coin(2, "x")}, {coin(2, "x")}, {coin(2, "x")}},
			pubKey1,
			getMsgReleaseCoinsToManySignature(lockerID,
				[]csdkTypes.AccAddress{accAddress1, accAddress2, accAddress3},
				[]csdkTypes.Coins{{coin(2, "x")}, {coin(2, "x")}, {coin(2, "x")}}),
		},
	}
}
