package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
	plan "github.com/sentinel-official/hub/x/plan/types"
)

func QueryNode(ctx context.CLIContext, address hub.NodeAddress) (*types.Node, error) {
	params := types.NewQueryNodeParams(address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryNode)
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

func QueryNodes(ctx context.CLIContext, status hub.Status, page, limit int) (types.Nodes, error) {
	params := types.NewQueryNodesParams(status, page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryNodes)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no nodes found")
	}

	var nodes types.Nodes
	if err := ctx.Codec.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

func QueryNodesForProvider(ctx context.CLIContext, address hub.ProvAddress, status hub.Status, page, limit int) (types.Nodes, error) {
	params := types.NewQueryNodesForProviderParams(address, status, page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryNodesForProvider)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no nodes found")
	}

	var nodes types.Nodes
	if err := ctx.Codec.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

func QueryNodesForPlan(ctx context.CLIContext, id uint64, page, limit int) (types.Nodes, error) {
	params := plan.NewQueryNodesForPlanParams(id, page, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", plan.StoreKey, plan.QuerierRoute, plan.QueryNodesForPlan)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no nodes found")
	}

	var nodes types.Nodes
	if err := ctx.Codec.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}
