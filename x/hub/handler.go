package hub

import (
	"fmt"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

func NewHandler(k Keeper) sdkTypes.Handler {
	return func(ctx sdkTypes.Context, msg sdkTypes.Msg) sdkTypes.Result {
		switch msg := msg.(type) {
		case MsgLockCoins:
			return handleLockCoins(ctx, k, msg)
		case MsgUnlockCoins:
			return handleUnlockCoins(ctx, k, msg)
		case MsgUnlockAndShareCoins:
			return handleUnlockAndShareCoins(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", msg.Type())
			return sdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleLockCoins(ctx sdkTypes.Context, k Keeper, msg MsgLockCoins) sdkTypes.Result {
	locker := k.GetLocker(ctx, msg.LockerId)

	if locker != nil {
		return sdkTypes.Result{}
	}

	k.LockCoins(ctx, msg.LockerId, msg.Address, msg.Coins)

	ibcPacket := hubTypes.IBCPacket{
		SrcChainId:  "sentinel-hub",
		DestChainId: msg.FromChainId,
		Message: hubTypes.IBCMsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  msg.Address,
			Coins:    msg.Coins,
			Locked:   true,
		},
	}

	if err := k.ibcKeeper.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return sdkTypes.Result{}
}

func handleUnlockCoins(ctx sdkTypes.Context, k Keeper, msg MsgUnlockCoins) sdkTypes.Result {
	locker := k.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return sdkTypes.Result{}
	}

	k.UnlockCoins(ctx, msg.LockerId)

	ibcPacket := hubTypes.IBCPacket{
		SrcChainId:  "sentinel-hub",
		DestChainId: msg.FromChainId,
		Message: hubTypes.IBCMsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  locker.Address,
			Coins:    locker.Coins,
			Locked:   false,
		},
	}

	if err := k.ibcKeeper.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return sdkTypes.Result{}
}

func handleUnlockAndShareCoins(ctx sdkTypes.Context, k Keeper, msg MsgUnlockAndShareCoins) sdkTypes.Result {
	locker := k.GetLocker(ctx, msg.LockerId)

	if locker == nil {
		return sdkTypes.Result{}
	}

	k.UnlockAndShareCoins(ctx, msg.LockerId, msg.Addrs, msg.Shares)

	ibcPacket := hubTypes.IBCPacket{
		SrcChainId:  "sentinel-hub",
		DestChainId: msg.FromChainId,
		Message: hubTypes.IBCMsgCoinLocker{
			LockerId: msg.LockerId,
			Address:  locker.Address,
			Coins:    locker.Coins,
			Locked:   false,
		},
	}

	if err := k.ibcKeeper.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return sdkTypes.Result{}
}
