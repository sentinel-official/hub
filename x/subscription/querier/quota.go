package querier

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func queryQuota(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryQuotaParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	quota, found := k.GetQuota(ctx, params.ID, params.Address)
	if !found {
		return nil, nil
	}

	res, err := types.ModuleCdc.MarshalJSON(quota)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryQuotas(ctx sdk.Context, req abci.RequestQuery, k keeper.Keeper) ([]byte, sdk.Error) {
	var params types.QueryQuotasParams
	if err := types.ModuleCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	quotas := k.GetQuotas(ctx, params.ID)

	start, end := client.Paginate(len(quotas), params.Page, params.Limit, len(quotas))
	if start < 0 || end < 0 {
		quotas = types.Quotas{}
	} else {
		quotas = quotas[start:end]
	}

	res, err := types.ModuleCdc.MarshalJSON(quotas)
	if err != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
