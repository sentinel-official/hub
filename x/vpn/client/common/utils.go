package common

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func GetBandwidthSignDataBytes(cliCtx context.CLIContext, cdc *codec.Codec,
	subscriptionID string, bandwidth sdkTypes.Bandwidth) ([]byte, error) {

	subscription, err := QuerySubscription(cliCtx, cdc, subscriptionID)
	if err != nil {
		return nil, err
	}

	node, err := QueryNode(cliCtx, cdc, subscription.NodeID)
	if err != nil {
		return nil, err
	}

	sessionID := vpn.SessionID(subscription.ID, subscription.SessionsCount)
	return sdkTypes.NewBandwidthSignData(sessionID, bandwidth, node.Owner, subscription.Client).Bytes(), nil
}
