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

	k.SetActiveNodeIDs(ctx, _height, nil)
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

		k.SetSession(ctx, session)

		subscription.ConsumedDeposit = consumedDeposit
		subscription.ConsumedBandwidth = consumedBandwidth
		subscription.CalculatedBandwidth = calculatedBandwidth
		subscription.SessionsCount++

		if subscription.ConsumedDeposit.IsEqual(subscription.TotalDeposit) {
			subscription.Status = types.StatusInactive
			subscription.StatusModifiedAt = height
		}

		k.SetSubscription(ctx, subscription)
	}

	k.SetActiveSessionIDs(ctx, _height, nil)
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
	node := types.Node{
		ID:                 sdkTypes.NewIDFromUInt64(k.GetNodesCount(ctx)),
		Owner:              msg.From,
		Deposit:            csdkTypes.NewInt64Coin(k.Deposit(ctx).Denom, 0),
		Type:               msg.Type_,
		Version:            msg.Version,
		Moniker:            msg.Moniker,
		PricesPerGB:        msg.PricesPerGB,
		InternetSpeed:      msg.InternetSpeed,
		Encryption:         msg.Encryption,
		SubscriptionsCount: 0,
		Status:             types.StatusRegistered,
		StatusModifiedAt:   ctx.BlockHeight(),
	}

	allTags, err := k.AddNode(ctx, node)
	if err != nil {
		return err.Result()
	}

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

	k.RemoveActiveNodeID(ctx, node.StatusModifiedAt, node.ID)
	if msg.Status == types.StatusActive {
		k.AddActiveNodeID(ctx, ctx.BlockHeight(), node.ID)
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

	pricePerGB := node.FindPricePerGB(msg.Deposit.Denom)
	bandwidth, err := node.DepositToBandwidth(msg.Deposit)
	if err != nil {
		return err.Result()
	}

	subscription := types.Subscription{
		ID:                  sdkTypes.NewIDFromUInt64(k.GetSubscriptionsCount(ctx)),
		NodeID:              node.ID,
		Client:              msg.From,
		PricePerGB:          pricePerGB,
		TotalDeposit:        msg.Deposit,
		TotalBandwidth:      bandwidth,
		ConsumedDeposit:     csdkTypes.NewInt64Coin(msg.Deposit.Denom, 0),
		ConsumedBandwidth:   sdkTypes.NewBandwidthFromInt64(0, 0),
		CalculatedBandwidth: sdkTypes.NewBandwidthFromInt64(0, 0),
		SessionsCount:       0,
		Status:              types.StatusActive,
		StatusModifiedAt:    ctx.BlockHeight(),
	}

	tags, err := k.AddSubscription(ctx, node, subscription)
	if err != nil {
		return err.Result()
	}

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

	_, found = k.GetSessionIDBySubscriptionID(ctx, subscription.ID, subscription.SessionsCount)
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

	var session types.Session

	id, found := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, subscription.SessionsCount)
	if !found {
		session = types.Session{
			ID:                  sdkTypes.NewIDFromUInt64(k.GetSessionsCount(ctx)),
			SubscriptionID:      subscription.ID,
			Bandwidth:           sdkTypes.NewBandwidthFromInt64(0, 0),
			CalculatedBandwidth: sdkTypes.NewBandwidthFromInt64(0, 0),
		}

		k.SetSessionsCount(ctx, session.ID.UInt64()+1)
		k.SetSessionIDBySubscriptionID(ctx, subscription.ID, subscription.SessionsCount, session.ID)
	} else {
		session, _ = k.GetSession(ctx, id)
	}

	if msg.Bandwidth.AllLT(session.Bandwidth) ||
		subscription.TotalBandwidth.AllLT(subscription.CalculatedBandwidth.Add(msg.Bandwidth)) {

		return types.ErrorInvalidBandwidth().Result()
	}

	node, _ := k.GetNode(ctx, subscription.NodeID)
	data := sdkTypes.NewBandwidthSignData(subscription.ID, subscription.SessionsCount, msg.Bandwidth,
		node.Owner, subscription.Client).Bytes()

	if !node.OwnerPubKey.VerifyBytes(data, msg.NodeOwnerSign) ||
		!subscription.ClientPubKey.VerifyBytes(data, msg.ClientSign) {

		return types.ErrorInvalidBandwidthSign().Result()
	}

	k.RemoveActiveSessionID(ctx, session.StatusModifiedAt, session.ID)
	k.AddActiveSessionID(ctx, ctx.BlockHeight(), session.ID)

	session.Bandwidth = msg.Bandwidth
	session.CalculatedBandwidth = msg.Bandwidth.CeilTo(sdkTypes.GB.Quo(subscription.PricePerGB.Amount))
	session.NodeOwnerSign = msg.NodeOwnerSign
	session.ClientSign = msg.ClientSign
	session.Status = types.StatusActive
	session.StatusModifiedAt = ctx.BlockHeight()

	k.SetSession(ctx, session)
	return csdkTypes.Result{Tags: allTags}
}
