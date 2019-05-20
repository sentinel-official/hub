package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func QueryNode(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) (*vpn.Node, error) {
	params := vpn.NewQueryNodeParams(id)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNode), paramBytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no node found")
	}

	var details vpn.Node
	if err := cdc.UnmarshalJSON(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

func QueryNodesOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, address csdkTypes.AccAddress) ([]byte, error) {
	params := vpn.NewQueryNodesOfAddressParams(address)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodesOfAddress), paramBytes)
}

func QuerySubscription(cliCtx context.CLIContext, cdc *codec.Codec, id string) (*vpn.Subscription, error) {
	subscriptionKey := vpn.SubscriptionKey(sdkTypes.IDFromString(id))
	res, err := cliCtx.QueryStore(subscriptionKey, vpn.StoreKeySubscription)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("no session found")
	}

	var details vpn.Subscription
	if err := cdc.UnmarshalBinaryLengthPrefixed(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
}
