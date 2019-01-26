package vpn

import (
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func NewHandler(nk Keeper, bk bank.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgRegisterNode:
			return handleRegisterNode(ctx, nk, bk, msg)
		default:
			return errorUnknownMsgType(reflect.TypeOf(msg).Name()).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, nk Keeper, bk bank.Keeper, msg MsgRegisterNode) csdkTypes.Result {
	details := sdkTypes.VPNNodeDetails{
		Owner:        msg.Owner,
		LockedAmount: msg.AmountToLock,
		APIPort:      msg.APIPort,
		NetSpeed:     msg.NetSpeed,
		EncMethod:    msg.EncMethod,
		PerGBAmount:  msg.PerGBAmount,
		Version:      msg.Version,
	}

	tags, err := AddVPN(ctx, nk, bk, &details)
	if err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: tags}
}
