package hub

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgLockCoins{}, "lock_coins", nil)
	cdc.RegisterConcrete(MsgUnlockCoins{}, "unlock_coins", nil)
	cdc.RegisterConcrete(MsgSplitUnlockCoins{}, "split_unlock_coins", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
