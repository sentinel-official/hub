package vpn

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func InitGenesis(ctx csdkTypes.Context, k Keeper, data types.GenesisState) {
	k.SetParams(ctx, data.Params)

	for _, node := range data.Nodes {
		k.SetNode(ctx, node)

		count := k.GetNodesCount(ctx, node.Owner)
		k.SetNodesCount(ctx, node.Owner, count+1)
	}

	for _, subscription := range data.Subscriptions {
		k.SetSubscription(ctx, subscription)
	}

	for _, session := range data.Sessions {
		k.SetSession(ctx, session)
	}
}

func ExportGenesis(ctx csdkTypes.Context, k Keeper) types.GenesisState {
	params := k.GetParams(ctx)
	nodes := k.GetAllNodes(ctx)
	subscriptions := k.GetAllSubscriptions(ctx)
	sessions := k.GetAllSessions(ctx)

	return types.NewGenesisState(nodes, subscriptions, sessions, params)
}

func ValidateGenesis(data types.GenesisState) error {
	if len(data.Params.Deposit.Denom) == 0 || data.Params.Deposit.IsZero() {
		return fmt.Errorf("invalid deposit in the params %s", data.Params)
	}

	nodeIDsMap := make(map[string]bool, len(data.Nodes))
	for _, node := range data.Nodes {
		if err := node.IsValid(); err != nil {
			return fmt.Errorf("%s for the node %s", err.Error(), node)
		}

		if node.Deposit.Denom != data.Params.Deposit.Denom {
			return fmt.Errorf("invalid deposit for the node %s", node)
		}

		if nodeIDsMap[node.ID.String()] {
			return fmt.Errorf("duplicate node id for the node %s", node)
		}

		nodeIDsMap[node.ID.String()] = true
	}

	return nil
}
