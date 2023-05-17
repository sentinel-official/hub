// DO NOT COVER

package simulation

import (
	"github.com/cosmos/cosmos-sdk/client"
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
	cdc codec.Codec,
	txConfig client.TxConfig,
	params simulationtypes.AppParams,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) []simulationtypes.WeightedOperation {
	var operations []simulationtypes.WeightedOperation
	operations = append(operations, providersimulation.WeightedOperations(cdc, txConfig, params, ak, bk, k.Provider)...)
	operations = append(operations, nodesimulation.WeightedOperations(params, cdc, ak, bk, k.Node)...)
	operations = append(operations, plansimulation.WeightedOperations(params, cdc, ak, bk, k.Plan)...)
	operations = append(operations, subscriptionsimulation.WeightedOperations(params, cdc, ak, bk, k.Subscription)...)
	operations = append(operations, sessionsimulation.WeightedOperations(params, cdc, ak, bk, k.Session)...)

	return operations
}
