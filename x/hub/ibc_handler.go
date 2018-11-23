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
				return handleLockCoins(ctx, ibcKeeper, hubKeeper, msg)
			case MsgReleaseCoins:
				return handleReleaseCoins(ctx, ibcKeeper, hubKeeper, msg)
			case MsgReleaseCoinsToMany:
				return handleReleaseCoinsToMany(ctx, ibcKeeper, hubKeeper, msg)
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

func handleLockCoins(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(MsgLockCoins)
	sequence := ibcKeeper.GetIngressLength(ctx, string(ibc.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID)))

	if ibcMsg.Sequence != sequence {
		// TODO: Replace with ErrInvalidIBCSequence
		panic("ibcmsg.sequence != sequence")
	}

	if !msg.Verify() {
		// TODO: Replace with ErrIBCPacketMsgVerificationFailed
		panic("!msg.verify()")
	}

	lockerID := ibcMsg.IBCPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()
	locker := hubKeeper.GetLocker(ctx, lockerID)

	if locker != nil {
		// TODO: Replace with ErrLockerAlreadyExists
		panic("locker != nil")
	}

	if err := hubKeeper.LockCoins(ctx, lockerID, address, msg.Coins); err != nil {
		// TODO: Replace with ErrLockCoins
		panic(err)
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcMsg.IBCPacket.DestChainID,
		DestChainID: ibcMsg.IBCPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   "LOCKED",
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		panic(err)
	}

	ibcKeeper.SetIngressLength(ctx, string(ibc.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID)), sequence+1)

	// TODO: Replace with SuccessLockCoins
	return csdkTypes.Result{}
}

func handleReleaseCoins(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(MsgReleaseCoins)
	sequence := ibcKeeper.GetIngressLength(ctx, string(ibc.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID)))

	if ibcMsg.Sequence != sequence {
		// TODO: Replace with ErrInvalidIBCSequence
		panic("ibcmsg.sequence != sequence")
	}

	if !msg.Verify() {
		// TODO: Replace with ErrIBCPacketMsgVerificationFailed
		panic("!msg.verify()")
	}

	lockerID := ibcMsg.IBCPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()
	locker := hubKeeper.GetLocker(ctx, lockerID)

	if locker == nil {
		// TODO: Replace with ErrLockerNotExists
		panic("locker == nil")
	}

	if !locker.Address.Equals(address) {
		// TODO: Replace with ErrInvalidLockerOwnerAddress
		panic("locker.address != address")
	}

	if err := hubKeeper.ReleaseCoins(ctx, msg.LockerID); err != nil {
		// TODO: Replace with ErrReleaseCoins
		panic(err)
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcMsg.IBCPacket.DestChainID,
		DestChainID: ibcMsg.IBCPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   "RELEASED",
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		panic(err)
	}

	ibcKeeper.SetIngressLength(ctx, string(ibc.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID)), sequence+1)

	// TODO: Replace with SuccessReleaseCoins
	return csdkTypes.Result{}
}

func handleReleaseCoinsToMany(ctx csdkTypes.Context, ibcKeeper ibc.Keeper, hubKeeper Keeper, ibcMsg ibc.MsgIBCTransaction) csdkTypes.Result {
	msg, _ := ibcMsg.IBCPacket.Message.(MsgReleaseCoinsToMany)
	sequence := ibcKeeper.GetIngressLength(ctx, string(ibc.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID)))

	if ibcMsg.Sequence != sequence {
		// TODO: Replace with ErrInvalidIBCSequence
		panic("ibcmsg.sequence != sequence")
	}

	if !msg.Verify() {
		// TODO: Replace with ErrIBCPacketMsgVerificationFailed
		panic("!msg.verify()")
	}

	lockerID := ibcMsg.IBCPacket.SrcChainID + "/" + msg.LockerID
	address := msg.PubKey.Address().Bytes()
	locker := hubKeeper.GetLocker(ctx, lockerID)

	if locker == nil {
		// TODO: Replace with ErrLockerNotExists
		panic("locker == nil")
	}

	if !locker.Address.Equals(address) {
		// TODO: Replace with ErrInvalidLockerOwnerAddress
		panic("locker.address != address")
	}

	if err := hubKeeper.ReleaseCoinsToMany(ctx, msg.LockerID, msg.Addresses, msg.Shares); err != nil {
		// TODO: Replace with ErrReleaseCoinsToMany
		panic(err)
	}

	packet := sdkTypes.IBCPacket{
		SrcChainID:  ibcMsg.IBCPacket.DestChainID,
		DestChainID: ibcMsg.IBCPacket.SrcChainID,
		Message: MsgLockerStatus{
			LockerID: msg.LockerID,
			Status:   "RELEASED",
		},
	}

	if err := ibcKeeper.PostIBCPacket(ctx, packet); err != nil {
		// TODO: Replace with ErrPostIBCPacket
		panic(err)
	}

	ibcKeeper.SetIngressLength(ctx, string(ibc.IngressLengthKey(ibcMsg.IBCPacket.SrcChainID)), sequence+1)

	// TODO: Replace with SuccessReleaseCoinsToMany
	return csdkTypes.Result{}
}
