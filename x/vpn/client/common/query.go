package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func QueryNode(cliCtx context.CLIContext, cdc *codec.Codec, id string) (*vpn.NodeDetails, error) {
	nodeKey := vpn.NodeKey(vpn.NewNodeID(id))
	res, err := cliCtx.QueryStore(nodeKey, vpn.StoreKeyNode)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("no node found")
	}

	var details vpn.NodeDetails
	if err := cdc.UnmarshalBinaryLengthPrefixed(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

func QueryNodesOfOwner(cliCtx context.CLIContext, cdc *codec.Codec, owner csdkTypes.AccAddress) ([]byte, error) {
	params := vpn.NewQueryNodesOfOwnerParams(owner)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodesOfOwner), paramBytes)
}

func QuerySession(cliCtx context.CLIContext, cdc *codec.Codec, id string) (*vpn.SessionDetails, error) {
	sessionKey := vpn.SessionKey(vpn.NewSessionID(id))
	res, err := cliCtx.QueryStore(sessionKey, vpn.StoreKeySession)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("no session found")
	}

	var details vpn.SessionDetails
	if err := cdc.UnmarshalBinaryLengthPrefixed(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
}
