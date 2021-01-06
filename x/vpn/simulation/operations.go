package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	node "github.com/sentinel-official/hub/x/node/simulation"
	plan "github.com/sentinel-official/hub/x/plan/simulation"
	provider "github.com/sentinel-official/hub/x/provider/simulation"
	session "github.com/sentinel-official/hub/x/session/simulation"
	subscription "github.com/sentinel-official/hub/x/subscription/simulation"
	"github.com/sentinel-official/hub/x/vpn/expected"
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func WeightedOperations(params simulation.AppParams, cdc *codec.Codec, ak expected.AccountKeeper, k keeper.Keeper) (operations simulation.WeightedOperations) {
	operations = append(operations, provider.WeightedOperations(params, cdc, ak, k.Provider)...)
	operations = append(operations, plan.WeightedOperations(params, cdc, ak, k.Provider, k.Node, k.Plan)...)
	operations = append(operations, node.WeightedOperations(params, cdc, ak, k.Provider, k.Node)...)
	operations = append(operations, subscription.WeightedOperations(params, cdc, ak, k.Plan, k.Node, k.Subscription)...)
	operations = append(operations, session.WeightedOperations(params, cdc, ak, k.Node, k.Subscription)...)

	return operations
}
