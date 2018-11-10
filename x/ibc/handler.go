package ibc

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
)

func NewHandler(ibc Keeper, hubKeeper hub.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case IBCMsgLockCoins:
				return handleLockCoins(ctx, ibc, hubKeeper, ibcMsg)
			case IBCMsgReleaseCoins:
				return handleReleaseCoins(ctx, ibc, hubKeeper, ibcMsg)
			case IBCMsgReleaseCoinsToMany:
				return handleReleaseCoinsToMany(ctx, ibc, hubKeeper, ibcMsg)
			default:
				errMsg := fmt.Sprintf("Unrecognized IBC Msg : %v", ibcMsg)
				return csdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", msg.Type())
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleLockCoins(ctx csdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgLockCoins) csdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker != nil {
		return csdkTypes.Result{}
	}

	hubKeeper.LockCoins(ctx, msg.LockerId, msg.Address, msg.Coins)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  ibcMsg.DestChainId,
		DestChainId: ibcMsg.SrcChainId,
		Message: hub.MsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  msg.Address,
			Coins:    msg.Coins,
			Locked:   true,
		},
	}

	if err := ibc.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}

func handleReleaseCoins(ctx csdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgReleaseCoins) csdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return csdkTypes.Result{}
	}

	hubKeeper.ReleaseCoins(ctx, msg.LockerId)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  ibcMsg.DestChainId,
		DestChainId: ibcMsg.SrcChainId,
		Message: hub.MsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  locker.Address,
			Coins:    locker.Coins,
			Locked:   false,
		},
	}

	if err := ibc.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}

func handleReleaseCoinsToMany(ctx csdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgReleaseCoinsToMany) csdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return csdkTypes.Result{}
	}

	hubKeeper.ReleaseCoinsToMany(ctx, msg.LockerId, msg.Addresses, msg.Shares)

	ibcPacket := sdkTypes.IBCPacket{
		SrcChainId:  ibcMsg.DestChainId,
		DestChainId: ibcMsg.SrcChainId,
		Message: hub.MsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  locker.Address,
			Coins:    locker.Coins,
			Locked:   false,
		},
	}

	if err := ibc.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return csdkTypes.Result{}
}
