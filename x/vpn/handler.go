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
		case types.MsgUpdateSessionBandwidthInfo:
			return handleUpdateSessionBandwidthInfo(ctx, vk, msg)
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
		node, err := vk.GetNode(ctx, nodeID)
		if err != nil {
			panic(err)
		}

		node.Status = types.StatusInactive
		node.StatusModifiedAtHeight = height
		if err := vk.SetNode(ctx, node); err != nil {
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
		session, err := vk.GetSession(ctx, sessionID)
		if err != nil {
			panic(err)
		}

		session.Status = types.StatusEnd
		session.StatusModifiedAtHeight = height
		if err := vk.SetSession(ctx, session); err != nil {
			panic(err)
		}

		payAmount := session.Amount()
		remainingAmount := session.DepositAmount.Sub(payAmount)

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

	depositAmount := csdkTypes.Coins{msg.DepositAmount}
	_, tags, err := bk.SubtractCoins(ctx, msg.From, depositAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	ownerPubKey, err := ak.GetPubKey(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	height := ctx.BlockHeight()
	node := types.Node{
		Owner:                  msg.From,
		OwnerPubKey:            ownerPubKey,
		DepositAmount:          msg.DepositAmount,
		Moniker:                msg.Moniker,
		PricesPerGB:            msg.PricesPerGB,
		InternetSpeed:          msg.InternetSpeed,
		EncryptionMethod:       msg.EncryptionMethod,
		Type:                   msg.Type_,
		Version:                msg.Version,
		ModifiedAtHeight:       height,
		Status:                 types.StatusRegistered,
		StatusModifiedAtHeight: height,
	}

	tags, err = vk.AddNode(ctx, &node)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeDetails(ctx csdkTypes.Context, vk keeper.Keeper,
	msg types.MsgUpdateNodeDetails) csdkTypes.Result {

	allTags := csdkTypes.EmptyTags()

	node, err := vk.GetNode(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if node == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status != types.StatusRegistered &&
		node.Status != types.StatusActive &&
		node.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	_node := types.Node{
		Moniker:          msg.Moniker,
		PricesPerGB:      msg.PricesPerGB,
		InternetSpeed:    msg.InternetSpeed,
		EncryptionMethod: msg.EncryptionMethod,
		Type:             msg.Type_,
		Version:          msg.Version,
	}
	node.UpdateDetails(_node)
	node.ModifiedAtHeight = ctx.BlockHeight()

	if err := vk.SetNode(ctx, node); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", msg.ID.String())

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, vk keeper.Keeper,
	msg types.MsgUpdateNodeStatus) csdkTypes.Result {

	allTags := csdkTypes.EmptyTags()

	node, err := vk.GetNode(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if node == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status != types.StatusRegistered &&
		node.Status != types.StatusActive &&
		node.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	if err := vk.RemoveActiveNodeIDAtHeight(ctx, node.StatusModifiedAtHeight, node.ID); err != nil {
		return err.Result()
	}

	height := ctx.BlockHeight()
	if msg.Status == types.StatusActive {
		if err := vk.AddActiveNodeIDAtHeight(ctx, height, node.ID); err != nil {
			return err.Result()
		}
	}

	node.Status = msg.Status
	node.StatusModifiedAtHeight = height
	if err := vk.SetNode(ctx, node); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", msg.ID.String())

	return csdkTypes.Result{Tags: allTags}
}

func handleDeregisterNode(ctx csdkTypes.Context, vk keeper.Keeper, bk bank.Keeper,
	msg types.MsgDeregisterNode) csdkTypes.Result {

	allTags := csdkTypes.EmptyTags()

	node, err := vk.GetNode(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if node == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status != types.StatusRegistered &&
		node.Status != types.StatusInactive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	if err := vk.RemoveActiveNodeIDAtHeight(ctx, node.StatusModifiedAtHeight, node.ID); err != nil {
		return err.Result()
	}

	node.Status = types.StatusDeregistered
	node.StatusModifiedAtHeight = ctx.BlockHeight()
	if err := vk.SetNode(ctx, node); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", msg.ID.String())

	releaseAmount := csdkTypes.Coins{node.DepositAmount}
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

	node, err := vk.GetNode(ctx, msg.NodeID)
	if err != nil {
		return err.Result()
	}
	if node == nil {
		return types.ErrorNodeNotExists().Result()
	}
	if node.Status != types.StatusActive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	pricePerGB := node.FindPricePerGB(msg.DepositAmount.Denom)
	toProvide, err := node.AmountToBandwidth(msg.DepositAmount)
	if err != nil {
		return err.Result()
	}

	lockAmount := csdkTypes.Coins{msg.DepositAmount}
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
	session := types.Session{
		NodeID:          msg.NodeID,
		NodeOwner:       node.Owner,
		NodeOwnerPubKey: node.OwnerPubKey,
		Client:          msg.From,
		ClientPubKey:    clientPubKey,
		DepositAmount:   msg.DepositAmount,
		PricePerGB:      pricePerGB,
		BandwidthInfo: types.SessionBandwidthInfo{
			ToProvide:        toProvide,
			ModifiedAtHeight: height,
		},
		Status:                 types.StatusInit,
		StatusModifiedAtHeight: height,
	}

	tags, err = vk.AddSession(ctx, &session)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateSessionBandwidthInfo(ctx csdkTypes.Context, vk keeper.Keeper,
	msg types.MsgUpdateSessionBandwidthInfo) csdkTypes.Result {

	allTags := csdkTypes.EmptyTags()

	session, err := vk.GetSession(ctx, msg.ID)
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

	if err := vk.RemoveActiveSessionIDsAtHeight(ctx, session.StatusModifiedAtHeight, session.ID); err != nil {
		return err.Result()
	}

	height := ctx.BlockHeight()
	if err := vk.AddActiveSessionIDsAtHeight(ctx, height, session.ID); err != nil {
		return err.Result()
	}

	if err := session.UpdateSessionBandwidthInfo(msg.Consumed, msg.NodeOwnerSign, msg.ClientSign, height); err != nil {
		return types.ErrorBandwidthUpdate(err.Error()).Result()
	}
	if session.Status == StatusInit {
		session.StartedAtHeight = height
	}
	if session.Status != StatusActive {
		session.Status = StatusActive
		session.StatusModifiedAtHeight = height
	}

	if err := vk.SetSession(ctx, session); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: allTags}
}
