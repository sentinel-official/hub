package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func QueryNode(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) ([]byte, error) {
	params := vpn.NewQueryNodeParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNode), paramBytes)
}

func QueryNodesOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, address csdkTypes.AccAddress) ([]byte, error) {
	params := vpn.NewQueryNodesOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodesOfAddress), paramBytes)
}

func QueryAllNodes(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllNodes), nil)
}

func QuerySubscription(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) ([]byte, error) {
	params := vpn.NewQuerySubscriptionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscription), paramBytes)
}

func QuerySubscriptionsOfNode(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) ([]byte, error) {
	params := vpn.NewQuerySubscriptionsOfNodePrams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfNode), paramBytes)
}

func QuerySubscriptionsOAddress(cliCtx context.CLIContext, cdc *codec.Codec, address csdkTypes.AccAddress) ([]byte, error) {
	params := vpn.NewQuerySubscriptionsOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfAddress), paramBytes)
}

func QueryAllSubscriptions(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllSubscriptions), nil)
}

func QuerySession(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) ([]byte, error) {
	params := vpn.NewQuerySessionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySession), paramBytes)
}

func QuerySessionsOfSubscription(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) ([]byte, error) {
	params := vpn.NewQuerySessionsOfSubscriptionPrams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySessionsOfSubscription), paramBytes)
}

func QueryAllSessions(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllSessions), nil)
}
