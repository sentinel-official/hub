package vpn

import (
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func NewHandler(vk keeper.Keeper, ak auth.AccountKeeper, bk bank.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case types.MsgRegisterNode:
			return handleRegisterNode(ctx, vk, bk, msg)
		case types.MsgUpdateNodeDetails:
			return handleUpdateNodeDetails(ctx, vk, msg)
		case types.MsgUpdateNodeStatus:
			return handleUpdateNodeStatus(ctx, vk, msg)
		case types.MsgDeregisterNode:
			return handleDeregisterNode(ctx, vk, bk, msg)
		case types.MsgInitSession:
			return handleInitSession(ctx, vk, bk, msg)
		case types.MsgUpdateSessionBandwidth:
			return handleUpdateSessionBandwidth(ctx, vk, ak, msg)
		default:
			return types.ErrorUnknownMsgType(reflect.TypeOf(msg).Name()).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper, msg types.MsgRegisterNode) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	count, err := vk.GetNodesCount(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	id := types.NodeKey(msg.From, count)
	if details, err := vk.GetNodeDetails(ctx, id); true {
		if err != nil {
			return err.Result()
		}
		if details != nil {
			return types.ErrorNodeAlreadyExists().Result()
		}
	}

	lockAmount := csdkTypes.Coins{msg.AmountToLock}
	_, tags, err := bk.SubtractCoins(ctx, msg.From, lockAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	details := types.NodeDetails{
		ID:           id,
		Owner:        msg.From,
		LockedAmount: msg.AmountToLock,
		APIPort:      msg.APIPort,
		NetSpeed:     msg.NetSpeed,
		EncMethod:    msg.EncMethod,
		PricesPerGB:  msg.PricesPerGB,
		Version:      msg.Version,
		NodeType:     msg.NodeType,
		Status:       types.StatusRegistered,
		StatusAt:     ctx.BlockHeader().Time,
	}
	if err := vk.SetNodeDetails(ctx, id, &details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(id))

	if err := vk.SetNodesCount(ctx, msg.From, count+1); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeDetails(ctx csdkTypes.Context, vk keeper.Keeper, msg types.MsgUpdateNodeDetails) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	details, err := vk.GetNodeDetails(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if details == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if !details.Owner.Equals(msg.From) {
		return types.ErrorUnauthorized().Result()
	}
	if details.Status != types.StatusRegistered &&
		details.Status != types.StatusActive &&
		details.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	if msg.APIPort != 0 {
		details.APIPort = msg.APIPort
	}
	if !msg.NetSpeed.Download.IsZero() && !msg.NetSpeed.Upload.IsZero() {
		details.NetSpeed = msg.NetSpeed
	}
	if len(msg.EncMethod) != 0 {
		details.EncMethod = msg.EncMethod
	}
	if msg.PricesPerGB != nil && msg.PricesPerGB.Len() != 0 {
		details.PricesPerGB = msg.PricesPerGB
	}
	if len(msg.Version) != 0 {
		details.Version = msg.Version
	}
	details.DetailsAt = ctx.BlockHeader().Time

	if err := vk.SetNodeDetails(ctx, msg.ID, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(msg.ID))

	return csdkTypes.Result{Tags: allTags}
}

func handleDeregisterNode(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper, msg types.MsgDeregisterNode) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	details, err := vk.GetNodeDetails(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if details == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if !msg.From.Equals(details.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if details.Status != types.StatusRegistered &&
		details.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	details.Status = types.StatusDeregistered
	details.StatusAt = ctx.BlockHeader().Time
	if err := vk.SetNodeDetails(ctx, msg.ID, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(msg.ID))

	releaseAmount := csdkTypes.Coins{details.LockedAmount}
	_, tags, err := bk.AddCoins(ctx, msg.From, releaseAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, vk keeper.Keeper, msg types.MsgUpdateNodeStatus) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	details, err := vk.GetNodeDetails(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if details == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if !details.Owner.Equals(msg.From) {
		return types.ErrorUnauthorized().Result()
	}
	if details.Status != types.StatusRegistered &&
		details.Status != types.StatusActive &&
		details.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	details.Status = msg.Status
	details.StatusAt = ctx.BlockHeader().Time

	if err := vk.SetNodeDetails(ctx, msg.ID, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(msg.ID))

	return csdkTypes.Result{Tags: allTags}
}

func handleInitSession(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper, msg types.MsgInitSession) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	node, err := vk.GetNodeDetails(ctx, msg.NodeID)
	if err != nil {
		return err.Result()
	}
	if node == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if node.Status != types.StatusActive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	pricePerGB := node.FindPricePerGB(msg.AmountToLock.Denom)
	bandwidth, err := node.CalculateBandwidth(msg.AmountToLock)
	if err != nil {
		return err.Result()
	}

	count, err := vk.GetSessionsCount(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	id := types.SessionKey(msg.From, count)
	if details, err := vk.GetSessionDetails(ctx, id); true {
		if err != nil {
			return err.Result()
		}
		if details != nil {
			return types.ErrorSessionAlreadyExists().Result()
		}
	}

	lockAmount := csdkTypes.Coins{msg.AmountToLock}
	_, tags, err := bk.SubtractCoins(ctx, msg.From, lockAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	details := types.SessionDetails{
		ID:           id,
		NodeID:       msg.NodeID,
		Client:       msg.From,
		LockedAmount: msg.AmountToLock,
		PricePerGB:   pricePerGB,
		Bandwidth: types.SessionBandwidth{
			ToProvide: bandwidth,
		},
		Status:   types.StatusInit,
		StatusAt: ctx.BlockHeader().Time,
	}

	if err := vk.SetSessionDetails(ctx, id, &details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("session_id", []byte(id))

	if err := vk.SetSessionsCount(ctx, msg.From, count+1); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateSessionBandwidth(ctx csdkTypes.Context, vk keeper.Keeper, ak auth.AccountKeeper, msg types.MsgUpdateSessionBandwidth) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	session, err := vk.GetSessionDetails(ctx, msg.SessionID)
	if err != nil {
		return err.Result()
	}
	if session == nil {
		return types.ErrorSessionNotExists().Result()
	}
	if session.Status != types.StatusInit &&
		session.Status != types.StatusActive &&
		session.Status != types.StatusInactive {
		return types.ErrorInvalidSessionStatus().Result()
	}

	if msg.Bandwidth.LTE(session.Bandwidth.Consumed) ||
		session.Bandwidth.ToProvide.LT(msg.Bandwidth) {
		return types.ErrorInvalidBandwidth().Result()
	}

	node, err := vk.GetNodeDetails(ctx, session.NodeID)
	if err != nil {
		return err.Result()
	}
	if node == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if node.Status != types.StatusActive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	signData := types.NewBandwidthSignData(msg.SessionID, msg.Bandwidth, node.Owner, session.Client)
	signBytes, err := signData.GetBytes()
	if err != nil {
		return err.Result()
	}

	nodeOwner := ak.GetAccount(ctx, node.Owner)
	client := ak.GetAccount(ctx, session.Client)

	if !client.GetPubKey().VerifyBytes(signBytes, msg.ClientSign) ||
		!nodeOwner.GetPubKey().VerifyBytes(signBytes, msg.NodeOwnerSign) {
		return types.ErrorInvalidSign().Result()
	}

	session.Bandwidth.Consumed = msg.Bandwidth
	session.Bandwidth.ClientSign = msg.ClientSign
	session.Bandwidth.NodeOwnerSign = msg.NodeOwnerSign
	session.Bandwidth.UpdatedAt = ctx.BlockHeader().Time
	if session.Status == types.StatusInit {
		session.Status = types.StatusActive
		session.StatusAt = ctx.BlockHeader().Time
	}

	if err := vk.SetSessionDetails(ctx, msg.SessionID, session); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: allTags}
}
