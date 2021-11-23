package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	nodesimulation "github.com/sentinel-official/hub/x/node/simulation"
	plansimulation "github.com/sentinel-official/hub/x/plan/simulation"
	providersimulation "github.com/sentinel-official/hub/x/provider/simulation"
	sessionsimulation "github.com/sentinel-official/hub/x/session/simulation"
	subscriptionsimulation "github.com/sentinel-official/hub/x/subscription/simulation"
	"github.com/sentinel-official/hub/x/vpn/expected"
	"github.com/sentinel-official/hub/x/vpn/keeper"
)

func WeightedOperations(
	params simulationtypes.AppParams,
	appCodec codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) []simulationtypes.WeightedOperation {
	var operations []simulationtypes.WeightedOperation
	operations = append(operations, providersimulation.WeightedOperations(params, appCodec, ak, bk, k.Provider)...)
	operations = append(operations, nodesimulation.WeightedOperations(params, appCodec, ak, bk, k.Node)...)
	operations = append(operations, plansimulation.WeightedOperations(params, appCodec, ak, bk, k.Plan)...)
	operations = append(operations, subscriptionsimulation.WeightedOperations(params, appCodec, ak, bk, k.Subscription)...)
	operations = append(operations, sessionsimulation.WeightedOperations(params, appCodec, ak, bk, k.Session)...)

	return operations
}
