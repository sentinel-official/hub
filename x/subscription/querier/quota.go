package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func queryQuota(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryQuotaParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	quota, found := k.GetQuota(ctx, params.ID, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(quota)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}

func queryQuotas(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, error) {
	var params types.QueryQuotasParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errors.Wrap(types.ErrorUnmarshal, err.Error())
	}

	quotas := k.GetQuotas(ctx, params.ID, params.Skip, params.Limit)

	res, err := types.ModuleCdc.MarshalJSON(quotas)
	if err != nil {
		return nil, errors.Wrap(types.ErrorMarshal, err.Error())
	}

	return res, nil
}
