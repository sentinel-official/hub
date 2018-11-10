package ibc

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(IBCMsgCoinLocker{}, "x/ibc/ibc_msg_coin_locker", nil)
	cdc.RegisterConcrete(IBCMsgLockCoins{}, "x/ibc/ibc_msg_lock_coins", nil)
	cdc.RegisterConcrete(IBCMsgReleaseCoins{}, "x/ibc/ibc_msg_release_coins", nil)
	cdc.RegisterConcrete(IBCMsgReleaseCoinsToMany{}, "x/ibc/ibc_msg_release_coins_to_many", nil)

	cdc.RegisterConcrete(MsgIBCTransaction{}, "x/ibc/msg_ibc_transaction", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
