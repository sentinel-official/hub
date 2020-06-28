package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/plan/types"
)

func QueryPlan(ctx context.CLIContext, id uint64) (*types.Plan, error) {
	params := types.NewQueryPlanParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryPlan)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no plan found")
	}

	var plan types.Plan
	if err := ctx.Codec.UnmarshalJSON(res, &plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

func QueryPlans(ctx context.CLIContext) (types.Plans, error) {
	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryPlans)
	res, _, err := ctx.QueryWithData(path, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no plans found")
	}

	var plans types.Plans
	if err := ctx.Codec.UnmarshalJSON(res, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}

func QueryPlansForProvider(ctx context.CLIContext, address hub.ProvAddress) (types.Plans, error) {
	params := types.NewQueryPlansForProviderParams(address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryPlansForProvider)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no plans found")
	}

	var plans types.Plans
	if err := ctx.Codec.UnmarshalJSON(res, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}
