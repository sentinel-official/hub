package hub

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func Test_handleLockCoins(t *testing.T) {
	multiStore, accountKey, ibcKey, coinLockerKey := setupMultiStore()
	ctx := csdkTypes.NewContext(multiStore, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	sdkTypes.RegisterCodec(cdc)
	RegisterCodec(cdc)

	accountKeeper := auth.NewAccountKeeper(cdc, accountKey, auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)
	hubKeeper := NewBaseKeeper(cdc, coinLockerKey, bankKeeper)

	var result csdkTypes.Result

	account1 := auth.NewBaseAccountWithAddress(accAddress1)
	account2 := auth.NewBaseAccountWithAddress(accAddress2)

	if err := account1.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account2.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account1.SetPubKey(pubKey1); err != nil {
		panic(err)
	}

	if err := account2.SetPubKey(pubKey2); err != nil {
		panic(err)
	}

	accountKeeper.SetAccount(ctx, &account1)
	accountKeeper.SetAccount(ctx, &account2)

	result = handleLockCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 0, getIBCPacketMsgLockCoins("locker_id")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handleLockCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 0, getIBCPacketMsgLockCoins("locker_id")})
	require.Equal(t, errorInvalidIBCSequence().Result(), result)

	result = handleLockCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 1, getIBCPacketMsgLockCoins("locker_id")})
	require.Equal(t, errorLockerAlreadyExists().Result(), result)
}

func Test_handleReleaseCoins(t *testing.T) {
	multiStore, accountKey, ibcKey, coinLockerKey := setupMultiStore()
	ctx := csdkTypes.NewContext(multiStore, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	sdkTypes.RegisterCodec(cdc)
	RegisterCodec(cdc)

	accountKeeper := auth.NewAccountKeeper(cdc, accountKey, auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)
	hubKeeper := NewBaseKeeper(cdc, coinLockerKey, bankKeeper)

	var result csdkTypes.Result

	account1 := auth.NewBaseAccountWithAddress(accAddress1)
	account2 := auth.NewBaseAccountWithAddress(accAddress2)

	if err := account1.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account2.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account1.SetPubKey(pubKey1); err != nil {
		panic(err)
	}

	if err := account2.SetPubKey(pubKey2); err != nil {
		panic(err)
	}

	accountKeeper.SetAccount(ctx, &account1)
	accountKeeper.SetAccount(ctx, &account2)

	result = handleLockCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 0, getIBCPacketMsgLockCoins("locker_id")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handleReleaseCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 0, getIBCPacketMsgReleaseCoins("locker_id")})
	require.Equal(t, errorInvalidIBCSequence().Result(), result)

	result = handleReleaseCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 1, getIBCPacketMsgReleaseCoins("locker_id_x")})
	require.Equal(t, errorLockerNotExists().Result(), result)

	result = handleReleaseCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 1, getIBCPacketMsgReleaseCoins("locker_id")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handleReleaseCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 2, getIBCPacketMsgReleaseCoins("locker_id")})
	require.Equal(t, errorInvalidLockerStatus().Result(), result)
}

func Test_handleReleaseCoinsToMany(t *testing.T) {
	multiStore, accountKey, ibcKey, coinLockerKey := setupMultiStore()
	ctx := csdkTypes.NewContext(multiStore, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	sdkTypes.RegisterCodec(cdc)
	RegisterCodec(cdc)

	accountKeeper := auth.NewAccountKeeper(cdc, accountKey, auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)
	hubKeeper := NewBaseKeeper(cdc, coinLockerKey, bankKeeper)

	var result csdkTypes.Result

	account1 := auth.NewBaseAccountWithAddress(accAddress1)
	account2 := auth.NewBaseAccountWithAddress(accAddress2)

	if err := account1.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account2.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account1.SetPubKey(pubKey1); err != nil {
		panic(err)
	}

	if err := account2.SetPubKey(pubKey2); err != nil {
		panic(err)
	}

	accountKeeper.SetAccount(ctx, &account1)
	accountKeeper.SetAccount(ctx, &account2)

	result = handleLockCoins(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 0, getIBCPacketMsgLockCoins("locker_id")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handleReleaseCoinsToMany(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 0, getIBCPacketMsgReleaseCoinsToMany("locker_id")})
	require.Equal(t, errorInvalidIBCSequence().Result(), result)

	result = handleReleaseCoinsToMany(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 1, getIBCPacketMsgReleaseCoinsToMany("locker_id_x")})
	require.Equal(t, errorLockerNotExists().Result(), result)

	result = handleReleaseCoinsToMany(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 1, getIBCPacketMsgReleaseCoinsToMany("locker_id")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handleReleaseCoinsToMany(ctx, ibcKeeper, hubKeeper,
		ibc.MsgIBCTransaction{accAddress2, 2, getIBCPacketMsgReleaseCoinsToMany("locker_id")})
	require.Equal(t, errorInvalidLockerStatus().Result(), result)
}

func Test_NewIBCHubHandler(t *testing.T) {
	multiStore, accountKey, ibcKey, coinLockerKey := setupMultiStore()
	ctx := csdkTypes.NewContext(multiStore, abciTypes.Header{}, false, log.NewNopLogger())

	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	sdkTypes.RegisterCodec(cdc)
	RegisterCodec(cdc)

	accountKeeper := auth.NewAccountKeeper(cdc, accountKey, auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)
	hubKeeper := NewBaseKeeper(cdc, coinLockerKey, bankKeeper)

	var result csdkTypes.Result

	account1 := auth.NewBaseAccountWithAddress(accAddress1)
	account2 := auth.NewBaseAccountWithAddress(accAddress2)

	if err := account1.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account2.SetCoins(csdkTypes.Coins{coin(100, "x")}); err != nil {
		panic(err)
	}

	if err := account1.SetPubKey(pubKey1); err != nil {
		panic(err)
	}

	if err := account2.SetPubKey(pubKey2); err != nil {
		panic(err)
	}

	accountKeeper.SetAccount(ctx, &account1)
	accountKeeper.SetAccount(ctx, &account2)

	handler := NewIBCHubHandler(ibcKeeper, hubKeeper)

	result = handler(ctx, ibc.MsgIBCTransaction{accAddress2, 0, getIBCPacketMsgLockCoins("locker_id_1")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handler(ctx, ibc.MsgIBCTransaction{accAddress2, 1, getIBCPacketMsgReleaseCoins("locker_id_1")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handler(ctx, ibc.MsgIBCTransaction{accAddress2, 2, getIBCPacketMsgLockCoins("locker_id_2")})
	require.Equal(t, csdkTypes.Result{}, result)

	result = handler(ctx, ibc.MsgIBCTransaction{accAddress2, 3, getIBCPacketMsgReleaseCoinsToMany("locker_id_2")})
	require.Equal(t, csdkTypes.Result{}, result)
}
