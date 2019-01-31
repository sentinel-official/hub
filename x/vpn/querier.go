package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryNode         = "node"
	QueryNodes        = "nodes"
	QueryNodesOfOwner = "nodesOfOwner"
)

func NewQuerier(k Keeper, cdc *codec.Codec) csdkTypes.Querier {
	return func(ctx csdkTypes.Context, path []string, req abciTypes.RequestQuery) (res []byte, err csdkTypes.Error) {
		switch path[0] {
		case QueryNode:
			return queryNode(ctx, cdc, req, k)
		case QueryNodes:
			return queryNodes(ctx, cdc, k)
		case QueryNodesOfOwner:
			return queryNodesOfOwner(ctx, cdc, req, k)
		default:
			return nil, errorInvalidQueryType(path[0])
		}
	}
}

type QueryNodeParams struct {
	ID string
}

func NewQueryNodeParams(id string) QueryNodeParams {
	return QueryNodeParams{
		ID: id,
	}
}

func queryNode(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery, k Keeper) ([]byte, csdkTypes.Error) {
	var params QueryNodeParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errorUnmarshal()
	}

	details, err := k.GetNodeDetails(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if details == nil {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(details)
	if resErr != nil {
		return nil, errorMarshal()
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

func queryNodesOfOwner(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery, k Keeper) ([]byte, csdkTypes.Error) {
	var params QueryNodesOfOwnerPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, errorUnmarshal()
	}

	nodes, err := k.GetNodesOfOwner(ctx, params.Owner)
	if err != nil {
		return nil, err
	}

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, errorMarshal()
	}

	return res, nil
}

func queryNodes(ctx csdkTypes.Context, cdc *codec.Codec, k Keeper) ([]byte, csdkTypes.Error) {
	nodes, err := k.GetNodes(ctx)
	if err != nil {
		return nil, err
	}

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, errorMarshal()
	}

	return res, nil
}
