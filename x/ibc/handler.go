package ibc

import (
	"fmt"

	ccsdkTypes "github.com/cosmos/cosmos-sdk/types"
	csdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
)

func NewHandler(ibc Keeper, hubKeeper hub.Keeper) ccsdkTypes.Handler {
	return func(ctx ccsdkTypes.Context, msg ccsdkTypes.Msg) ccsdkTypes.Result {
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
				return ccsdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", msg.Type())
			return ccsdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleLockCoins(ctx ccsdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgLockCoins) ccsdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker != nil {
		return ccsdkTypes.Result{}
	}

	hubKeeper.LockCoins(ctx, msg.LockerId, msg.Address, msg.Coins)

	ibcPacket := csdkTypes.IBCPacket{
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

	return ccsdkTypes.Result{}
}

func handleReleaseCoins(ctx ccsdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgReleaseCoins) ccsdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return ccsdkTypes.Result{}
	}

	hubKeeper.ReleaseCoins(ctx, msg.LockerId)

	ibcPacket := csdkTypes.IBCPacket{
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

	return ccsdkTypes.Result{}
}

func handleReleaseCoinsToMany(ctx ccsdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgReleaseCoinsToMany) ccsdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return ccsdkTypes.Result{}
	}

	hubKeeper.ReleaseCoinsToMany(ctx, msg.LockerId, msg.Addresses, msg.Shares)

	ibcPacket := csdkTypes.IBCPacket{
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

	return ccsdkTypes.Result{}
}
