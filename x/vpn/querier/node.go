// nolint:dupl
package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

const (
	QueryNode           = "node"
	QueryNodesOfAddress = "nodesOfAddress"
	QueryAllNodes       = "allNodes"
)

type QueryNodeParams struct {
	ID hub.ID
}

func NewQueryNodeParams(id hub.ID) QueryNodeParams {
	return QueryNodeParams{
		ID: id,
	}
}

// nolint:dupl
func queryNode(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

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
	Address sdk.AccAddress
}

func NewQueryNodesOfAddressParams(address sdk.AccAddress) QueryNodesOfAddressPrams {
	return QueryNodesOfAddressPrams{
		Address: address,
	}
}

func queryNodesOfAddress(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

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

func queryAllNodes(ctx sdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, sdk.Error) {
	nodes := k.GetAllNodes(ctx)

	res, resErr := cdc.MarshalJSON(nodes)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}
