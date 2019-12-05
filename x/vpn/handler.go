package vpn

import (
	"bytes"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgRegisterNode:
			return handleRegisterNode(ctx, k, msg)
		case types.MsgUpdateNodeInfo:
			return handleUpdateNodeInfo(ctx, k, msg)
		case types.MsgAddFreeClient:
			return handleAddFreeClient(ctx, k, msg)
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

func EndBlock(ctx sdk.Context, k keeper.Keeper) {
	height := ctx.BlockHeight()
	_height := height - k.SessionInactiveInterval(ctx)

	ids := k.GetActiveSessionIDs(ctx, _height)
	for _, id := range ids {
		session, _ := k.GetSession(ctx, id.(hub.SessionID))
		subscription, _ := k.GetSubscription(ctx, session.SubscriptionID)

		bandwidth := session.Bandwidth.CeilTo(hub.GB.Quo(subscription.PricePerGB.Amount))

		freeClients := k.GetFreeClientsOfNode(ctx, subscription.NodeID)

		isFreeClient := false
		for _, client := range freeClients {
			if client.Client.Equals(subscription.Client) {
				isFreeClient = true
			}
		}

		pay := sdk.Coin{}
		if !isFreeClient {
			amount := bandwidth.Sum().Mul(subscription.PricePerGB.Amount).Quo(hub.GB)
			pay = sdk.NewCoin(subscription.PricePerGB.Denom, amount)

			if !pay.IsZero() {
				node, _ := k.GetNode(ctx, subscription.NodeID)

				if err := k.SendDeposit(ctx, subscription.Client, node.Owner, pay); err != nil {
					panic(err)
				}
			}
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
}

func handleRegisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterNode) sdk.Result {
	nc := k.GetNodesCount(ctx)
	node := types.Node{
		ID:               hub.NewNodeID(nc),
		Owner:            msg.From,
		Deposit:          sdk.NewInt64Coin(k.Deposit(ctx).Denom, 0),
		Type:             msg.T,
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

		if err := k.AddDeposit(ctx, node.Owner, node.Deposit); err != nil {
			return err.Result()
		}
	}

	k.SetNode(ctx, node)
	k.SetNodeIDByAddress(ctx, node.Owner, nca, node.ID)

	k.SetNodesCount(ctx, nc+1)
	k.SetNodesCountOfAddress(ctx, node.Owner, nca+1)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleUpdateNodeInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateNodeInfo) sdk.Result {
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
		Type:          msg.T,
		Version:       msg.Version,
		Moniker:       msg.Moniker,
		PricesPerGB:   msg.PricesPerGB,
		InternetSpeed: msg.InternetSpeed,
		Encryption:    msg.Encryption,
	}
	node = node.UpdateInfo(_node)

	k.SetNode(ctx, node)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleAddFreeClient(ctx sdk.Context, k keeper.Keeper, msg types.MsgAddFreeClient) sdk.Result {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if !msg.From.Equals(node.Owner) {
		return types.ErrorUnauthorized().Result()
	}
	if node.Status == types.StatusDeRegistered {
		return types.ErrorInvalidNodeStatus().Result()
	}

	freeClient := types.NewFreeClient(msg.NodeID, msg.Client)

	k.SetFreeClient(ctx, freeClient)
	k.SetFreeClientOfNode(ctx, freeClient)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleDeregisterNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeregisterNode) sdk.Result {
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

	if node.Deposit.IsPositive() {
		if err := k.SubtractDeposit(ctx, node.Owner, node.Deposit); err != nil {
			return err.Result()
		}
	}

	node.Status = types.StatusDeRegistered
	node.StatusModifiedAt = ctx.BlockHeight()

	k.SetNode(ctx, node)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleStartSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgStartSubscription) sdk.Result {
	node, found := k.GetNode(ctx, msg.NodeID)
	if !found {
		return types.ErrorNodeDoesNotExist().Result()
	}
	if node.Status != types.StatusRegistered {
		return types.ErrorInvalidNodeStatus().Result()
	}

	if err := k.AddDeposit(ctx, msg.From, msg.Deposit); err != nil {
		return err.Result()
	}

	bandwidth, err := node.DepositToBandwidth(msg.Deposit)
	if err != nil {
		return err.Result()
	}

	pricePerGB := node.FindPricePerGB(msg.Deposit.Denom)

	sc := k.GetSubscriptionsCount(ctx)
	subscription := types.Subscription{
		ID:                 hub.NewSubscriptionID(sc),
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

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleEndSubscription(ctx sdk.Context, k keeper.Keeper, msg types.MsgEndSubscription) sdk.Result {
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

	if err := k.SubtractDeposit(ctx, subscription.Client, subscription.RemainingDeposit); err != nil {
		return err.Result()
	}

	subscription.Status = types.StatusInactive
	subscription.StatusModifiedAt = ctx.BlockHeight()

	k.SetSubscription(ctx, subscription)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleUpdateSessionInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateSessionInfo) sdk.Result {
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
	data := hub.NewBandwidthSignatureData(subscription.ID, scs, msg.Bandwidth).Bytes()
	if !msg.NodeOwnerSignature.VerifyBytes(data, msg.NodeOwnerSignature.Signature) {
		return types.ErrorInvalidBandwidthSignature().Result()
	}
	if !msg.ClientSignature.VerifyBytes(data, msg.ClientSignature.Signature) {
		return types.ErrorInvalidBandwidthSignature().Result()
	}

	if subscription.RemainingBandwidth.AnyLT(msg.Bandwidth) {
		return types.ErrorInvalidBandwidth().Result()
	}

	var session types.Session

	id, found := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, scs)
	if !found {
		sc := k.GetSessionsCount(ctx)
		session = types.Session{
			ID:             hub.NewSessionID(sc),
			SubscriptionID: subscription.ID,
			Bandwidth:      hub.NewBandwidthFromInt64(0, 0),
		}

		k.SetSessionsCount(ctx, sc+1)
		k.SetSessionIDBySubscriptionID(ctx, subscription.ID, scs, session.ID)
	} else {
		session, _ = k.GetSession(ctx, id)
	}

	k.RemoveSessionIDFromActiveList(ctx, session.StatusModifiedAt, session.ID)
	k.AddSessionIDToActiveList(ctx, ctx.BlockHeight(), session.ID)

	session.Bandwidth = msg.Bandwidth
	session.Status = types.StatusActive
	session.StatusModifiedAt = ctx.BlockHeight()

	k.SetSession(ctx, session)

	return sdk.Result{Events: ctx.EventManager().Events()}
}
