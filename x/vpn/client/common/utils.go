package common

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func GetSessionBandwidthSignBytes(cliCtx context.CLIContext, cdc *codec.Codec,
	sessionID string, bandwidth sdkTypes.Bandwidth) ([]byte, error) {
	session, err := QuerySession(cliCtx, cdc, sessionID)
	if err != nil {
		return nil, err
	}

	node, err := QueryNode(cliCtx, cdc, session.NodeID.String())
	if err != nil {
		return nil, err
	}

	sign := vpn.NewBandwidthSign(session.ID, bandwidth, node.Owner, session.Client)
	return sign.GetBytes()
}
