// nolint:dupl
package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	QueryNode           = "node"
	QueryNodesOfAddress = "nodesOfAddress"
	QueryAllNodes       = "allNodes"
)

type QueryNodeParams struct {
	ID sdk.ID
}

func NewQueryNodeParams(id sdk.ID) QueryNodeParams {
	return QueryNodeParams{
		ID: id,
	}
}

// nolint:dupl
func queryNode(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

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
	Address csdk.AccAddress
}

func NewQueryNodesOfAddressParams(address csdk.AccAddress) QueryNodesOfAddressPrams {
	return QueryNodesOfAddressPrams{
		Address: address,
	}
}

func queryNodesOfAddress(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

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

func queryAllNodes(ctx csdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, csdk.Error) {
	nodes := k.GetAllNodes(ctx)

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
