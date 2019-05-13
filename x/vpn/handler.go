package vpn

import (
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func NewHandler(k keeper.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case types.MsgRegisterNode:
			return handleRegisterNode(ctx, k, msg)
		case types.MsgUpdateNodeDetails:
			return handleUpdateNodeDetails(ctx, k, msg)
		case types.MsgUpdateNodeStatus:
			return handleUpdateNodeStatus(ctx, k, msg)
		case types.MsgDeregisterNode:
			return handleDeregisterNode(ctx, k, msg)
		case types.MsgInitSession:
			return handleInitSession(ctx, k, msg)
		case types.MsgUpdateSessionBandwidthInfo:
			return handleUpdateSessionBandwidthInfo(ctx, k, msg)
		default:
			return types.ErrorUnknownMsgType(reflect.TypeOf(msg).Name()).Result()
		}
	}
}

func endBlockNodes(ctx csdkTypes.Context, k keeper.Keeper, height int64) csdkTypes.Tags {
	tags := csdkTypes.EmptyTags()

	inactiveHeight := height - k.NodeInactiveInterval(ctx)
	nodeIDs := k.GetActiveNodeIDsAtHeight(ctx, inactiveHeight)
	for _, nodeID := range nodeIDs {
		node, _ := k.GetNode(ctx, nodeID)
		node.Status = types.StatusInactive
		node.StatusModifiedAtHeight = height

		k.SetNode(ctx, node)
		tags = tags.AppendTag("node_id", node.ID.String())
	}

	k.SetActiveNodeIDsAtHeight(ctx, inactiveHeight, nil)
	return tags
}

func endBlockSessions(ctx csdkTypes.Context, k keeper.Keeper, height int64) (csdkTypes.Tags, csdkTypes.Error) {
	allTags := csdkTypes.EmptyTags()

	inactiveHeight := height - k.SessionEndInterval(ctx)
	sessionIDs := k.GetActiveSessionIDsAtHeight(ctx, inactiveHeight)
	for _, sessionID := range sessionIDs {
		session, _ := k.GetSession(ctx, sessionID)
		session.Status = types.StatusEnd
		session.StatusModifiedAtHeight = height

		k.SetSession(ctx, session)
		allTags = allTags.AppendTag("session_id", session.ID.String())

		pay := session.Amount()
		remaining := session.Deposit.Sub(pay)

		if !pay.IsZero() {
			tags, err := k.AddAndSubtractDeposit(ctx, session.NodeOwner, pay)
			if err != nil {
				return nil, err
			}

			allTags = allTags.AppendTags(tags)
		}

		if !remaining.IsZero() {
			tags, err := k.AddAndSubtractDeposit(ctx, session.Client, remaining)
			if err != nil {
				return nil, err
			}

			allTags = allTags.AppendTags(tags)
		}
	}

	k.SetActiveSessionIDsAtHeight(ctx, inactiveHeight, nil)
	return allTags, nil
}

func EndBlock(ctx csdkTypes.Context, k keeper.Keeper) (csdkTypes.Tags, csdkTypes.Error) {
	allTags := csdkTypes.EmptyTags()
	height := ctx.BlockHeight()

	tags := endBlockNodes(ctx, k, height)
	allTags = allTags.AppendTags(tags)

	tags, err := endBlockSessions(ctx, k, height)
	if err != nil {
		return nil, err
	}

	allTags = allTags.AppendTags(tags)
	return allTags, nil
}

func handleRegisterNode(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgRegisterNode) csdkTypes.Result {
	node := types.Node{
		Owner:            msg.From,
		Moniker:          msg.Moniker,
		PricesPerGB:      msg.PricesPerGB,
		InternetSpeed:    msg.InternetSpeed,
		EncryptionMethod: msg.EncryptionMethod,
		Type:             msg.Type_,
		Version:          msg.Version,

		Status:                 types.StatusRegister,
		StatusModifiedAtHeight: ctx.BlockHeight(),
	}

	tags, err := k.AddNode(ctx, node)
	if err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: tags}
}

