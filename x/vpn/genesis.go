package vpn

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func InitGenesis(ctx csdkTypes.Context, k Keeper, data types.GenesisState) {
	k.SetParams(ctx, data.Params)

	for _, node := range data.Nodes {
		k.SetNode(ctx, node)

		count := k.GetNodesCount(ctx, node.Owner)
		k.SetNodesCount(ctx, node.Owner, count+1)

		if node.Status == types.StatusActive {
			k.AddActiveNodeIDAtHeight(ctx, ctx.BlockHeight(), node.ID)
		}
	}

	for _, session := range data.Sessions {
		k.SetSession(ctx, session)

		count := k.GetSessionsCount(ctx, session.Client)
		k.SetSessionsCount(ctx, session.Client, count+1)

		if session.Status == types.StatusActive {
			k.AddActiveSessionIDAtHeight(ctx, ctx.BlockHeight(), session.ID)
		}
	}
}

func ExportGenesis(ctx csdkTypes.Context, k Keeper) types.GenesisState {
	params := k.GetParams(ctx)
	nodes := k.GetNodes(ctx)
	sessions := k.GetAllSessions(ctx)

	return types.NewGenesisState(nodes, sessions, params)
}

func ValidateGenesis(data types.GenesisState) error {
	return nil
}
