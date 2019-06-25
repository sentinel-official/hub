// nolint:dupl
package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/sentinel-hub/types"
	"github.com/sentinel-official/sentinel-hub/x/vpn"
)

func QueryNode(cliCtx context.CLIContext, cdc *codec.Codec, _id string) (*vpn.Node, error) {
	id := hub.NewIDFromString(_id)
	params := vpn.NewQueryNodeParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNode), paramBytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no node found")
	}

	var node vpn.Node
	if err = cdc.UnmarshalJSON(res, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

func QueryNodesOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, _address string) ([]vpn.Node, error) {
	address, err := sdk.AccAddressFromBech32(_address)
	if err != nil {
		return nil, err
	}

	params := vpn.NewQueryNodesOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodesOfAddress), paramBytes)
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
	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllNodes), nil)
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

func QuerySubscription(cliCtx context.CLIContext, cdc *codec.Codec, _id string) (*vpn.Subscription, error) {
	id := hub.NewIDFromString(_id)
	params := vpn.NewQuerySubscriptionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscription), paramBytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscription found")
	}

	var subscription vpn.Subscription
	if err = cdc.UnmarshalJSON(res, &subscription); err != nil {
		return nil, err
	}

	return &subscription, nil
}

func QuerySubscriptionsOfNode(cliCtx context.CLIContext, cdc *codec.Codec, _id string) ([]vpn.Subscription, error) {
	id := hub.NewIDFromString(_id)
	params := vpn.NewQuerySubscriptionsOfNodePrams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfNode), paramBytes)
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

func QuerySubscriptionsOfAddress(cliCtx context.CLIContext, cdc *codec.Codec,
	_address string) ([]vpn.Subscription, error) {

	address, err := sdk.AccAddressFromBech32(_address)
	if err != nil {
		return nil, err
	}

	params := vpn.NewQuerySubscriptionsOfAddressParams(address)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfAddress), paramBytes)
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
	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllSubscriptions), nil)
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

func QuerySessionsCountOfSubscription(cliCtx context.CLIContext, cdc *codec.Codec, _id string) (uint64, error) {
	id := hub.NewIDFromString(_id)
	params := vpn.NewQuerySessionsCountOfSubscriptionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return 0, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySessionsCountOfSubscription), paramBytes)
	if err != nil {
		return 0, err
	}
	if res == nil {
		return 0, fmt.Errorf("no sessions count found")
	}

	var count uint64
	if err = cdc.UnmarshalJSON(res, &count); err != nil {
		return 0, err
	}

	return count, nil
}

func QuerySession(cliCtx context.CLIContext, cdc *codec.Codec, _id string) (*vpn.Session, error) {
	id := hub.NewIDFromString(_id)
	params := vpn.NewQuerySessionParams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySession), paramBytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no session found")
	}

	var session vpn.Session
	if err = cdc.UnmarshalJSON(res, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func QuerySessionOfSubscription(cliCtx context.CLIContext, cdc *codec.Codec,
	_id string, index uint64) (*vpn.Session, error) {

	id := hub.NewIDFromString(_id)
	params := vpn.NewQuerySessionOfSubscriptionPrams(id, index)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySessionOfSubscription), paramBytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no session found")
	}

	var session vpn.Session
	if err = cdc.UnmarshalJSON(res, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func QuerySessionsOfSubscription(cliCtx context.CLIContext, cdc *codec.Codec, _id string) ([]vpn.Session, error) {
	id := hub.NewIDFromString(_id)
	params := vpn.NewQuerySessionsOfSubscriptionPrams(id)

	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySessionsOfSubscription), paramBytes)
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
	res, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryAllSessions), nil)
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
