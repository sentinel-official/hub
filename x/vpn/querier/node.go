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
	QueryNode           = "node"
	QueryNodesOfAddress = "nodesOfAddress"
	QueryAllNodes       = "allNodes"
)

type QueryNodeParams struct {
	ID sdkTypes.ID
}

func NewQueryNodeParams(id sdkTypes.ID) QueryNodeParams {
	return QueryNodeParams{
		ID: id,
	}
}

func queryNode(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery,
	k keeper.Keeper) ([]byte, csdkTypes.Error) {

	var params QueryNodeParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	node, found := k.GetNode(ctx, params.ID)
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
	k keeper.Keeper) ([]byte, csdkTypes.Error) {

	var params QueryNodesOfAddressPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	nodes := k.GetNodesOfAddress(ctx, params.Address)

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryAllNodes(ctx csdkTypes.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, csdkTypes.Error) {
	nodes := k.GetAllNodes(ctx)

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
