package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/subscription/keeper"
	"github.com/sentinel-official/hub/x/dvpn/subscription/types"
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

func querySubscriptions(ctx sdk.Context, _ abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	res, err := types.ModuleCdc.MarshalJSON(k.GetSubscriptions(ctx))
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsOfAddress(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsOfAddressParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	res, err := types.ModuleCdc.MarshalJSON(k.GetSubscriptionsForAddress(ctx, params.Address))
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsOfPlan(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsOfPlanParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	res, err := types.ModuleCdc.MarshalJSON(k.GetSubscriptionsForPlan(ctx, params.ID))
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySubscriptionsOfNode(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySubscriptionsOfNodeParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return nil, nil
}
