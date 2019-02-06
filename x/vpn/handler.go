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

func endBlockNodes(ctx csdkTypes.Context, vk keeper.Keeper) {
	inactiveHeight := ctx.BlockHeight() - 50
	if inactiveHeight <= 0 {
		return
	}

	nodeIDs, err := vk.GetActiveNodeIDsAtHeight(ctx, inactiveHeight)
	if err != nil {
		panic(err)
	}

	for _, nodeID := range nodeIDs {
		node, err := vk.GetNodeDetails(ctx, nodeID)
		if err != nil {
			panic(err)
		}

		node.Status = types.StatusInactive
		node.StatusAtHeight = ctx.BlockHeight()
		if err := vk.SetNodeDetails(ctx, node); err != nil {
			panic(err)
		}
	}

	if err := vk.SetActiveNodeIDsAtHeight(ctx, inactiveHeight, nil); err != nil {
		panic(err)
	}
}

func endBlockSessions(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper) {
	endHeight := ctx.BlockHeight() - 35
	if endHeight <= 0 {
		return
	}

	sessionIDs, err := vk.GetActiveSessionIDsAtHeight(ctx, endHeight)
	if err != nil {
		panic(err)
	}

	for _, sessionID := range sessionIDs {
		session, err := vk.GetSessionDetails(ctx, sessionID)
		if err != nil {
			panic(err)
		}

		session.Status = types.StatusEnd
		session.StatusAtHeight = ctx.BlockHeight()
		if err := vk.SetSessionDetails(ctx, session); err != nil {
			panic(err)
		}

		node, err := vk.GetNodeDetails(ctx, session.NodeID)
		if err != nil {
			panic(err)
		}

		payAmount := session.Amount()
		remainingAmount := session.LockedAmount.Minus(payAmount)

		if !payAmount.IsZero() {
			_, _, err := bk.AddCoins(ctx, node.Owner, csdkTypes.Coins{payAmount})
			if err != nil {
				panic(err)
			}
		}

		if !remainingAmount.IsZero() {
			_, _, err := bk.AddCoins(ctx, session.Client, csdkTypes.Coins{remainingAmount})
			if err != nil {
				panic(err)
			}
		}
	}

	if err := vk.SetActiveSessionIDsAtHeight(ctx, endHeight, nil); err != nil {
		panic(err)
	}
}

func EndBlock(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper) {
	endBlockNodes(ctx, vk)
	endBlockSessions(ctx, vk, bk)
}

func handleRegisterNode(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper,
	msg types.MsgRegisterNode) csdkTypes.Result {

	allTags := csdkTypes.EmptyTags()

	lockAmount := csdkTypes.Coins{msg.AmountToLock}
	_, tags, err := bk.SubtractCoins(ctx, msg.From, lockAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	details := types.NodeDetails{
		Owner:           msg.From,
		LockedAmount:    msg.AmountToLock,
		APIPort:         msg.APIPort,
		NetSpeed:        msg.NetSpeed,
		EncMethod:       msg.EncMethod,
		PricesPerGB:     msg.PricesPerGB,
		Version:         msg.Version,
		NodeType:        msg.NodeType,
		Status:          types.StatusRegistered,
		StatusAtHeight:  ctx.BlockHeight(),
		DetailsAtHeight: ctx.BlockHeight(),
	}

	tags, err = vk.AddNode(ctx, &details)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeDetails(ctx csdkTypes.Context, vk keeper.Keeper,
	msg types.MsgUpdateNodeDetails) csdkTypes.Result {

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
		details.Status != types.StatusActive &&
		details.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	newDetails := types.NodeDetails{
		APIPort:     msg.APIPort,
		NetSpeed:    msg.NetSpeed,
		EncMethod:   msg.EncMethod,
		PricesPerGB: msg.PricesPerGB,
		Version:     msg.Version,
	}
	details.UpdateDetails(newDetails)
	details.DetailsAtHeight = ctx.BlockHeight()

	if err := vk.SetNodeDetails(ctx, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(msg.ID))

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, vk keeper.Keeper,
	msg types.MsgUpdateNodeStatus) csdkTypes.Result {

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
		details.Status != types.StatusActive &&
		details.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	if err := vk.RemoveActiveNodeIDAtHeight(ctx, details.StatusAtHeight, details.ID); err != nil {
		return err.Result()
	}
	if msg.Status == types.StatusActive {
		if err := vk.AddActiveNodeIDAtHeight(ctx, ctx.BlockHeight(), details.ID); err != nil {
			return err.Result()
		}
	}

	details.Status = msg.Status
	details.StatusAtHeight = ctx.BlockHeight()
	if err := vk.SetNodeDetails(ctx, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(msg.ID))

	return csdkTypes.Result{Tags: allTags}
}

func handleDeregisterNode(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper,
	msg types.MsgDeregisterNode) csdkTypes.Result {

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

	if err := vk.RemoveActiveNodeIDAtHeight(ctx, details.StatusAtHeight, details.ID); err != nil {
		return err.Result()
	}

	details.Status = types.StatusDeregistered
	details.StatusAtHeight = ctx.BlockHeight()
	if err := vk.SetNodeDetails(ctx, details); err != nil {
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

func handleInitSession(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper,
	msg types.MsgInitSession) csdkTypes.Result {

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

	lockAmount := csdkTypes.Coins{msg.AmountToLock}
	_, tags, err := bk.SubtractCoins(ctx, msg.From, lockAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	details := types.SessionDetails{
		NodeID:       msg.NodeID,
		Client:       msg.From,
		LockedAmount: msg.AmountToLock,
		PricePerGB:   pricePerGB,
		Bandwidth: types.SessionBandwidth{
			ToProvide:       bandwidth,
			UpdatedAtHeight: ctx.BlockHeight(),
		},
		Status:         types.StatusInit,
		StatusAtHeight: ctx.BlockHeight(),
	}

	tags, err = vk.AddSession(ctx, &details)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateSessionBandwidth(ctx csdkTypes.Context, vk keeper.Keeper, ak auth.AccountKeeper,
	msg types.MsgUpdateSessionBandwidth) csdkTypes.Result {

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

	if err := vk.RemoveActiveSessionIDsAtHeight(ctx, session.StatusAtHeight, session.ID); err != nil {
		return err.Result()
	}
	if err := vk.AddActiveSessionIDsAtHeight(ctx, ctx.BlockHeight(), session.ID); err != nil {
		return err.Result()
	}

	sign := types.NewBandwidthSign(msg.SessionID, msg.Bandwidth, node.Owner, session.Client)
	if err := keeper.VerifyAndUpdateSessionBandwidth(ctx, ak, session, sign,
		msg.ClientSign, msg.NodeOwnerSign); err != nil {
		return err.Result()
	}

	if err := vk.SetSessionDetails(ctx, session); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: allTags}
}
