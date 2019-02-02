package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func QueryNode(cliCtx context.CLIContext, cdc *codec.Codec, nodeID string) ([]byte, error) {
	params := vpn.NewQueryNodeParams(nodeID)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNode), paramBytes)
}

func QueryNodes(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodes), nil)
}

func QueryNodesOfOwner(cliCtx context.CLIContext, cdc *codec.Codec, owner csdkTypes.AccAddress) ([]byte, error) {
	params := vpn.NewQueryNodesOfOwnerParams(owner)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodesOfOwner), paramBytes)
}
