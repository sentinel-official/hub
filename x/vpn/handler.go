package vpn

import (
	"fmt"
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func NewHandler(k keeper.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case types.MsgRegisterNode:
			return handleRegisterNode(ctx, k, msg)
		case types.MsgUpdateNodeInfo:
			return handleUpdateNodeInfo(ctx, k, msg)
		case types.MsgUpdateNodeStatus:
			return handleUpdateNodeStatus(ctx, k, msg)
		case types.MsgDeregisterNode:
			return handleDeregisterNode(ctx, k, msg)
		case types.MsgStartSubscription:
			return handleStartSubscription(ctx, k, msg)
		case types.MsgEndSubscription:
			return handleEndSubscription(ctx, k, msg)
		case types.MsgUpdateSessionInfo:
			return handleUpdateSessionInfo(ctx, k, msg)
		default:
			return types.ErrorUnknownMsgType(reflect.TypeOf(msg).Name()).Result()
		}
	}
}

func endBlockNodes(ctx csdkTypes.Context, k keeper.Keeper, height int64) csdkTypes.Tags {
	allTags := csdkTypes.EmptyTags()

	_height := height - k.NodeInactiveInterval(ctx)
	ids := k.GetActiveNodeIDs(ctx, _height)

	for _, id := range ids {
		node, _ := k.GetNode(ctx, id)

		node.Status = types.StatusInactive
		node.StatusModifiedAt = height

		k.SetNode(ctx, node)
	}

	k.DeleteActiveNodeIDs(ctx, _height)
	return allTags
}

func endBlockSessions(ctx csdkTypes.Context, k keeper.Keeper, height int64) csdkTypes.Tags {
	allTags := csdkTypes.EmptyTags()

	_height := height - k.SessionInactiveInterval(ctx)
	ids := k.GetActiveSessionIDs(ctx, _height)

	for _, id := range ids {
		session, _ := k.GetSession(ctx, id)
		subscription, _ := k.GetSubscription(ctx, session.SubscriptionID)

		amount := session.CalculatedBandwidth.Sum().
			Mul(subscription.PricePerGB.Amount).Quo(sdkTypes.GB.Add(sdkTypes.GB))
		pay := csdkTypes.NewCoin(subscription.PricePerGB.Denom, amount)

		consumedDeposit := subscription.ConsumedDeposit.Add(pay)
		consumedBandwidth := subscription.ConsumedBandwidth.Add(session.Bandwidth)
		calculatedBandwidth := subscription.CalculatedBandwidth.Add(session.CalculatedBandwidth)

		if subscription.TotalDeposit.IsLT(consumedDeposit) {
			panic(fmt.Errorf("subscription total deposit is less than "+
				"consumed deposit: %s < %s", subscription.TotalDeposit, consumedDeposit))
		}
		if subscription.TotalBandwidth.AllLT(calculatedBandwidth) {
			panic(fmt.Errorf("subscription total bandwidth is less than "+
				"calculated bandwidth: %s < %s", subscription.TotalBandwidth, calculatedBandwidth))
		}

		if !pay.IsZero() {
			node, _ := k.GetNode(ctx, subscription.NodeID)

			tags, err := k.SendDeposit(ctx, subscription.Client, node.Owner, pay)
			if err != nil {
				panic(err)
			}

			allTags = allTags.AppendTags(tags)
		}

		session.Status = types.StatusInactive
		session.StatusModifiedAt = height

		subscription.ConsumedDeposit = consumedDeposit
		subscription.ConsumedBandwidth = consumedBandwidth
		subscription.CalculatedBandwidth = calculatedBandwidth

		k.SetSession(ctx, session)
		k.SetSubscription(ctx, subscription)

		scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
		k.SetSessionsCountOfSubscription(ctx, subscription.ID, scs+1)
	}

	k.DeleteActiveSessionIDs(ctx, _height)
	return allTags
}

func EndBlock(ctx csdkTypes.Context, k keeper.Keeper) csdkTypes.Tags {
	allTags := csdkTypes.EmptyTags()
	height := ctx.BlockHeight()

	tags := endBlockNodes(ctx, k, height)
	allTags = allTags.AppendTags(tags)

	tags = endBlockSessions(ctx, k, height)
	allTags = allTags.AppendTags(tags)

	return allTags
}

