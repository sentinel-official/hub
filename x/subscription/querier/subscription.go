package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	plan, found := k.GetSubscription(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(plan)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptions(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptions(ctx, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsForAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsForAddressParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptionsForAddress(ctx, params.Address, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsForPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsForPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptionsForPlan(ctx, params.ID, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsForNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsForNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptionsForNode(ctx, params.Address, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(subscriptions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
