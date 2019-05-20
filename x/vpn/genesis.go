package vpn

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func InitGenesis(ctx csdkTypes.Context, k Keeper, data types.GenesisState) {
	k.SetParams(ctx, data.Params)

	for _, node := range data.Nodes {
		k.SetNode(ctx, node)

		count := k.GetNodesCount(ctx, node.Owner)
		k.SetNodesCount(ctx, node.Owner, count+1)
	}

	for _, subscription := range data.Subscriptions {
		k.SetSubscription(ctx, subscription)

		node, _ := k.GetNode(ctx, subscription.NodeID)
		node.SubscriptionsCount += 1
		k.SetNode(ctx, node)
	}

	for _, session := range data.Sessions {
		k.SetSession(ctx, session)

		subscription, _ := k.GetSubscription(ctx, session.SubscriptionID)
		subscription.SessionsCount += 1
		k.SetSubscription(ctx, subscription)
	}
}

func ExportGenesis(ctx csdkTypes.Context, k Keeper) types.GenesisState {
	params := k.GetParams(ctx)
	nodes := k.GetAllNodes(ctx)
	subscriptions := k.GetAllSubscriptions(ctx)
	sessions := k.GetAllSessions(ctx)

	return types.NewGenesisState(nodes, subscriptions, sessions, params)
}

func ValidateGenesis(data types.GenesisState) error {
	return nil
}
