package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
	"github.com/sentinel-official/hub/x/dvpn/types"
)

func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case provider.ModuleName:
			return provider.Querier(ctx, path[1:], req, k.Provider)
		case node.ModuleName:
			return node.Querier(ctx, path[1:], req, k.Node)
		case subscription.ModuleName:
			return subscription.Querier(ctx, path[1:], req, k.Subscription)
		default:
			return nil, types.ErrorUnknownQueryType(path[0])
		}
	}
}
