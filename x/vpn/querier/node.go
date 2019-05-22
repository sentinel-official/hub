package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	QueryNode           = "node"
	QueryNodesOfAddress = "nodesOfAddress"
)

func NewQuerier(vk keeper.Keeper, cdc *codec.Codec) csdkTypes.Querier {
	return func(ctx csdkTypes.Context, path []string, req abciTypes.RequestQuery) (res []byte, err csdkTypes.Error) {
		switch path[0] {
		case QueryNode:
			return queryNode(ctx, cdc, req, vk)
		case QueryNodesOfAddress:
			return queryNodesOfAddress(ctx, cdc, req, vk)
		default:
			return nil, types.ErrorInvalidQueryType(path[0])
		}
	}
}

type QueryNodeParams struct {
	ID uint64
}

func NewQueryNodeParams(id uint64) QueryNodeParams {
	return QueryNodeParams{
		ID: id,
	}
}

func queryNode(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery,
	vk keeper.Keeper) ([]byte, csdkTypes.Error) {

	var params QueryNodeParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	node, found := vk.GetNode(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(node)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QueryNodesOfAddressPrams struct {
	Address csdkTypes.AccAddress
}

func NewQueryNodesOfAddressParams(address csdkTypes.AccAddress) QueryNodesOfAddressPrams {
	return QueryNodesOfAddressPrams{
		Address: address,
	}
}

func queryNodesOfAddress(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery,
	vk keeper.Keeper) ([]byte, csdkTypes.Error) {

	var params QueryNodesOfAddressPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	nodes := vk.GetNodesOfAddress(ctx, params.Address)

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
