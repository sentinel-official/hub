package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	nodesim "github.com/sentinel-official/hub/x/node/simulation"
	plansim "github.com/sentinel-official/hub/x/plan/simulation"
	providersim "github.com/sentinel-official/hub/x/provider/simulation"
	sessionsim "github.com/sentinel-official/hub/x/session/simulation"
	subscriptionsim "github.com/sentinel-official/hub/x/subscription/simulation"
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func WeightedOperations(ap simulation.AppParams, cdc codec.JSONMarshaler, k keeper.Keeper) []simulation.WeightedOperation {
	var operations []simulation.WeightedOperation

	operations = append(operations, nodesim.WeightedOperations(ap, cdc, k.Node)...)
	operations = append(operations, sessionsim.WeightedOperations(ap, cdc, k.Session)...)
	operations = append(operations, subscriptionsim.WeightedOperations(ap, cdc, k.Subscription)...)
	operations = append(operations, plansim.WeightedOperations(ap, cdc, k.Plan)...)
	operations = append(operations, providersim.WeightedOperations(ap, cdc, k.Provider)...)

	return operations
}
