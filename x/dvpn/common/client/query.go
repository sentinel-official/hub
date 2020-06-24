package client

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func QueryNodesOfPlan(ctx context.CLIContext, id uint64) (node.Nodes, error) {
	params := subscription.NewQueryNodesOfPlanParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, _, err := ctx.QueryWithData(subscription.QueryNodesOfPlanPath, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no nodes found")
	}

	var nodes node.Nodes
	if err := ctx.Codec.UnmarshalJSON(res, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}
