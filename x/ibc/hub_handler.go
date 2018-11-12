package ibc

import (
	"fmt"
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
)

func NewHubHandler(ibc Keeper, hubKeeper hub.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case hub.MsgLockCoins:
				return handleLockCoins(ctx, ibc, hubKeeper, msg.IBCPacket)
			case hub.MsgReleaseCoins:
				return handleReleaseCoins(ctx, ibc, hubKeeper, msg.IBCPacket)
			case hub.MsgReleaseCoinsToMany:
				return handleReleaseCoinsToMany(ctx, ibc, hubKeeper, msg.IBCPacket)
			default:
				errMsg := fmt.Sprintf("Unrecognized IBC Msg: %v", reflect.TypeOf(ibcMsg))
				return csdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", msg.Type())
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleLockCoins(ctx csdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgLockCoins)
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker != nil {
		return csdkTypes.Result{}
	}

	hubKeeper.LockCoins(ctx, msg.LockerId, msg.Address, msg.Coins)

	packet := sdkTypes.IBCPacket{
		SrcChainId:  ibcPacket.DestChainId,
		DestChainId: ibcPacket.SrcChainId,
		Message: hub.MsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  msg.Address,
			Coins:    msg.Coins,
			Locked:   true,
		},
	}

	if err := ibc.PostIBCPacket(ctx, packet); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}

func handleReleaseCoins(ctx csdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgReleaseCoins)
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return csdkTypes.Result{}
	}

	hubKeeper.ReleaseCoins(ctx, msg.LockerId)

	packet := sdkTypes.IBCPacket{
		SrcChainId:  ibcPacket.DestChainId,
		DestChainId: ibcPacket.SrcChainId,
		Message: hub.MsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  locker.Address,
			Coins:    locker.Coins,
			Locked:   false,
		},
	}

	if err := ibc.PostIBCPacket(ctx, packet); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}

func handleReleaseCoinsToMany(ctx csdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgReleaseCoinsToMany)
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return csdkTypes.Result{}
	}

	hubKeeper.ReleaseCoinsToMany(ctx, msg.LockerId, msg.Addresses, msg.Shares)

	packet := sdkTypes.IBCPacket{
		SrcChainId:  ibcPacket.DestChainId,
		DestChainId: ibcPacket.SrcChainId,
		Message: hub.MsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  locker.Address,
			Coins:    locker.Coins,
			Locked:   false,
		},
	}

	if err := ibc.PostIBCPacket(ctx, packet); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}
