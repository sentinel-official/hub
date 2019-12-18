package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func queryParameters(ctx sdk.Context, k keeper.Keeper) ([]byte, sdk.Error) {
	params := k.GetParams(ctx)

	res, err := types.ModuleCdc.MarshalJSON(params)
	if err != nil {
		return nil, types.ErrorUnmarshal()
	}

	return res, nil
}
