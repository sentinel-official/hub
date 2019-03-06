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
			return handleRegisterNode(ctx, vk, ak, bk, msg)
		case types.MsgUpdateNodeDetails:
			return handleUpdateNodeDetails(ctx, vk, msg)
		case types.MsgUpdateNodeStatus:
			return handleUpdateNodeStatus(ctx, vk, msg)
		case types.MsgDeregisterNode:
			return handleDeregisterNode(ctx, vk, bk, msg)
		case types.MsgInitSession:
			return handleInitSession(ctx, vk, ak, bk, msg)
		case types.MsgUpdateSessionBandwidth:
			return handleUpdateSessionBandwidth(ctx, vk, msg)
		default:
			return types.ErrorUnknownMsgType(reflect.TypeOf(msg).Name()).Result()
		}
	}
}

func endBlockNodes(ctx csdkTypes.Context, vk keeper.Keeper, height int64) {
	inactiveHeight := height - 50
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
		node.StatusAtHeight = height
		if err := vk.SetNodeDetails(ctx, node); err != nil {
			panic(err)
		}
	}

	if err := vk.SetActiveNodeIDsAtHeight(ctx, inactiveHeight, nil); err != nil {
		panic(err)
	}
}

func endBlockSessions(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper, height int64) {
	inactiveHeight := height - 25
	sessionIDs, err := vk.GetActiveSessionIDsAtHeight(ctx, inactiveHeight)
	if err != nil {
		panic(err)
	}

	for _, sessionID := range sessionIDs {
		session, err := vk.GetSessionDetails(ctx, sessionID)
		if err != nil {
			panic(err)
		}

		session.Status = types.StatusEnd
		session.StatusAtHeight = height
		if err := vk.SetSessionDetails(ctx, session); err != nil {
			panic(err)
		}

		payAmount := session.Amount()
		remainingAmount := session.LockedAmount.Minus(payAmount)

		if !payAmount.IsZero() {
			_, _, err := bk.AddCoins(ctx, session.NodeOwner, csdkTypes.Coins{payAmount})
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

	if err := vk.SetActiveSessionIDsAtHeight(ctx, inactiveHeight, nil); err != nil {
		panic(err)
	}
}

func EndBlock(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper) {
	height := ctx.BlockHeight()
	endBlockNodes(ctx, vk, height)
	endBlockSessions(ctx, vk, bk, height)
}

func handleRegisterNode(ctx csdkTypes.Context, vk keeper.Keeper, ak auth.AccountKeeper, bk bank.Keeper,
	msg types.MsgRegisterNode) csdkTypes.Result {

	allTags := csdkTypes.EmptyTags()

	lockAmount := csdkTypes.Coins{msg.AmountToLock}
	_, tags, err := bk.SubtractCoins(ctx, msg.From, lockAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	nodeOwnerPubKey, err := ak.GetPubKey(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	height := ctx.BlockHeight()
	details := types.NodeDetails{
		Owner:           msg.From,
		PubKey:          nodeOwnerPubKey,
		LockedAmount:    msg.AmountToLock,
		APIPort:         msg.APIPort,
		NetSpeed:        msg.NetSpeed,
		Encryption:      msg.Encryption,
		PricesPerGB:     msg.PricesPerGB,
		Version:         msg.Version,
		NodeType:        msg.NodeType,
		Status:          types.StatusRegistered,
		StatusAtHeight:  height,
		DetailsAtHeight: height,
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
		Encryption:  msg.Encryption,
		PricesPerGB: msg.PricesPerGB,
		Version:     msg.Version,
	}
	details.UpdateDetails(newDetails)
	details.DetailsAtHeight = ctx.BlockHeight()

	if err := vk.SetNodeDetails(ctx, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", msg.ID.String())

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

	height := ctx.BlockHeight()
	if msg.Status == types.StatusActive {
		if err := vk.AddActiveNodeIDAtHeight(ctx, height, details.ID); err != nil {
			return err.Result()
		}
	}

	details.Status = msg.Status
	details.StatusAtHeight = height
	if err := vk.SetNodeDetails(ctx, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", msg.ID.String())

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
	allTags = allTags.AppendTag("node_id", msg.ID.String())

	releaseAmount := csdkTypes.Coins{details.LockedAmount}
	_, tags, err := bk.AddCoins(ctx, msg.From, releaseAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleInitSession(ctx csdkTypes.Context, vk keeper.Keeper, ak auth.AccountKeeper, bk bank.Keeper,
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

	clientPubKey, err := ak.GetPubKey(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	height := ctx.BlockHeight()
	details := types.SessionDetails{
		NodeID:          msg.NodeID,
		NodeOwner:       node.Owner,
		NodeOwnerPubKey: node.PubKey,
		Client:          msg.From,
		ClientPubKey:    clientPubKey,
		LockedAmount:    msg.AmountToLock,
		PricePerGB:      pricePerGB,
		Bandwidth: types.SessionBandwidth{
			ToProvide:       bandwidth,
			UpdatedAtHeight: height,
		},
		Status:         types.StatusInit,
		StatusAtHeight: height,
	}

	tags, err = vk.AddSession(ctx, &details)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateSessionBandwidth(ctx csdkTypes.Context, vk keeper.Keeper,
	msg types.MsgUpdateSessionBandwidth) csdkTypes.Result {

	allTags := csdkTypes.EmptyTags()

	session, err := vk.GetSessionDetails(ctx, msg.ID)
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

	if err := vk.RemoveActiveSessionIDsAtHeight(ctx, session.StatusAtHeight, session.ID); err != nil {
		return err.Result()
	}

	height := ctx.BlockHeight()
	if err := vk.AddActiveSessionIDsAtHeight(ctx, height, session.ID); err != nil {
		return err.Result()
	}

	if err := session.SetNewSessionBandwidth(msg.Bandwidth, msg.ClientSign, msg.NodeOwnerSign, height); err != nil {
		return types.ErrorBandwidthUpdate(err.Error()).Result()
	}
	if session.Status == StatusInit {
		session.StartedAtHeight = height
	}
	if session.Status != StatusActive {
		session.Status = StatusActive
		session.StatusAtHeight = height
	}

	if err := vk.SetSessionDetails(ctx, session); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: allTags}
}
