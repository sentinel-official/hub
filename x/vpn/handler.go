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
		case types.MsgRemoveFreeClient:
			return handleRemoveFreeClient(ctx, k, msg)
		case types.MsgDeregisterNode:
			return handleDeregisterNode(ctx, k, msg)
		case types.MsgStartSubscription:
			return handleStartSubscription(ctx, k, msg)
		case types.MsgEndSubscription:
			return handleEndSubscription(ctx, k, msg)
		case types.MsgUpdateSessionInfo:
			return handleUpdateSessionInfo(ctx, k, msg)
		case types.MsgRegisterResolver:
			return handleRegisterResolver(ctx, k, msg)
		case types.MsgUpdateResolverInfo:
			return handleUpdateResolverInfo(ctx, k, msg)
		case types.MsgDeregisterResolver:
			return handleDeregisterResolver(ctx, k, msg)

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

		pay := sdk.NewInt(0)
		if !types.IsFreeClient(freeClients, subscription.Client) {
			amount := bandwidth.Sum().Mul(subscription.PricePerGB.Amount).Quo(hub.GB)
			payCoin := sdk.NewCoin(subscription.PricePerGB.Denom, amount)

			pay = payCoin.Amount
			if !pay.IsZero() {
				node, _ := k.GetNode(ctx, subscription.NodeID)

				if err := k.SendDeposit(ctx, subscription.Client, node.Owner, payCoin); err != nil {
					panic(err)
				}
			}
		}

		session.Status = types.StatusInactive
		session.StatusModifiedAt = height
		k.SetSession(ctx, session)

		subscription.RemainingDeposit.Amount = subscription.RemainingDeposit.Amount.Sub(pay)
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

	k.SetFreeNodesOfClient(ctx, freeClient)
	k.SetFreeClientOfNode(ctx, freeClient)

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleRemoveFreeClient(ctx sdk.Context, k keeper.Keeper, msg types.MsgRemoveFreeClient) sdk.Result {
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

	client := k.GetFreeClientOfNode(ctx, msg.NodeID, msg.Client)
	if client == nil {
		return types.ErrorFreeClientDoesNotExist().Result()
	}

	k.RemoveFreeClient(ctx, msg.NodeID, msg.Client)

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

	freeClients := k.GetFreeClientsOfNode(ctx, msg.NodeID)

	if !types.IsFreeClient(freeClients, msg.From) {
		if err := k.AddDeposit(ctx, msg.From, msg.Deposit); err != nil {
			return err.Result()
		}
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

	freeClients := k.GetFreeClientsOfNode(ctx, subscription.NodeID)

	if !types.IsFreeClient(freeClients, msg.From) {
		if err := k.SubtractDeposit(ctx, subscription.Client, subscription.RemainingDeposit); err != nil {
			return err.Result()
		}
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

func handleRegisterResolver(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterResolver) sdk.Result {
	_, found := k.GetResolver(ctx, msg.From)
	if found {
		return types.ErrorResolverAlreadyExist().Result()
	}

	resolver := types.Resolver{
		Owner:      msg.From,
		Commission: msg.Commission,
		Status:     types.StatusRegistered,
	}

	k.SetResolver(ctx, resolver)

	return sdk.Result{}
}

func handleUpdateResolverInfo(ctx sdk.Context, k keeper.Keeper, msg types.MsgUpdateResolverInfo) sdk.Result {
	resolver, found := k.GetResolver(ctx, msg.From)
	if !found {
		return types.ErrorResolverDoesNotExist().Result()
	}
	if resolver.Status == types.StatusDeRegistered {
		return types.ErrorInvalidResolverStatus().Result()
	}

	_resolver := types.Resolver{
		Owner:      msg.From,
		Commission: msg.Commission,
	}
	resolver = resolver.UpdateInfo(_resolver)

	k.SetResolver(ctx, resolver)

	return sdk.Result{}
}

func handleDeregisterResolver(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeregisterResolver)sdk.Result{
	resolver, found := k.GetResolver(ctx, msg.From)
	if !found {
		return types.ErrorResolverDoesNotExist().Result()
	}
	if resolver.Status != types.StatusRegistered {
		return types.ErrorInvalidResolverStatus().Result()
	}
	
	resolver.Status = types.StatusDeRegistered
	k.SetResolver(ctx,resolver)

	return sdk.Result{}
}
