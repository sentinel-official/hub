package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func QueryNode(cliCtx context.CLIContext, cdc *codec.Codec, id sdkTypes.ID) (vpn.Node, error) {
	params := vpn.NewQueryNodeParams(id)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return vpn.Node{}, err
	}

	res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNode), paramBytes)
	if err != nil {
		return vpn.Node{}, err
	}
	if res == nil {
		return vpn.Node{}, fmt.Errorf("no node found")
	}

	var node vpn.Node
	if err := cdc.UnmarshalJSON(res, &node); err != nil {
		return vpn.Node{}, err
	}

	return node, nil
}

func QueryNodesOfAddress(cliCtx context.CLIContext, cdc *codec.Codec, address csdkTypes.AccAddress) ([]byte, error) {
	params := vpn.NewQueryNodesOfAddressParams(address)
	paramBytes, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	return cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QueryNodesOfAddress), paramBytes)
}

// nolint: dupl
func QuerySubscription(cliCtx context.CLIContext, cdc *codec.Codec, id string) (vpn.Subscription, error) {
	subscriptionKey := vpn.SubscriptionKey(sdkTypes.IDFromString(id))
	res, err := cliCtx.QueryStore(subscriptionKey, vpn.StoreKeySubscription)
	if err != nil {
		return vpn.Subscription{}, err
	}
	if len(res) == 0 {
		return vpn.Subscription{}, fmt.Errorf("no subscription found")
	}

	var subscription vpn.Subscription
	if err := cdc.UnmarshalBinaryLengthPrefixed(res, &subscription); err != nil {
		return vpn.Subscription{}, err
	}

	return subscription, nil
}

// nolint: dupl
func QuerySession(cliCtx context.CLIContext, cdc *codec.Codec, id string) (vpn.Session, error) {
	sessionKey := vpn.SessionKey(sdkTypes.IDFromString(id))
	res, err := cliCtx.QueryStore(sessionKey, vpn.StoreKeySession)
	if err != nil {
		return vpn.Session{}, err
	}
	if len(res) == 0 {
		return vpn.Session{}, fmt.Errorf("no session found")
	}

	var session vpn.Session
	if err := cdc.UnmarshalBinaryLengthPrefixed(res, &session); err != nil {
		return vpn.Session{}, err
	}

	return session, nil
}
