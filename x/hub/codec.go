package hub

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCoinLocker{}, "x/hub/msg_coin_locker", nil)

	cdc.RegisterConcrete(MsgLockCoins{}, "x/hub/msg_lock_coins", nil)
	cdc.RegisterConcrete(MsgReleaseCoins{}, "x/hub/msg_release_coins", nil)
	cdc.RegisterConcrete(MsgReleaseCoinsToMany{}, "x/hub/msg_release_coins_to_many", nil)

	cdc.RegisterConcrete(IBCMsgCoinLocker{}, "x/hub/ibc_msg_coin_locker", nil)
	cdc.RegisterConcrete(IBCMsgLockCoins{}, "x/hub/ibc_msg_lock_coins", nil)
	cdc.RegisterConcrete(IBCMsgReleaseCoins{}, "x/hub/ibc_msg_release_coins", nil)
	cdc.RegisterConcrete(IBCMsgReleaseCoinsToMany{}, "x/hub/ibc_msg_release_coins_to_many", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
