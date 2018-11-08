package hub

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

func handleLockCoins(ctx sdkTypes.Context, k Keeper, msg MsgLockCoins) sdkTypes.Result {
	lockId := msg.LockId
	addr := msg.Address
	coins := msg.Coins
	lockedCoins := k.GetLockedCoins(ctx, lockId)

	if lockedCoins.Address != nil {
		return sdkTypes.Result{}
	}

	k.LockCoins(ctx, lockId, addr, coins)

	ibcPacket := hubTypes.IBCPacket{
		SrcChain:  "sentinel-hub",
		DestChain: msg.ChainId,
		Msg: hubTypes.IBCMsgLockCoins{
			LockId:  lockId,
			Address: addr,
			Coins:   coins,
		},
	}

	if err := k.ibcKeeper.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return sdkTypes.Result{}
}

func handleUnlockCoins(ctx sdkTypes.Context, k Keeper, msg MsgUnlockCoins) sdkTypes.Result {
	lockId := msg.LockId
	lockedCoins := k.GetLockedCoins(ctx, lockId)

	if lockedCoins.Address == nil {
		return sdkTypes.Result{}
	}

	k.UnlockCoins(ctx, lockId)

	ibcPacket := hubTypes.IBCPacket{
		SrcChain:  "sentinel-hub",
		DestChain: msg.ChainId,
		Msg: hubTypes.IBCMsgLockCoins{
			LockId:  lockId,
			Address: lockedCoins.Address,
			Coins:   lockedCoins.Coins,
		},
	}

	if err := k.ibcKeeper.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return sdkTypes.Result{}
}

func handleSplitUnlockCoins(ctx sdkTypes.Context, k Keeper, msg MsgSplitUnlockCoins) sdkTypes.Result {
	lockId := msg.LockId
	splits := msg.Splits
	lockedCoins := k.GetLockedCoins(ctx, lockId)

	if lockedCoins.Address == nil {
		return sdkTypes.Result{}
	}

	k.SplitUnlockCoins(ctx, lockId, splits)

	ibcPacket := hubTypes.IBCPacket{
		SrcChain:  "sentinel-hub",
		DestChain: msg.ChainId,
		Msg: hubTypes.IBCMsgLockCoins{
			LockId:  lockId,
			Address: lockedCoins.Address,
			Coins:   lockedCoins.Coins,
		},
	}

	if err := k.ibcKeeper.PostIBCPacket(ctx, ibcPacket); err != nil {
		panic(err)
	}

	return sdkTypes.Result{}
}
