package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func querySession(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	session, found := k.GetSession(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(session)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySessionOfSubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionOfSubscriptionPrams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	id, found := k.GetSessionIDBySubscriptionID(ctx, params.ID, params.Index)
	if !found {
		return nil, nil
	}

	session, found := k.GetSession(ctx, id)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(session)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func querySessionsOfSubscription(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QuerySessionsOfSubscriptionPrams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	sessions := k.GetSessionsOfSubscription(ctx, params.ID)

	res, err := types.ModuleCdc.MarshalJSON(sessions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryAllSessions(ctx sdk.Context, k keeper.Keeper) ([]byte, sdk.Error) {
	sessions := k.GetAllSessions(ctx)

	res, err := types.ModuleCdc.MarshalJSON(sessions)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
