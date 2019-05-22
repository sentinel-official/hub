package common

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func GetBandwidthSignDataBytes(cliCtx context.CLIContext, cdc *codec.Codec, subscriptionID sdkTypes.ID,
	bandwidth sdkTypes.Bandwidth) ([]byte, error) {

	res, err := QuerySubscription(cliCtx, cdc, subscriptionID)
	if err != nil {
		return nil, err
	}

	var subscription vpn.Subscription
	if err := cdc.UnmarshalJSON(res, &subscription); err != nil {
		return nil, err
	}

	res, err = QueryNode(cliCtx, cdc, subscription.NodeID)
	if err != nil {
		return nil, err
	}

	var node vpn.Node
	if err := cdc.UnmarshalJSON(res, &node); err != nil {
		return nil, err
	}

	return sdkTypes.NewBandwidthSignData(subscription.ID, 0, bandwidth,
		node.Owner, subscription.Client).Bytes(), nil
}
