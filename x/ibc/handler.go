package ibc

import (
	"fmt"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
	"github.com/ironman0x7b2/sentinel-hub/x/hub"
)

func NewHandler(ibc Keeper, hubKeeper hub.Keeper) sdkTypes.Handler {
	return func(ctx sdkTypes.Context, msg sdkTypes.Msg) sdkTypes.Result {
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
				return sdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", msg.Type())
			return sdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleLockCoins(ctx sdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgLockCoins) sdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker != nil {
		return sdkTypes.Result{}
	}

	hubKeeper.LockCoins(ctx, msg.LockerId, msg.Address, msg.Coins)

	ibcPacket := hubTypes.IBCPacket{
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

	return sdkTypes.Result{}
}

func handleReleaseCoins(ctx sdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgReleaseCoins) sdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return sdkTypes.Result{}
	}

	hubKeeper.ReleaseCoins(ctx, msg.LockerId)

	ibcPacket := hubTypes.IBCPacket{
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

	return sdkTypes.Result{}
}

func handleReleaseCoinsToMany(ctx sdkTypes.Context, ibc Keeper, hubKeeper hub.Keeper, ibcMsg IBCMsgReleaseCoinsToMany) sdkTypes.Result {
	msg := ibcMsg.Message
	locker := hubKeeper.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return sdkTypes.Result{}
	}

	hubKeeper.ReleaseCoinsToMany(ctx, msg.LockerId, msg.Addresses, msg.Shares)

	ibcPacket := hubTypes.IBCPacket{
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

	return sdkTypes.Result{}
}
