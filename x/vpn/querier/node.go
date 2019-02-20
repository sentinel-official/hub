package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	QueryNode         = "node"
	QueryNodesOfOwner = "nodesOfOwner"
)

func NewQuerier(vk keeper.Keeper, cdc *codec.Codec) csdkTypes.Querier {
	return func(ctx csdkTypes.Context, path []string, req abciTypes.RequestQuery) (res []byte, err csdkTypes.Error) {
		switch path[0] {
		case QueryNode:
			return queryNode(ctx, cdc, req, vk)
		case QueryNodesOfOwner:
			return queryNodesOfOwner(ctx, cdc, req, vk)
		default:
			return nil, types.ErrorInvalidQueryType(path[0])
		}
	}
}

type QueryNodeParams struct {
	ID sdkTypes.ID
}

func NewQueryNodeParams(id sdkTypes.ID) QueryNodeParams {
	return QueryNodeParams{
		ID: id,
	}
}

func queryNode(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery, vk keeper.Keeper) ([]byte, csdkTypes.Error) {
	var params QueryNodeParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	details, err := vk.GetNodeDetails(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if details == nil {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(details)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QueryNodesOfOwnerPrams struct {
	Owner csdkTypes.AccAddress
}

func NewQueryNodesOfOwnerParams(owner csdkTypes.AccAddress) QueryNodesOfOwnerPrams {
	return QueryNodesOfOwnerPrams{
		Owner: owner,
	}
}

func queryNodesOfOwner(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery, vk keeper.Keeper) ([]byte, csdkTypes.Error) {
	var params QueryNodesOfOwnerPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}
	nodes, err := vk.GetNodesOfOwner(ctx, params.Owner)
	if err != nil {
		return nil, err
	}

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
