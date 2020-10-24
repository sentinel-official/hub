package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	provider, found := k.GetNode(ctx, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(provider)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryNodes(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	var nodes types.Nodes
	if params.Status.Equal(hub.StatusActive) {
		nodes = k.GetActiveNodes(ctx)
	} else if params.Status.Equal(hub.StatusInactive) {
		nodes = k.GetInActiveNodes(ctx)
	} else {
		nodes = append(
			k.GetActiveNodes(ctx),
			k.GetInActiveNodes(ctx)...,
		)
	}

	start, end := client.Paginate(len(nodes), params.Page, params.Limit, len(nodes))
	if start < 0 || end < 0 {
		nodes = types.Nodes{}
	} else {
		nodes = nodes[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryNodesForProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesForProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	var nodes types.Nodes
	if params.Status.Equal(hub.StatusActive) {
		nodes = k.GetActiveNodesForProvider(ctx, params.Address)
	} else if params.Status.Equal(hub.StatusInactive) {
		nodes = k.GetInActiveNodesForProvider(ctx, params.Address)
	} else {
		nodes = append(
			k.GetActiveNodesForProvider(ctx, params.Address),
			k.GetInActiveNodesForProvider(ctx, params.Address)...,
		)
	}

	start, end := client.Paginate(len(nodes), params.Page, params.Limit, len(nodes))
	if start < 0 || end < 0 {
		nodes = types.Nodes{}
	} else {
		nodes = nodes[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
