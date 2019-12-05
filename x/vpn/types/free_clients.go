package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type FreeClient struct {
	NodeID hub.NodeID     `json:"node_id"`
	Client sdk.AccAddress `json:"client"`
}

func (fc FreeClient) String() string {
	return fmt.Sprintf(`FreeClient
  NodeID:  %s
  Client:  %s`, fc.NodeID, fc.Client)
}

func NewFreeClient(nodeID hub.NodeID, client sdk.AccAddress) FreeClient {
	return FreeClient{
		NodeID: nodeID,
		Client: client,
	}
}
