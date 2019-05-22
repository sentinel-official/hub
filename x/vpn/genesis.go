package vpn

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func InitGenesis(ctx csdkTypes.Context, k Keeper, data types.GenesisState) {
	k.SetParams(ctx, data.Params)

	for _, node := range data.Nodes {
		node.SubscriptionsCount = 0
		k.SetNode(ctx, node)

		k.SetNodeIDByAddress(ctx, node.Owner, k.GetNodesCountOfAddress(ctx, node.Owner), node.ID)
		k.SetNodesCount(ctx, k.GetNodesCount(ctx)+1)
		k.SetNodesCountOfAddress(ctx, node.Owner, k.GetNodesCountOfAddress(ctx, node.Owner)+1)
	}

	for _, subscription := range data.Subscriptions {
		subscription.SessionsCount = 0
		k.SetSubscription(ctx, subscription)

		node, _ := k.GetNode(ctx, subscription.ID)
		k.SetSubscriptionIDByNodeID(ctx, node.ID, node.SubscriptionsCount, subscription.ID)
		k.SetSubscriptionIDByAddress(ctx, subscription.Client, k.GetSubscriptionsCountOfAddress(ctx, subscription.Client), subscription.ID)
		k.SetSubscriptionsCount(ctx, k.GetSubscriptionsCount(ctx)+1)
		node.SubscriptionsCount++
		k.SetNode(ctx, node)
		k.SetSubscriptionsCountOfAddress(ctx, subscription.Client, k.GetSubscriptionsCountOfAddress(ctx, subscription.Client)+1)
	}

	for _, session := range data.Sessions {
		k.SetSession(ctx, session)

		subscription, _ := k.GetSubscription(ctx, session.SubscriptionID)
		k.SetSessionIDBySubscriptionID(ctx, subscription.ID, subscription.SessionsCount, session.ID)
		k.SetSessionsCount(ctx, k.GetSessionsCount(ctx)+1)
		subscription.SessionsCount++
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

// nolint: gocyclo
func ValidateGenesis(data types.GenesisState) error {
	if len(data.Params.Deposit.Denom) == 0 || data.Params.Deposit.IsZero() {
		return fmt.Errorf("invalid deposit for the %s", data.Params)
	}

	sessionsMap := make(map[uint64]bool, len(data.Sessions))
	for _, session := range data.Sessions {
		if err := session.IsValid(); err != nil {
			return fmt.Errorf("%s for the %s", err.Error(), session)
		}

		if sessionsMap[session.ID] {
			return fmt.Errorf("duplicate id for the %s", session)
		}

		sessionsMap[session.ID] = true
	}

	subscriptionsMap := make(map[uint64]bool, len(data.Subscriptions))
	for _, subscription := range data.Subscriptions {
		if err := subscription.IsValid(); err != nil {
			return fmt.Errorf("%s for the %s", err.Error(), subscription)
		}

		if subscriptionsMap[subscription.ID] {
			return fmt.Errorf("duplicate id for the %s", subscription)
		}

		subscriptionsMap[subscription.ID] = true
	}

	nodeIDsMap := make(map[uint64]bool, len(data.Nodes))
	for _, node := range data.Nodes {
		if err := node.IsValid(); err != nil {
			return fmt.Errorf("%s for the %s", err.Error(), node)
		}

		if node.Deposit.Denom != data.Params.Deposit.Denom {
			return fmt.Errorf("invalid deposit for the %s", node)
		}

		if nodeIDsMap[node.ID] {
			return fmt.Errorf("duplicate id for the %s", node)
		}

		nodeIDsMap[node.ID] = true
	}

	return nil
}
