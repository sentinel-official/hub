package vpn

import (
	"bytes"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/sentinel-hub/types"
	"github.com/sentinel-official/sentinel-hub/x/vpn/keeper"
	"github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
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

func endBlockNodes(ctx sdk.Context, k keeper.Keeper, height int64) sdk.Tags {
	allTags := sdk.EmptyTags()

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

func endBlockSessions(ctx sdk.Context, k keeper.Keeper, height int64) sdk.Tags {
	allTags := sdk.EmptyTags()

	_height := height - k.SessionInactiveInterval(ctx)
	ids := k.GetActiveSessionIDs(ctx, _height)

	for _, id := range ids {
		session, _ := k.GetSession(ctx, id)
		subscription, _ := k.GetSubscription(ctx, session.SubscriptionID)

		bandwidth := session.Bandwidth.CeilTo(hub.GB.Quo(subscription.PricePerGB.Amount))
		amount := bandwidth.Sum().Mul(subscription.PricePerGB.Amount).Quo(hub.GB)
		pay := sdk.NewCoin(subscription.PricePerGB.Denom, amount)

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

		subscription.RemainingDeposit = subscription.RemainingDeposit.Sub(pay)
		subscription.RemainingBandwidth = subscription.RemainingBandwidth.Sub(bandwidth)
		k.SetSubscription(ctx, subscription)

		scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
		k.SetSessionsCountOfSubscription(ctx, subscription.ID, scs+1)
	}

	k.DeleteActiveSessionIDs(ctx, _height)
	return allTags
}

func EndBlock(ctx sdk.Context, k keeper.Keeper) sdk.Tags {
	allTags := sdk.EmptyTags()
	height := ctx.BlockHeight()

	tags := endBlockNodes(ctx, k, height)
	allTags = allTags.AppendTags(tags)

	tags = endBlockSessions(ctx, k, height)
	allTags = allTags.AppendTags(tags)

	return allTags
}

func handleRegisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterNode) sdk.Result {
	allTags := sdk.EmptyTags()

	nc := k.GetNodesCount(ctx)
	node := types.Node{
		ID:               hub.NewIDFromUInt64(nc),
		Owner:            msg.From,
		Deposit:          sdk.NewInt64Coin(k.Deposit(ctx).Denom, 0),
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
	return sdk.Result{Tags: allTags}
}

func handleUpdateNodeInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateNodeInfo) sdk.Result {
	allTags := sdk.EmptyTags()

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
	return sdk.Result{Tags: allTags}
}

func handleUpdateNodeStatus(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateNodeStatus) sdk.Result {
	allTags := sdk.EmptyTags()

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
	return sdk.Result{Tags: allTags}
}

func handleDeregisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeregisterNode) sdk.Result {
	allTags := sdk.EmptyTags()

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

	if node.Deposit.IsPositive() {
		tags, err := k.SubtractDeposit(ctx, node.Owner, node.Deposit)
		if err != nil {
			return err.Result()
		}

		allTags = allTags.AppendTags(tags)
	}

	node.Status = types.StatusDeRegistered
	node.StatusModifiedAt = ctx.BlockHeight()

	k.SetNode(ctx, node)
	return sdk.Result{Tags: allTags}
}

func handleStartSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartSubscription) sdk.Result {
	allTags := sdk.EmptyTags()

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
		ID:                 hub.NewIDFromUInt64(sc),
		NodeID:             node.ID,
		Client:             msg.From,
		PricePerGB:         pricePerGB,
		TotalDeposit:       msg.Deposit,
		RemainingDeposit:   msg.Deposit,
		RemainingBandwidth: bandwidth,
		Status:             types.StatusActive,
		StatusModifiedAt:   ctx.BlockHeight(),
	}

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionsCount(ctx, sc+1)

	nsc := k.GetSubscriptionsCountOfNode(ctx, node.ID)
	k.SetSubscriptionIDByNodeID(ctx, node.ID, nsc, subscription.ID)
	k.SetSubscriptionsCountOfNode(ctx, node.ID, nsc+1)

	sca := k.GetSubscriptionsCountOfAddress(ctx, subscription.Client)
	k.SetSubscriptionIDByAddress(ctx, subscription.Client, sca, subscription.ID)
	k.SetSubscriptionsCountOfAddress(ctx, subscription.Client, sca+1)

	allTags = allTags.AppendTag(types.TagSubscriptionID, subscription.ID.String())
	return sdk.Result{Tags: allTags}
}

func handleEndSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgEndSubscription) sdk.Result {
	allTags := sdk.EmptyTags()

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

	tags, err := k.SubtractDeposit(ctx, subscription.Client, subscription.RemainingDeposit)
	if err != nil {
		return err.Result()
	}

	allTags = allTags.AppendTags(tags)

	subscription.Status = types.StatusInactive
	subscription.StatusModifiedAt = ctx.BlockHeight()
	k.SetSubscription(ctx, subscription)

	return sdk.Result{Tags: allTags}
}

func handleUpdateSessionInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateSessionInfo) sdk.Result {
	allTags := sdk.EmptyTags()

	subscription, found := k.GetSubscription(ctx, msg.SubscriptionID)
	if !found {
		return types.ErrorSubscriptionDoesNotExist().Result()
	}
	if subscription.Status == types.StatusInactive {
		return types.ErrorInvalidSubscriptionStatus().Result()
	}
	if !bytes.Equal(msg.ClientSignature.PubKey.Address(), subscription.Client.Bytes()) {
		return types.ErrorUnauthorized().Result()
	}

	node, _ := k.GetNode(ctx, subscription.NodeID)
	if !bytes.Equal(msg.NodeOwnerSignature.PubKey.Address(), node.Owner.Bytes()) {
		return types.ErrorUnauthorized().Result()
	}

	scs := k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	data := types.NewBandwidthSignatureData(subscription.ID, scs, msg.Bandwidth).Bytes()
	if !msg.NodeOwnerSignature.VerifyBytes(data, msg.NodeOwnerSignature.Signature) {
		return types.ErrorInvalidBandwidthSignature().Result()
	}
	if !msg.ClientSignature.VerifyBytes(data, msg.ClientSignature.Signature) {
		return types.ErrorInvalidBandwidthSignature().Result()
	}

	var session types.Session
	id, found := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
	if !found {
		sc := k.GetSessionsCount(ctx)
		session = types.Session{
			ID:             hub.NewIDFromUInt64(sc),
			SubscriptionID: subscription.ID,
			Bandwidth:      hub.NewBandwidthFromInt64(0, 0),
		}

		k.SetSessionsCount(ctx, sc+1)
		k.SetSessionIDBySubscriptionID(ctx, subscription.ID, scs, session.ID)
	} else {
		session, _ = k.GetSession(ctx, id)
	}

	if subscription.RemainingBandwidth.AnyLT(msg.Bandwidth) {
		return types.ErrorInvalidBandwidth().Result()
	}

	k.RemoveSessionIDFromActiveList(ctx, session.StatusModifiedAt, session.ID)
	k.AddSessionIDToActiveList(ctx, ctx.BlockHeight(), session.ID)

	session.Bandwidth = msg.Bandwidth
	session.Status = types.StatusActive
	session.StatusModifiedAt = ctx.BlockHeight()

	k.SetSession(ctx, session)
	return sdk.Result{Tags: allTags}
}
