package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

func queryPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	plan, found := k.GetPlan(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(plan)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryPlans(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryPlansParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	plans := k.GetPlans(ctx, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(plans)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryPlansForProvider(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryPlansForProviderParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	plans := k.GetPlansForProvider(ctx, params.Address, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(plans)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryNodesForPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryNodesForPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	nodes := k.GetNodesForPlan(ctx, params.ID, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(nodes)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}
