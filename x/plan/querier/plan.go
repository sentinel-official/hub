package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	node "github.com/sentinel-official/hub/x/node/types"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func queryPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	plan, found := k.GetPlan(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(plan)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryPlans(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlansParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	plans := k.GetPlans(ctx)

	start, end := client.Paginate(len(plans), params.Page, params.Limit, len(plans))
	if start < 0 || end < 0 {
		plans = types.Plans{}
	} else {
		plans = plans[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(plans)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryPlansForProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryPlansForProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	plans := k.GetPlansForProvider(ctx, params.Address)

	start, end := client.Paginate(len(plans), params.Page, params.Limit, len(plans))
	if start < 0 || end < 0 {
		plans = types.Plans{}
	} else {
		plans = plans[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(plans)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryNodesForPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryNodesForPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	nodes := k.GetNodesForPlan(ctx, params.ID)

	start, end := client.Paginate(len(nodes), params.Page, params.Limit, len(nodes))
	if start < 0 || end < 0 {
		nodes = node.Nodes{}
	} else {
		nodes = nodes[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
