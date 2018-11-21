package hub

import (
	"fmt"
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
)

func NewIBCHubHandler(ibcKeeper ibc.Keeper, hubKeeper Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case ibc.MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case MsgLockCoins:
				return handleLockCoins(ctx, ibcKeeper, hubKeeper, msg.IBCPacket)
			case MsgReleaseCoins:
				return handleReleaseCoins(ctx, ibcKeeper, hubKeeper, msg.IBCPacket)
			case MsgReleaseCoinsToMany:
				return handleReleaseCoinsToMany(ctx, ibcKeeper, hubKeeper, msg.IBCPacket)
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

func handleLockCoins(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(MsgLockCoins)

	lockerID := ibcPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()

	locker := hubKeeper.GetLocker(ctx, lockerID)

	if locker != nil {
		// TODO: Replace with ErrLockerAlreadyExists
		return csdkTypes.Result{}
	}

	if err := hubKeeper.LockCoins(ctx, lockerID, address, msg.Coins); err != nil {
		// TODO: Replace with ErrLockCoins
		return csdkTypes.Result{}
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcPacket.DestChainID,
		DestChainID: ibcPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   "LOCKED",
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		return csdkTypes.Result{}
	}

	// TODO: Replace with SuccessLockCoins
	return csdkTypes.Result{}
}

func handleReleaseCoins(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(MsgReleaseCoins)

	lockerID := ibcPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()

	locker := hubKeeper.GetLocker(ctx, lockerID)

	if locker == nil {
		// TODO: Replace with ErrLockerNotExists
		return csdkTypes.Result{}
	}

	if !locker.Address.Equals(address) {
		// TODO: Replace with ErrInvalidLockerOwnerAddress
		return csdkTypes.Result{}
	}

	if err := hubKeeper.ReleaseCoins(ctx, msg.LockerID); err != nil {
		// TODO: Replace with ErrReleaseCoins
		return csdkTypes.Result{}
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcPacket.DestChainID,
		DestChainID: ibcPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   "RELEASED",
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		return csdkTypes.Result{}
	}

	// TODO: Replace with SuccessReleaseCoins
	return csdkTypes.Result{}
}

func handleReleaseCoinsToMany(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(MsgReleaseCoinsToMany)

	lockerID := ibcPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()

	locker := hubKeeper.GetLocker(ctx, lockerID)

	if locker == nil {
		// TODO: Replace with ErrLockerNotExists
		return csdkTypes.Result{}
	}

	if !locker.Address.Equals(address) {
		// TODO: Replace with ErrInvalidLockerOwnerAddress
		return csdkTypes.Result{}
	}

	if err := hubKeeper.ReleaseCoinsToMany(ctx, msg.LockerID, msg.Addresses, msg.Shares); err != nil {
		// TODO: Replace with ErrReleaseCoinsToMany
		return csdkTypes.Result{}
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcPacket.DestChainID,
		DestChainID: ibcPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   "RELEASED",
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		return csdkTypes.Result{}
	}

	// TODO: Replace with SuccessReleaseCoinsToMany
	return csdkTypes.Result{}
}