func handleUpdateNodeDetails(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgUpdateNodeDetails) csdkTypes.Result {
	tags := csdkTypes.EmptyTags()

	node, found := k.GetNode(ctx, msg.ID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status == types.StatusDeregister {
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

	k.SetNode(ctx, node)
	tags = tags.AppendTag("node_id", msg.ID.String())

	return csdkTypes.Result{Tags: tags}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgUpdateNodeStatus) csdkTypes.Result {
	tags := csdkTypes.EmptyTags()

	node, found := k.GetNode(ctx, msg.ID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status == types.StatusDeregister {
		return types.ErrorInvalidNodeStatus().Result()
	}

	k.RemoveActiveNodeIDAtHeight(ctx, node.StatusModifiedAtHeight, node.ID)

	height := ctx.BlockHeight()
	if msg.Status == types.StatusActive {
		k.AddActiveNodeIDAtHeight(ctx, height, node.ID)
	}

	node.Status = msg.Status
	node.StatusModifiedAtHeight = height

	k.SetNode(ctx, node)
	tags = tags.AppendTag("node_id", msg.ID.String())

	return csdkTypes.Result{Tags: tags}
}

func handleDeregisterNode(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgDeregisterNode) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	node, found := k.GetNode(ctx, msg.ID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status == types.StatusActive ||
		node.Status == types.StatusDeregister {

		return types.ErrorInvalidNodeStatus().Result()
	}

	node.Status = types.StatusDeregister
	node.StatusModifiedAtHeight = ctx.BlockHeight()

	k.SetNode(ctx, node)
	allTags = allTags.AppendTag("node_id", msg.ID.String())

	tags, err := k.AddAndSubtractDeposit(ctx, node.Owner, node.Deposit)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}

func handleInitSession(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgInitSession) csdkTypes.Result {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if node.Status != types.StatusActive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	pricePerGB := node.FindPricePerGB(msg.Deposit.Denom)

	toProvide, err := node.AmountToBandwidth(msg.Deposit)
	if err != nil {
		return err.Result()
	}

	height := ctx.BlockHeight()
	session := types.Session{
		NodeID:          node.ID,
		NodeOwner:       node.Owner,
		NodeOwnerPubKey: node.OwnerPubKey,
		Client:          msg.From,
		Deposit:         msg.Deposit,
		PricePerGB:      pricePerGB,
		BandwidthInfo: types.SessionBandwidthInfo{
			ToProvide:        toProvide,
			ModifiedAtHeight: height,
		},
		Status:                 types.StatusInit,
		StatusModifiedAtHeight: height,
	}

	tags, err := k.AddSession(ctx, session)
	if err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: tags}
}

func handleUpdateSessionBandwidthInfo(ctx csdkTypes.Context, k keeper.Keeper,
	msg types.MsgUpdateSessionBandwidthInfo) csdkTypes.Result {

	tags := csdkTypes.EmptyTags()

	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return types.ErrorSessionDoesNotExist().Result()
	}
	if session.Status == types.StatusEnd {
		return types.ErrorInvalidSessionStatus().Result()
	}

	k.RemoveActiveSessionIDAtHeight(ctx, session.StatusModifiedAtHeight, session.ID)

	height := ctx.BlockHeight()
	k.AddActiveSessionIDAtHeight(ctx, height, session.ID)

	if err := session.UpdateSessionBandwidthInfo(msg.Consumed,
		msg.NodeOwnerSign, msg.ClientSign, height); err != nil {

		return types.ErrorBandwidthUpdate(err.Error()).Result()
	}

	session.Status = StatusActive
	session.StatusModifiedAtHeight = height

	k.SetSession(ctx, session)
	return csdkTypes.Result{Tags: tags}
}
