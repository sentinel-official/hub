package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/sentinel-official/hub/x/swap/types"
)

func QuerySwap(ctx context.CLIContext, txHash types.EthereumHash) (*types.Swap, error) {
	params := types.NewQuerySwapParams(txHash)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bytes))

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySwap)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no swap found")
	}

	var swap types.Swap
	if err = ctx.Codec.UnmarshalJSON(res, &swap); err != nil {
		return nil, err
	}

	return &swap, nil
}

func QuerySwaps(ctx context.CLIContext, skip, limit int) (types.Swaps, error) {
	params := types.NewQuerySwapsParams(skip, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySwaps)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no swaps found")
	}

	var swaps types.Swaps
	if err = ctx.Codec.UnmarshalJSON(res, &swaps); err != nil {
		return nil, err
	}

	return swaps, nil
}
