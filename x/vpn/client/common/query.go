package common

import (
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func QueryNode(ctx context.CLIContext, s string) (*types.Node, error) {
	id, err := hub.NewNodeIDFromString(s)
	if err != nil {
		return nil, err
	}
	params := types.NewQueryNodeParams(id)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryNode)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no node found")
	}
	
	var node types.Node
	if err := ctx.Codec.UnmarshalJSON(res, &node); err != nil {
		return nil, err
	}
	
	return &node, nil
}

func QueryNodesOfAddress(ctx context.CLIContext, s string) ([]types.Node, error) {
	address, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQueryNodesOfAddressParams(address)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryNodesOfAddress)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no nodes found")
	}
	
	var nodes []types.Node
	if err := ctx.Codec.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}
	
	return nodes, nil
}

func QueryAllNodes(ctx context.CLIContext) ([]types.Node, error) {
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllNodes)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no nodes found")
	}
	
	var nodes []types.Node
	if err := ctx.Codec.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}
	
	return nodes, nil
}

func QueryFreeNodesOfClient(ctx context.CLIContext, address string) ([]hub.NodeID, error) {
	_address, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQueryNodesOfFreeClientPrams(_address)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFreeNodesOfClient)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no free clients found")
	}
	
	var freeNodes []hub.NodeID
	if err := ctx.Codec.UnmarshalJSON(res, &freeNodes); err != nil {
		return nil, err
	}
	
	return freeNodes, nil
}

func QueryFreeClientsOfNode(ctx context.CLIContext, id string) ([]sdk.AccAddress, error) {
	nodeID, err := hub.NewNodeIDFromString(id)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQueryFreeClientsOfNodeParams(nodeID)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFreeClientsOfNode)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no free clients found")
	}
	
	var freeClients []sdk.AccAddress
	if err := ctx.Codec.UnmarshalJSON(res, &freeClients); err != nil {
		return nil, err
	}
	
	return freeClients, nil
}

func QueryNodesOfResolver(ctx context.CLIContext, s string) ([]hub.NodeID, error) {
	id, err := hub.NewResolverIDFromString(s)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQueryNodesOfResolverPrams(id)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryNodesOfResolver)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no resolvers found")
	}
	
	var freeNodes []hub.NodeID
	if err := ctx.Codec.UnmarshalJSON(res, &freeNodes); err != nil {
		return nil, err
	}
	
	return freeNodes, nil
}

func QueryResolversOfNode(ctx context.CLIContext, id string) ([]sdk.AccAddress, error) {
	nodeID, err := hub.NewNodeIDFromString(id)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQueryResolversOfNodeParams(nodeID)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolversOfNode)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no nodes found")
	}
	
	var resolvers []sdk.AccAddress
	if err := ctx.Codec.UnmarshalJSON(res, &resolvers); err != nil {
		return nil, err
	}
	
	return resolvers, nil
}

func QuerySubscription(ctx context.CLIContext, s string) (*types.Subscription, error) {
	id, err := hub.NewSubscriptionIDFromString(s)
	if err != nil {
		return nil, err
	}
	params := types.NewQuerySubscriptionParams(id)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscription found")
	}
	
	var subscription types.Subscription
	if err := ctx.Codec.UnmarshalJSON(res, &subscription); err != nil {
		return nil, err
	}
	
	return &subscription, nil
}

func QuerySubscriptionsOfNode(ctx context.CLIContext, s string) ([]types.Subscription, error) {
	id, err := hub.NewNodeIDFromString(s)
	if err != nil {
		return nil, err
	}
	params := types.NewQuerySubscriptionsOfNodePrams(id)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySubscriptionsOfNode)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no subscriptions found")
	}
	
	var subscriptions []types.Subscription
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}
	
	return subscriptions, nil
}

func QuerySubscriptionsOfAddress(ctx context.CLIContext, s string) ([]types.Subscription, error) {
	address, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQuerySubscriptionsOfAddressParams(address)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySubscriptionsOfAddress)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no subscriptions found")
	}
	
	var subscriptions []types.Subscription
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}
	
	return subscriptions, nil
}

func QueryAllSubscriptions(ctx context.CLIContext) ([]types.Subscription, error) {
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllSubscriptions)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no subscriptions found")
	}
	
	var subscriptions []types.Subscription
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}
	
	return subscriptions, nil
}

func QuerySessionsCountOfSubscription(ctx context.CLIContext, s string) (uint64, error) {
	id, err := hub.NewSubscriptionIDFromString(s)
	if err != nil {
		return 0, err
	}
	params := types.NewQuerySessionsCountOfSubscriptionParams(id)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return 0, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySessionsCountOfSubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return 0, err
	}
	if res == nil {
		return 0, fmt.Errorf("no sessions count found")
	}
	
	var count uint64
	if err := ctx.Codec.UnmarshalJSON(res, &count); err != nil {
		return 0, err
	}
	
	return count, nil
}

func QuerySession(ctx context.CLIContext, s string) (*types.Session, error) {
	id, err := hub.NewSessionIDFromString(s)
	if err != nil {
		return nil, err
	}
	
	params := types.NewQuerySessionParams(id)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySession)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no session found")
	}
	
	var session types.Session
	if err := ctx.Codec.UnmarshalJSON(res, &session); err != nil {
		return nil, err
	}
	
	return &session, nil
}

func QuerySessionOfSubscription(ctx context.CLIContext, s string, index uint64) (*types.Session, error) {
	id, err := hub.NewSubscriptionIDFromString(s)
	if err != nil {
		return nil, err
	}
	params := types.NewQuerySessionOfSubscriptionPrams(id, index)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySessionOfSubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no session found")
	}
	
	var session types.Session
	if err := ctx.Codec.UnmarshalJSON(res, &session); err != nil {
		return nil, err
	}
	
	return &session, nil
}

func QuerySessionsOfSubscription(ctx context.CLIContext, s string) ([]types.Session, error) {
	id, err := hub.NewSubscriptionIDFromString(s)
	if err != nil {
		return nil, err
	}
	params := types.NewQuerySessionsOfSubscriptionPrams(id)
	
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySessionsOfSubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no sessions found")
	}
	
	var sessions []types.Session
	if err := ctx.Codec.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}
	
	return sessions, nil
}

func QueryAllSessions(ctx context.CLIContext) ([]types.Session, error) {
	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllSessions)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no sessions found")
	}
	
	var sessions []types.Session
	if err := ctx.Codec.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}
	
	return sessions, nil
}
