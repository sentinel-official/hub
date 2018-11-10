package hub

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCoinLocker{}, "coin_locker", nil)

	cdc.RegisterConcrete(MsgLockCoins{}, "lock_coins", nil)
	cdc.RegisterConcrete(MsgReleaseCoins{}, "release_coins", nil)
	cdc.RegisterConcrete(MsgReleaseCoinsToMany{}, "release_coins_to_many", nil)

}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