func handleRegisterNode(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgRegisterNode) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	nc := k.GetNodesCount(ctx)
	node := types.Node{
		ID:               sdkTypes.NewIDFromUInt64(nc),
		Owner:            msg.From,
		Deposit:          csdkTypes.NewInt64Coin(k.Deposit(ctx).Denom, 0),
		Type:             msg.Type_,
		Version:          msg.Version,
		Moniker:          msg.Moniker,
		PricesPerGB:      msg.PricesPerGB,
		InternetSpeed:    msg.InternetSpeed,
		Encryption:       msg.Encryption,
		Status:           types.StatusRegistered,
		StatusModifiedAt: ctx.BlockHeight(),
	}

	nca := k.GetNodesCountOfAddress(ctx, node.Owner)
	if nca >= k.FreeNodesCount(ctx) {
		node.Deposit = k.Deposit(ctx)

		tags, err := k.AddDeposit(ctx, node.Owner, node.Deposit)
		if err != nil {
			return err.Result()
		}

		allTags = allTags.AppendTags(tags)
	}

	k.SetNode(ctx, node)
	k.SetNodeIDByAddress(ctx, node.Owner, nca, node.ID)

	k.SetNodesCount(ctx, nc+1)
	k.SetNodesCountOfAddress(ctx, node.Owner, nca+1)

	allTags = allTags.AppendTag(types.TagNodeID, node.ID.String())
	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeInfo(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgUpdateNodeInfo) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	node, found := k.GetNode(ctx, msg.ID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status == types.StatusDeRegistered {
		return types.ErrorInvalidNodeStatus().Result()
	}

	_node := types.Node{
		Type:          msg.Type_,
		Version:       msg.Version,
		Moniker:       msg.Moniker,
		PricesPerGB:   msg.PricesPerGB,
		InternetSpeed: msg.InternetSpeed,
		Encryption:    msg.Encryption,
	}
	node = node.UpdateInfo(_node)

	k.SetNode(ctx, node)
	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNodeStatus(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgUpdateNodeStatus) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	node, found := k.GetNode(ctx, msg.ID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status == types.StatusDeRegistered {
		return types.ErrorInvalidNodeStatus().Result()
	}

	k.RemoveNodeIDFromActiveList(ctx, node.StatusModifiedAt, node.ID)
	if msg.Status == types.StatusActive {
		k.AddNodeIDToActiveList(ctx, ctx.BlockHeight(), node.ID)
	}

	node.Status = msg.Status
	node.StatusModifiedAt = ctx.BlockHeight()

	k.SetNode(ctx, node)
	return csdkTypes.Result{Tags: allTags}
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
	if node.Status == types.StatusActive || node.Status == types.StatusDeRegistered {
		return types.ErrorInvalidNodeStatus().Result()
	}

	node.Status = types.StatusDeRegistered
	node.StatusModifiedAt = ctx.BlockHeight()
	k.SetNode(ctx, node)

	if node.Deposit.IsPositive() {
		tags, err := k.SubtractDeposit(ctx, node.Owner, node.Deposit)
		if err != nil {
			return err.Result()
		}

		allTags = allTags.AppendTags(tags)
	}

	return csdkTypes.Result{Tags: allTags}
}

func handleStartSubscription(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgStartSubscription) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if node.Status != types.StatusActive {
		return types.ErrorInvalidNodeStatus().Result()
	}

	tags, err := k.AddDeposit(ctx, msg.From, msg.Deposit)
	if err != nil {
		return err.Result()
	}

	allTags = allTags.AppendTags(tags)

	bandwidth, err := node.DepositToBandwidth(msg.Deposit)
	if err != nil {
		return err.Result()
	}

	pricePerGB := node.FindPricePerGB(msg.Deposit.Denom)

	sc := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:                  sdkTypes.NewIDFromUInt64(sc),
		NodeID:              node.ID,
		Client:              msg.From,
		PricePerGB:          pricePerGB,
		TotalDeposit:        msg.Deposit,
		TotalBandwidth:      bandwidth,
		ConsumedDeposit:     csdkTypes.NewInt64Coin(msg.Deposit.Denom, 0),
		ConsumedBandwidth:   sdkTypes.NewBandwidthFromInt64(0, 0),
		CalculatedBandwidth: sdkTypes.NewBandwidthFromInt64(0, 0),
		Status:              types.StatusActive,
		StatusModifiedAt:    ctx.BlockHeight(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionsCount(ctx, sc+1)

	nsc := k.GetSubscriptionsCountOfNode(ctx, node.ID)
	k.SetSubscriptionIDByNodeID(ctx, node.ID, nsc, subscription.ID)
	k.SetSubscriptionsCountOfNode(ctx, node.ID, nsc+1)

	sca := k.GetSubscriptionsCountOfAddress(ctx, subscription.Client)
	k.SetSubscriptionIDByAddress(ctx, subscription.Client, sca, subscription.ID)
	k.SetSubscriptionsCountOfAddress(ctx, subscription.Client, sca+1)

	allTags = allTags.AppendTags(tags)
	return csdkTypes.Result{Tags: allTags}
}

func handleEndSubscription(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgEndSubscription) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if !msg.From.Equals(subscription.Client) {
		return types.ErrorUnauthorized().Result()
	}
	if subscription.Status != types.StatusActive {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)

	_, found = k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
	if found {
		return types.ErrorSessionAlreadyExists().Result()
	}

	remaining := subscription.TotalDeposit.Sub(subscription.ConsumedDeposit)

	tags, err := k.SubtractDeposit(ctx, subscription.Client, remaining)
	if err != nil {
		return err.Result()
	}

	allTags = allTags.AppendTags(tags)

	subscription.Status = types.StatusInactive
	subscription.StatusModifiedAt = ctx.BlockHeight()
	k.SetSubscription(ctx, subscription)

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateSessionInfo(ctx csdkTypes.Context, k keeper.Keeper, msg types.MsgUpdateSessionInfo) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	subscription, found := k.GetSubscription(ctx, msg.SubscriptionID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if subscription.Status == types.StatusInactive {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}

	scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)

	var session types.Session
	id, found := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
	if !found {
		sc := k.GetSessionsCount(ctx)
		session = types.Session{
			ID:                  sdkTypes.NewIDFromUInt64(sc),
			SubscriptionID:      subscription.ID,
			Bandwidth:           sdkTypes.NewBandwidthFromInt64(0, 0),
			CalculatedBandwidth: sdkTypes.NewBandwidthFromInt64(0, 0),
		}

		k.SetSessionsCount(ctx, sc+1)
		k.SetSessionIDBySubscriptionID(ctx, subscription.ID, scs, session.ID)
	} else {
		session, _ = k.GetSession(ctx, id)
	}

	if msg.Bandwidth.AllLT(session.Bandwidth) ||
		subscription.TotalBandwidth.AllLT(subscription.CalculatedBandwidth.Add(msg.Bandwidth)) {

		return types.ErrorInvalidBandwidth().Result()
	}

	node, _ := k.GetNode(ctx, subscription.NodeID)
	data := sdkTypes.NewBandwidthSignData(subscription.ID, scs, msg.Bandwidth, node.Owner, subscription.Client).Bytes()

	nodeOwnerPubKey, err := k.GetNodeOwnerPubKey(ctx, node.ID)
	if err != nil {
		return err.Result()
	}

	clientPubKey, err := k.GetSubscriptionClientPubKey(ctx, node.ID)
	if err != nil {
		return err.Result()
	}

	if !nodeOwnerPubKey.VerifyBytes(data, msg.NodeOwnerSign) ||
		!clientPubKey.VerifyBytes(data, msg.ClientSign) {

		return types.ErrorInvalidBandwidthSign().Result()
	}

	k.RemoveSessionIDFromActiveList(ctx, session.StatusModifiedAt, session.ID)
	k.AddSessionIDToActiveList(ctx, ctx.BlockHeight(), session.ID)

	session.Bandwidth = msg.Bandwidth
	session.CalculatedBandwidth = msg.Bandwidth.CeilTo(sdkTypes.GB.Quo(subscription.PricePerGB.Amount))
	session.NodeOwnerSign = msg.NodeOwnerSign
	session.ClientSign = msg.ClientSign
	session.Status = types.StatusActive
	session.StatusModifiedAt = ctx.BlockHeight()

	k.SetSession(ctx, session)
	return csdkTypes.Result{Tags: allTags}
}
