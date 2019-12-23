package vpn

import (
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/sentinel-official/hub/x/vpn/types"
)

func InitGenesis(ctx sdk.Context, k Keeper, data types.GenesisState) {
	k.SetParams(ctx, data.Params)
	
	for _, node := range data.Nodes {
		k.SetNode(ctx, node)
		
		nca := k.GetNodesCountOfAddress(ctx, node.Owner)
		k.SetNodeIDByAddress(ctx, node.Owner, nca, node.ID)
		
		k.SetNodesCount(ctx, k.GetNodesCount(ctx)+1)
		k.SetNodesCountOfAddress(ctx, node.Owner, nca+1)
	}
	
	for _, subscription := range data.Subscriptions {
		k.SetSubscription(ctx, subscription)
		
		scn := k.GetSubscriptionsCountOfNode(ctx, subscription.NodeID)
		k.SetSubscriptionIDByNodeID(ctx, subscription.NodeID, scn, subscription.ID)
		
		sca := k.GetSubscriptionsCountOfAddress(ctx, subscription.Client)
		k.SetSubscriptionIDByAddress(ctx, subscription.Client, sca, subscription.ID)
		
		k.SetSubscriptionsCount(ctx, k.GetSubscriptionsCount(ctx)+1)
		k.SetSubscriptionsCountOfNode(ctx, subscription.NodeID, scn+1)
		k.SetSubscriptionsCountOfAddress(ctx, subscription.Client, sca+1)
	}
	
	for _, session := range data.Sessions {
		k.SetSession(ctx, session)
		
		scs := k.GetSessionsCountOfSubscription(ctx, session.SubscriptionID)
		k.SetSessionIDBySubscriptionID(ctx, session.SubscriptionID, scs, session.ID)
		
		k.SetSessionsCount(ctx, k.GetSessionsCount(ctx)+1)
		k.SetSessionsCountOfSubscription(ctx, session.SubscriptionID, scs+1)
	}
	
	for _, resolver := range data.Resolvers {
		k.SetResolver(ctx, resolver)
		
		rc := k.GetResolverCount(ctx)
		rca := k.GetResolversCountOfAddress(ctx, resolver.Owner)
		
		k.SetResolverIDByAddress(ctx, resolver.Owner, rca, resolver.ID)
		k.SetResolverCount(ctx, rc+1)
		k.SetResolverCountOfAddress(ctx, resolver.Owner, rca+1)
	}
	
	for _, freeClient := range data.FreeClients {
		k.SetFreeClient(ctx, freeClient)
		k.SetFreeClientOfNode(ctx, freeClient.NodeID, freeClient.Client)
		k.SetFreeNodeOfClient(ctx, freeClient.Client, freeClient.NodeID)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	params := k.GetParams(ctx)
	nodes := k.GetAllNodes(ctx)
	subscriptions := k.GetAllSubscriptions(ctx)
	sessions := k.GetAllSessions(ctx)
	resolvers := k.GetAllResolvers(ctx)
	freeClients := k.GetFreeClients(ctx)
	
	return types.NewGenesisState(nodes, subscriptions, sessions, resolvers, freeClients, params)
}

func ValidateGenesis(data types.GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}
	
	sessionsMap := make(map[uint64]bool, len(data.Sessions))
	for _, session := range data.Sessions {
		if err := session.IsValid(); err != nil {
			return fmt.Errorf("%s for the %s", err.Error(), session)
		}
		
		if sessionsMap[session.ID.Uint64()] {
			return fmt.Errorf("duplicate id for the %s", session)
		}
		
		sessionsMap[session.ID.Uint64()] = true
	}
	
	subscriptionsMap := make(map[uint64]bool, len(data.Subscriptions))
	for _, subscription := range data.Subscriptions {
		if err := subscription.IsValid(); err != nil {
			return fmt.Errorf("%s for the %s", err.Error(), subscription)
		}
		
		if subscriptionsMap[subscription.ID.Uint64()] {
			return fmt.Errorf("duplicate id for the %s", subscription)
		}
		
		subscriptionsMap[subscription.ID.Uint64()] = true
	}
	
	nodeIDsMap := make(map[uint64]bool, len(data.Nodes))
	for _, node := range data.Nodes {
		if err := node.IsValid(); err != nil {
			return fmt.Errorf("%s for the %s", err.Error(), node)
		}
		
		if node.Deposit.Denom != data.Params.Deposit.Denom {
			return fmt.Errorf("invalid deposit for the %s", node)
		}
		
		if nodeIDsMap[node.ID.Uint64()] {
			return fmt.Errorf("duplicate id for the %s", node)
		}
		
		nodeIDsMap[node.ID.Uint64()] = true
	}
	
	return nil
}
