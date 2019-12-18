package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/sentinel-official/hub/x/vpn/types"
)

func QueryParams(ctx context.CLIContext) (types.Params, error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)

	bz, _, err := ctx.QueryWithData(route, nil)
	if err != nil {
		return types.Params{}, err
	}

	var params types.Params
	ctx.Codec.MustUnmarshalJSON(bz, &params)

	return params, nil
}
