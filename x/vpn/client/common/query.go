package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

// nolint:dupl
func QueryNode(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) (node vpn.Node, err error) {
	params := vpn.NewQueryNodeParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return node, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNode), paramBytes)
	if err != nil {
		return node, err
	}
	if res == nil {
		return node, fmt.Errorf("no node found")
	}

	if err = cdc.UnmarshalJSON(res, &node); err != nil {
		return node, err
	}

	return node, nil
}

// nolint:dupl
func QueryNodesOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, address csdkTypes.AccAddress) ([]vpn.Node, error) {
	params := vpn.NewQueryNodesOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodesOfAddress), paramBytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no nodes found")
	}

	var nodes []vpn.Node
	if err := cdc.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

func QueryAllNodes(cliCtx context.CLIContext, cdc *codec.Codec) ([]vpn.Node, error) {
	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllNodes), nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no nodes found")
	}

	var nodes []vpn.Node
	if err := cdc.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

// nolint:dupl
func QuerySubscription(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) (subscription vpn.Subscription, err error) {
	params := vpn.NewQuerySubscriptionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return subscription, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscription), paramBytes)
	if err != nil {
		return subscription, err
	}
	if res == nil {
		return subscription, fmt.Errorf("no subscription found")
	}

	if err = cdc.UnmarshalJSON(res, &subscription); err != nil {
		return subscription, err
	}

	return subscription, nil
}

// nolint:dupl
func QuerySubscriptionsOfNode(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) ([]vpn.Subscription, error) {
	params := vpn.NewQuerySubscriptionsOfNodePrams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfNode), paramBytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions []vpn.Subscription
	if err := cdc.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

// nolint:dupl
func QuerySubscriptionsOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, address csdkTypes.AccAddress) ([]vpn.Subscription, error) {
	params := vpn.NewQuerySubscriptionsOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfAddress), paramBytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions []vpn.Subscription
	if err := cdc.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QueryAllSubscriptions(cliCtx context.CLIContext, cdc *codec.Codec) ([]vpn.Subscription, error) {
	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllSubscriptions), nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions []vpn.Subscription
	if err := cdc.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QuerySessionsCountOfSubscription(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) (count uint64, err error) {
	params := vpn.NewQuerySessionsCountOfSubscriptionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return count, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySessionsCountOfSubscription), paramBytes)
	if err != nil {
		return count, err
	}
	if res == nil {
		return count, fmt.Errorf("no sessions count found")
	}

	if err = cdc.UnmarshalJSON(res, &count); err != nil {
		return count, err
	}

	return count, nil
}

// nolint:dupl
func QuerySession(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) (session vpn.Session, err error) {
	params := vpn.NewQuerySessionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return session, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySession), paramBytes)
	if err != nil {
		return session, err
	}
	if res == nil {
		return session, fmt.Errorf("no session found")
	}

	if err = cdc.UnmarshalJSON(res, &session); err != nil {
		return session, err
	}

	return session, nil
}

// nolint:dupl
func QuerySessionsOfSubscription(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) ([]vpn.Session, error) {
	params := vpn.NewQuerySessionsOfSubscriptionPrams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySessionsOfSubscription), paramBytes)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no sessions found")
	}

	var sessions []vpn.Session
	if err := cdc.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func QueryAllSessions(cliCtx context.CLIContext, cdc *codec.Codec) ([]vpn.Session, error) {
	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllSessions), nil)
	if err != nil {
		return nil, err
	}
	if string(res) == "[]" || string(res) == "null" {
		return nil, fmt.Errorf("no sessions found")
	}

	var sessions []vpn.Session
	if err := cdc.UnmarshalJSON(res, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}
